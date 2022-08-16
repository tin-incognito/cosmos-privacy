package coin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"privacy/x/privacy/common"
	"privacy/x/privacy/common/base58"
	"privacy/x/privacy/handler/key"
	"privacy/x/privacy/handler/operation"
)

type TxRandom [TxRandomGroupSize]byte

func NewTxRandom() *TxRandom {
	txRandom := new(operation.Point).Identity()
	index := uint32(0)

	res := new(TxRandom)
	res.SetTxConcealRandomPoint(txRandom)
	res.SetIndex(index)
	return res
}

func (t TxRandom) GetTxConcealRandomPoint() (*operation.Point, error) {
	return new(operation.Point).FromBytesS(t[operation.Ed25519KeySize+4:])
}

func (t TxRandom) GetTxOTARandomPoint() (*operation.Point, error) {
	return new(operation.Point).FromBytesS(t[:operation.Ed25519KeySize])
}

func (t TxRandom) GetIndex() (uint32, error) {
	return common.BytesToUint32(t[operation.Ed25519KeySize : operation.Ed25519KeySize+4])
}

func (t *TxRandom) SetTxConcealRandomPoint(txConcealRandom *operation.Point) {
	txRandomBytes := txConcealRandom.ToBytesS()
	copy(t[operation.Ed25519KeySize+4:], txRandomBytes)
}

func (t *TxRandom) SetTxOTARandomPoint(txRandom *operation.Point) {
	txRandomBytes := txRandom.ToBytesS()
	copy(t[:operation.Ed25519KeySize], txRandomBytes)
}

func (t *TxRandom) SetIndex(index uint32) {
	indexBytes := common.Uint32ToBytes(index)
	copy(t[operation.Ed25519KeySize:], indexBytes)
}

func (t TxRandom) Bytes() []byte {
	return t[:]
}

func (t *TxRandom) SetBytes(b []byte) error {
	if b == nil || len(b) != TxRandomGroupSize {
		return fmt.Errorf("cannot SetByte to TxRandom. Input is invalid")
	}
	_, err := new(operation.Point).FromBytesS(b[:operation.Ed25519KeySize])
	if err != nil {
		return fmt.Errorf("cannot TxRandomGroupSize.SetBytes: bytes is not operation.Point err: %v", err)
	}
	_, err = new(operation.Point).FromBytesS(b[operation.Ed25519KeySize+4:])
	if err != nil {
		return fmt.Errorf("cannot TxRandomGroupSize.SetBytes: bytes is not operation.Point err: %v", err)
	}
	copy(t[:], b)
	return nil
}

//nolint:revive // skip linter for this struct name
// Coin is the struct that will be stored to db
// If not privacy, mask and amount will be the original randomness and value
// If has privacy, mask and amount will be as paper monero
type Coin struct {
	// Public
	info       []byte
	publicKey  *operation.Point
	commitment *operation.Point
	keyImage   *operation.Point

	// sharedRandom and txRandom is shared secret between receiver and giver
	// sharedRandom is only visible when creating coins, when it is broadcast to network, it will be set to null
	sharedConcealRandom *operation.Scalar //rConceal: shared random when concealing output coin and blinding assetTag
	sharedRandom        *operation.Scalar // rOTA: shared random when creating one-time-address
	txRandom            *TxRandom         // rConceal*G + rOTA*G + index

	// mask = randomness
	// amount = value
	mask   *operation.Scalar
	amount *operation.Scalar
	// tag is nil unless confidential asset
	assetTag *operation.Point
}

// ParsePrivateKeyOfCoin retrieves the private OTA key of coin from the Master PrivateKey
func (c Coin) ParsePrivateKeyOfCoin(privKey key.PrivateKey) (*operation.Scalar, error) {
	keySet := new(key.KeySet)
	if err := keySet.InitFromPrivateKey(&privKey); err != nil {
		return nil, err
	}
	_, txRandomOTAPoint, index, err := c.GetTxRandomDetail()
	if err != nil {
		return nil, err
	}
	rK := new(operation.Point).ScalarMult(txRandomOTAPoint, keySet.OTAKey.GetOTASecretKey()) // (r_ota*G) * k = r_ota * K
	H := operation.HashToScalar(append(rK.ToBytesS(), common.Uint32ToBytes(index)...))       // Hash(r_ota*K, index)

	k := new(operation.Scalar).FromBytesS(privKey)
	return new(operation.Scalar).Add(H, k), nil // Hash(rK, index) + privSpend
}

// ParseKeyImageWithPrivateKey retrieves the keyImage of coin from the Master PrivateKey
func (c Coin) ParseKeyImageWithPrivateKey(privKey key.PrivateKey) (*operation.Point, error) {
	k, err := c.ParsePrivateKeyOfCoin(privKey)
	if err != nil {
		return nil, err
	}
	Hp := operation.HashToPoint(c.GetPublicKey().ToBytesS())
	return new(operation.Point).ScalarMult(Hp, k), nil
}

// Conceal the amount of coin using the publicView of the receiver
//
//	- AdditionalData: must be the publicView of the receiver
func (c *Coin) ConcealOutputCoin(additionalData *operation.Point) error {
	// If this coin is already encrypted or it is created by other person then cannot conceal
	if c.IsEncrypted() || c.GetSharedConcealRandom() == nil {
		return nil
	}
	publicView := additionalData
	rK := new(operation.Point).ScalarMult(publicView, c.GetSharedConcealRandom()) // rK = sharedConcealRandom * publicView

	hash := operation.HashToScalar(rK.ToBytesS()) // hash(rK)
	hash = operation.HashToScalar(hash.ToBytesS())
	mask := new(operation.Scalar).Add(c.GetRandomness(), hash) // mask = c.mask + hash

	hash = operation.HashToScalar(hash.ToBytesS())
	amount := new(operation.Scalar).Add(c.GetAmount(), hash) // amount = c.amout + hash
	c.SetRandomness(mask)
	c.SetAmount(amount)
	c.SetSharedConcealRandom(nil)
	c.SetSharedRandom(nil)
	return nil
}

// Conceal the input coin of a transaction: keep only the keyImage
func (c *Coin) ConcealInputCoin() {
	c.SetValue(0)
	c.SetRandomness(nil)
	c.SetPublicKey(nil)
	c.SetCommitment(nil)
	c.SetTxRandomDetail(new(operation.Point).Identity(), new(operation.Point).Identity(), 0)
}

// Decrypt a coin using the corresponding KeySet
func (c *Coin) Decrypt(keySet *key.KeySet) (*Coin, error) {
	if keySet == nil {
		return nil, fmt.Errorf("cannot Decrypt CoinV2: Keyset is empty")
	}

	// Must parse keyImage first in any situation
	if len(keySet.PrivateKey) > 0 {
		keyImage, err := c.ParseKeyImageWithPrivateKey(keySet.PrivateKey)
		if err != nil {
			return nil, fmt.Errorf("cannot parse key image with privateKey Coin - %v", err)
		}
		c.SetKeyImage(keyImage)
	}

	if !c.IsEncrypted() {
		return c, nil
	}

	viewKey := keySet.ReadonlyKey
	if len(viewKey.Rk) == 0 && len(keySet.PrivateKey) == 0 {
		return nil, fmt.Errorf("cannot Decrypt CoinV2: Keyset does not contain viewkey or privatekey")
	}

	if viewKey.GetPrivateView() != nil {
		txConcealRandomPoint, err := c.GetTxRandom().GetTxConcealRandomPoint()
		if err != nil {
			return nil, err
		}
		rK := new(operation.Point).ScalarMult(txConcealRandomPoint, viewKey.GetPrivateView())

		// Hash multiple times
		hash := operation.HashToScalar(rK.ToBytesS())
		hash = operation.HashToScalar(hash.ToBytesS())
		randomness := c.GetRandomness().Sub(c.GetRandomness(), hash)

		// Hash 1 more time to get value
		hash = operation.HashToScalar(hash.ToBytesS())
		value := c.GetAmount().Sub(c.GetAmount(), hash)

		commitment := operation.PedCom.CommitAtIndex(value, randomness, operation.PedersenValueIndex)
		// for `confidential asset` coin, we commit differently
		if c.GetAssetTag() != nil {
			com, err := ComputeCommitmentCA(c.GetAssetTag(), randomness, value)
			if err != nil {
				return nil, fmt.Errorf("cannot recompute commitment when decrypting")
			}
			commitment = com
		}
		if !operation.IsPointEqual(commitment, c.GetCommitment()) {
			return nil, fmt.Errorf("cannot Decrypt Coin: Commitment is not the same after decrypt")
		}
		c.SetRandomness(randomness)
		c.SetAmount(value)
	}
	return c, nil
}

func NewCoin() *Coin {
	c := new(Coin)
	c.info = []byte{}
	c.publicKey = new(operation.Point).Identity()
	c.commitment = new(operation.Point).Identity()
	c.keyImage = new(operation.Point).Identity()
	c.sharedRandom = new(operation.Scalar)
	c.txRandom = NewTxRandom()
	c.mask = new(operation.Scalar)
	c.amount = new(operation.Scalar)
	return c
}

func (c Coin) IsEncrypted() bool {
	if c.mask == nil || c.amount == nil {
		return true
	}
	tempCommitment := operation.PedCom.CommitAtIndex(c.amount, c.mask, operation.PedersenValueIndex)
	if c.GetAssetTag() != nil {
		// err is only for nil parameters, which we already checked, so here it is ignored
		com, _ := c.ComputeCommitmentCA()
		tempCommitment = com
	}
	return !operation.IsPointEqual(tempCommitment, c.commitment)
}

func (c Coin) GetRandomness() *operation.Scalar          { return c.mask }
func (c Coin) GetAmount() *operation.Scalar              { return c.amount }
func (c Coin) GetSharedRandom() *operation.Scalar        { return c.sharedRandom }
func (c Coin) GetSharedConcealRandom() *operation.Scalar { return c.sharedConcealRandom }
func (c Coin) GetPublicKey() *operation.Point            { return c.publicKey }
func (c Coin) GetCommitment() *operation.Point           { return c.commitment }
func (c Coin) GetKeyImage() *operation.Point             { return c.keyImage }
func (c Coin) GetInfo() []byte                           { return c.info }
func (c Coin) GetAssetTag() *operation.Point             { return c.assetTag }
func (c Coin) GetValue() uint64 {
	if c.IsEncrypted() {
		return 0
	}
	return c.amount.ToUint64Little()
}
func (c Coin) GetTxRandom() *TxRandom { return c.txRandom }
func (c Coin) GetTxRandomDetail() (*operation.Point, *operation.Point, uint32, error) {
	txRandomOTAPoint, err1 := c.txRandom.GetTxOTARandomPoint()
	txRandomConcealPoint, err2 := c.txRandom.GetTxConcealRandomPoint()
	index, err3 := c.txRandom.GetIndex()
	if err1 != nil || err2 != nil || err3 != nil {
		return nil, nil, 0, fmt.Errorf("cannot Get TxRandom: point or index is wrong")
	}
	return txRandomConcealPoint, txRandomOTAPoint, index, nil
}

func (c Coin) GetCoinDetailEncrypted() []byte {
	return c.GetAmount().ToBytesS()
}

func (c *Coin) GetCoinID() [operation.Ed25519KeySize]byte {
	if c.publicKey != nil {
		return c.publicKey.ToBytes()
	}
	return [operation.Ed25519KeySize]byte{}
}

func (c *Coin) SetRandomness(mask *operation.Scalar)           { c.mask = mask }
func (c *Coin) SetAmount(amount *operation.Scalar)             { c.amount = amount }
func (c *Coin) SetSharedRandom(sharedRandom *operation.Scalar) { c.sharedRandom = sharedRandom }
func (c *Coin) SetSharedConcealRandom(sharedConcealRandom *operation.Scalar) {
	c.sharedConcealRandom = sharedConcealRandom
}
func (c *Coin) SetTxRandom(txRandom *TxRandom) {
	if txRandom == nil {
		c.txRandom = nil
	} else {
		c.txRandom = NewTxRandom()
		_ = c.txRandom.SetBytes(txRandom.Bytes())
	}
}
func (c *Coin) SetTxRandomDetail(txConcealRandomPoint, txRandomPoint *operation.Point, index uint32) {
	res := new(TxRandom)
	res.SetTxConcealRandomPoint(txConcealRandomPoint)
	res.SetTxOTARandomPoint(txRandomPoint)
	res.SetIndex(index)
	c.txRandom = res
}

func (c *Coin) SetPublicKey(publicKey *operation.Point)   { c.publicKey = publicKey }
func (c *Coin) SetCommitment(commitment *operation.Point) { c.commitment = commitment }
func (c *Coin) SetKeyImage(keyImage *operation.Point)     { c.keyImage = keyImage }
func (c *Coin) SetValue(value uint64)                     { c.amount = new(operation.Scalar).FromUint64(value) }
func (c *Coin) SetInfo(b []byte) {
	c.info = make([]byte, len(b))
	copy(c.info, b)
}
func (c *Coin) SetAssetTag(at *operation.Point) { c.assetTag = at }

func (c Coin) Bytes() []byte {
	coinBytes := []byte{}
	info := c.GetInfo()
	byteLengthInfo := byte(getMin(len(info), MaxSizeInfoCoin))
	coinBytes = append(coinBytes, byteLengthInfo)
	coinBytes = append(coinBytes, info[:byteLengthInfo]...)

	if c.publicKey != nil {
		coinBytes = append(coinBytes, byte(operation.Ed25519KeySize))
		coinBytes = append(coinBytes, c.publicKey.ToBytesS()...)
	} else {
		coinBytes = append(coinBytes, byte(0))
	}

	if c.commitment != nil {
		coinBytes = append(coinBytes, byte(operation.Ed25519KeySize))
		coinBytes = append(coinBytes, c.commitment.ToBytesS()...)
	} else {
		coinBytes = append(coinBytes, byte(0))
	}

	if c.keyImage != nil {
		coinBytes = append(coinBytes, byte(operation.Ed25519KeySize))
		coinBytes = append(coinBytes, c.keyImage.ToBytesS()...)
	} else {
		coinBytes = append(coinBytes, byte(0))
	}

	if c.sharedRandom != nil {
		coinBytes = append(coinBytes, byte(operation.Ed25519KeySize))
		coinBytes = append(coinBytes, c.sharedRandom.ToBytesS()...)
	} else {
		coinBytes = append(coinBytes, byte(0))
	}

	if c.sharedConcealRandom != nil {
		coinBytes = append(coinBytes, byte(operation.Ed25519KeySize))
		coinBytes = append(coinBytes, c.sharedConcealRandom.ToBytesS()...)
	} else {
		coinBytes = append(coinBytes, byte(0))
	}

	if c.txRandom != nil {
		coinBytes = append(coinBytes, TxRandomGroupSize)
		coinBytes = append(coinBytes, c.txRandom.Bytes()...)
	} else {
		coinBytes = append(coinBytes, byte(0))
	}

	if c.mask != nil {
		coinBytes = append(coinBytes, byte(operation.Ed25519KeySize))
		coinBytes = append(coinBytes, c.mask.ToBytesS()...)
	} else {
		coinBytes = append(coinBytes, byte(0))
	}

	if c.amount != nil {
		coinBytes = append(coinBytes, byte(operation.Ed25519KeySize))
		coinBytes = append(coinBytes, c.amount.ToBytesS()...)
	} else {
		coinBytes = append(coinBytes, byte(0))
	}

	if c.assetTag != nil {
		coinBytes = append(coinBytes, byte(operation.Ed25519KeySize))
		coinBytes = append(coinBytes, c.assetTag.ToBytesS()...)
	} else {
		coinBytes = append(coinBytes, byte(0))
	}

	return coinBytes
}

func (c *Coin) SetBytes(coinBytes []byte) error {
	var err error
	if c == nil {
		return fmt.Errorf("cannot set bytes for unallocated CoinV2")
	}
	if len(coinBytes) == 0 {
		return fmt.Errorf("coinBytes is empty")
	}

	offset := 0
	c.info, err = parseInfoForSetBytes(&coinBytes, &offset)
	if err != nil {
		return fmt.Errorf("setBytes CoinV2 info error: %v", err)
	}

	c.publicKey, err = parsePointForSetBytes(&coinBytes, &offset)
	if err != nil {
		return fmt.Errorf("setBytes CoinV2 publicKey error: %v", err)
	}
	c.commitment, err = parsePointForSetBytes(&coinBytes, &offset)
	if err != nil {
		return fmt.Errorf("setBytes CoinV2 commitment error: %v", err)
	}
	c.keyImage, err = parsePointForSetBytes(&coinBytes, &offset)
	if err != nil {
		return fmt.Errorf("setBytes CoinV2 keyImage error: %v", err)
	}
	c.sharedRandom, err = parseScalarForSetBytes(&coinBytes, &offset)
	if err != nil {
		return fmt.Errorf("setBytes CoinV2 mask error: %v", err)
	}

	c.sharedConcealRandom, err = parseScalarForSetBytes(&coinBytes, &offset)
	if err != nil {
		return fmt.Errorf("setBytes CoinV2 mask error: %v", err)
	}

	if offset >= len(coinBytes) {
		return fmt.Errorf("offset is larger than len(bytes), cannot parse txRandom")
	}
	if coinBytes[offset] != TxRandomGroupSize {
		return fmt.Errorf("setBytes CoinV2 TxRandomGroup error: length of TxRandomGroup is not correct")
	}
	offset++
	if offset+TxRandomGroupSize > len(coinBytes) {
		return fmt.Errorf("setBytes CoinV2 TxRandomGroup error: length of coinBytes is too small")
	}
	c.txRandom = NewTxRandom()
	err = c.txRandom.SetBytes(coinBytes[offset : offset+TxRandomGroupSize])
	if err != nil {
		return fmt.Errorf("setBytes CoinV2 TxRandomGroup error: %v", err)
	}
	offset += TxRandomGroupSize

	if err != nil {
		return fmt.Errorf("setBytes CoinV2 txRandom error: %v", err)
	}
	c.mask, err = parseScalarForSetBytes(&coinBytes, &offset)
	if err != nil {
		return fmt.Errorf("setBytes CoinV2 mask error: %v", err)
	}
	c.amount, err = parseScalarForSetBytes(&coinBytes, &offset)
	if err != nil {
		return fmt.Errorf("setBytes CoinV2 amount error: %v", err)
	}

	if offset >= len(coinBytes) {
		// for parsing old serialization, which does not have assetTag field
		c.assetTag = nil
	} else {
		c.assetTag, err = parsePointForSetBytes(&coinBytes, &offset)
		if err != nil {
			return fmt.Errorf("setBytes CoinV2 assetTag error: %v", err)
		}
	}
	return nil
}

// HashH returns the SHA3-256 hashing of coin bytes array
func (c *Coin) HashH() *common.Hash {
	hash := common.HashH(c.Bytes())
	return &hash
}

func (c Coin) MarshalJSON() ([]byte, error) {
	data := c.Bytes()
	temp := base58.Base58Check{}.Encode(data, common.ZeroByte)
	return json.Marshal(temp)
}

func (c *Coin) UnmarshalJSON(data []byte) error {
	dataStr := ""
	_ = json.Unmarshal(data, &dataStr)
	temp, _, err := base58.Base58Check{}.Decode(dataStr)
	if err != nil {
		return err
	}
	err = c.SetBytes(temp)
	if err != nil {
		return err
	}
	return nil
}

func (c *Coin) CheckCoinValid(paymentAdd key.PaymentAddress, sharedRandom []byte, amount uint64) bool {
	if c.GetValue() != amount {
		return false
	}
	// check one-time address is corresponding to paymentaddress
	r := new(operation.Scalar).FromBytesS(sharedRandom)
	if !r.ScalarValid() {
		return false
	}

	publicOTA := paymentAdd.GetOTAPublicKey()
	if publicOTA == nil {
		return false
	}
	rK := new(operation.Point).ScalarMult(publicOTA, r)
	_, txOTARandomPoint, index, err := c.GetTxRandomDetail()
	if err != nil {
		return false
	}
	if !operation.IsPointEqual(new(operation.Point).ScalarMultBase(r), txOTARandomPoint) {
		return false
	}

	hash := operation.HashToScalar(append(rK.ToBytesS(), common.Uint32ToBytes(index)...))
	HrKG := new(operation.Point).ScalarMultBase(hash)
	tmpPubKey := new(operation.Point).Add(HrKG, paymentAdd.GetPublicSpend())
	return bytes.Equal(tmpPubKey.ToBytesS(), c.publicKey.ToBytesS())
}

// Check whether the utxo is from this keyset
func (c *Coin) DoesCoinBelongToKeySet(keySet *key.KeySet) (bool, *operation.Point) {
	_, txOTARandomPoint, index, err1 := c.GetTxRandomDetail()
	if err1 != nil {
		return false, nil
	}

	// Check if the utxo belong to this one-time-address
	rK := new(operation.Point).ScalarMult(txOTARandomPoint, keySet.OTAKey.GetOTASecretKey())

	hashed := operation.HashToScalar(
		append(rK.ToBytesS(), common.Uint32ToBytes(index)...),
	)

	HnG := new(operation.Point).ScalarMultBase(hashed)
	KCheck := new(operation.Point).Sub(c.GetPublicKey(), HnG)

	// // Retrieve the sharedConcealRandomPoint for generating the blinded assetTag
	// var rSharedConcealPoint *operation.Point
	// if keySet.ReadonlyKey.GetPrivateView() != nil {
	// 	rSharedConcealPoint = new(operation.Point).ScalarMult(txConcealRandomPoint, keySet.ReadonlyKey.GetPrivateView())
	// }

	return operation.IsPointEqual(KCheck, keySet.OTAKey.GetPublicSpend()), rK
}
