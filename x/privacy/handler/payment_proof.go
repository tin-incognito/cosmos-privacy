package handler

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"privacy/x/privacy/handler/bulletproofs"
	"privacy/x/privacy/handler/coin"

	"github.com/incognitochain/incognito-chain/common"
)

type PaymentProof struct {
	aggregatedRangeProof *bulletproofs.AggregatedRangeProof
	inputCoins           []*coin.Coin
	outputCoins          []*coin.Coin
}

// GetInputCoins is the getter for input coins.
func (proof PaymentProof) GetInputCoins() []coin.Coin { return proof.inputCoins }

// GetOutputCoins is the getter for output coins.
func (proof PaymentProof) GetOutputCoins() []coin.Coin {
	res := make([]coin.Coin, len(proof.outputCoins))
	for i := 0; i < len(proof.outputCoins); i++ {
		res[i] = *proof.outputCoins[i]
	}
	return res
}

// GetAggregatedRangeProof returns the Bulletproof in this, but as a generic range proof object.
func (proof PaymentProof) GetAggregatedRangeProof() *bulletproofs.AggregatedRangeProof {
	return proof.aggregatedRangeProof
}

func (proof *PaymentProof) SetInputCoins(v []coin.Coin) error {
	proof.inputCoins = make([]coin.Coin, len(v))
	for i := 0; i < len(v); i++ {
		b := v[i].Bytes()
		if err := proof.inputCoins[i].SetBytes(b); err != nil {
			return err
		}
	}
	return nil
}

func (proof *PaymentProof) SetOutputCoins(v []coin.Coin) error {
	var err error
	proof.outputCoins = make([]*coin.Coin, len(v))
	for i := 0; i < len(v); i++ {
		proof.outputCoins[i] = new(coin.Coin)
		b := v[i].Bytes()
		if err = proof.outputCoins[i].SetBytes(b); err != nil {
			return err
		}
	}
	return nil
}

func (proof *PaymentProof) SetAggregatedRangeProof(aggregatedRangeProof *bulletproofs.AggregatedRangeProof) {
	proof.aggregatedRangeProof = aggregatedRangeProof
}

func NewPaymentProof() *PaymentProof {
	proof := NewPaymentProof()
	aggregatedRangeProof := &bulletproofs.AggregatedRangeProof{}
	//TODO: @tin init proof here
	//aggregatedRangeProof.Init()
	proof.aggregatedRangeProof = aggregatedRangeProof
	proof.inputCoins = []coin.Coin{}
	proof.outputCoins = []*coin.CoinV2{}
	return proof
}

// MarshalJSON implements JSON Marshaller
func (proof PaymentProofV2) MarshalJSON() ([]byte, error) {
	data := proof.Bytes()
	// temp := base58.Base58Check{}.Encode(data, common.ZeroByte)
	temp := base64.StdEncoding.EncodeToString(data)
	return json.Marshal(temp)
}

// UnmarshalJSON implements JSON Unmarshaller
func (proof *PaymentProofV2) UnmarshalJSON(data []byte) error {
	dataStr := common.EmptyString
	errJson := json.Unmarshal(data, &dataStr)
	if errJson != nil {
		Logger.Log.Errorf("PaymentProofV2 unmarshalling dataStr error: %v\n", errJson)
		return errJson
	}
	temp, err := base64.StdEncoding.DecodeString(dataStr)
	if err != nil {
		Logger.Log.Errorf("PaymentProofV2 decodeing string dataStr error: %v\n", err)
		return err
	}

	errSetBytes := proof.SetBytes(temp)
	if errSetBytes != nil {
		Logger.Log.Errorf("PaymentProofV2 setbytes error: %v\n", errSetBytes)
		return errSetBytes
	}
	return nil
}

// Bytes does byte serialization for this payment proof
func (proof PaymentProofV2) Bytes() []byte {
	var bytes []byte
	bytes = append(bytes, proof.GetVersion())

	comOutputMultiRangeProof := proof.aggregatedRangeProof.Bytes()
	var rangeProofLength uint32 = uint32(len(comOutputMultiRangeProof))
	bytes = append(bytes, common.Uint32ToBytes(rangeProofLength)...)
	bytes = append(bytes, comOutputMultiRangeProof...)

	// InputCoins
	bytes = append(bytes, byte(len(proof.inputCoins)))
	for i := 0; i < len(proof.inputCoins); i++ {
		inputCoins := proof.inputCoins[i].Bytes()
		lenInputCoins := len(inputCoins)
		var lenInputCoinsBytes []byte
		if lenInputCoins < 256 {
			lenInputCoinsBytes = []byte{byte(lenInputCoins)}
		} else {
			lenInputCoinsBytes = common.IntToBytes(lenInputCoins)
		}

		bytes = append(bytes, lenInputCoinsBytes...)
		bytes = append(bytes, inputCoins...)
	}

	// OutputCoins
	bytes = append(bytes, byte(len(proof.outputCoins)))
	for i := 0; i < len(proof.outputCoins); i++ {
		outputCoins := proof.outputCoins[i].Bytes()
		lenOutputCoins := len(outputCoins)
		var lenOutputCoinsBytes []byte
		if lenOutputCoins < 256 {
			lenOutputCoinsBytes = []byte{byte(lenOutputCoins)}
		} else {
			lenOutputCoinsBytes = common.IntToBytes(lenOutputCoins)
		}

		bytes = append(bytes, lenOutputCoinsBytes...)
		bytes = append(bytes, outputCoins...)
	}

	return bytes
}

// SetBytes does byte deserialization for this payment proof
func (proof *PaymentProofV2) SetBytes(proofbytes []byte) *errhandler.PrivacyError {
	if len(proofbytes) == 0 {
		return errhandler.NewPrivacyErr(errhandler.InvalidInputToSetBytesErr, errors.New("Proof bytes is zero"))
	}
	if proofbytes[0] != proof.GetVersion() {
		Logger.Log.Errorf("proof bytes version is incorrect: %v != %v\n", proofbytes[0], proof.GetVersion())
		return errhandler.NewPrivacyErr(errhandler.SetBytesProofErr, errors.New("Proof bytes version is incorrect"))
	}
	proof.SetVersion()
	offset := 1

	// ComOutputMultiRangeProofSize *aggregatedRangeProof
	if offset+common.Uint32Size >= len(proofbytes) {
		Logger.Log.Errorf("out of range aggregated range proof: %v + %v >= %v\n", offset, common.Uint32Size, len(proofbytes))
		return errhandler.NewPrivacyErr(errhandler.SetBytesProofErr, errors.New("Out of range aggregated range proof"))
	}
	lenComOutputMultiRangeUint32, _ := common.BytesToUint32(proofbytes[offset : offset+common.Uint32Size])
	lenComOutputMultiRangeProof := int(lenComOutputMultiRangeUint32)
	offset += common.Uint32Size

	if offset+lenComOutputMultiRangeProof > len(proofbytes) {
		Logger.Log.Errorf("out of range aggregated range proof: %v + %v >= %v\n", offset, lenComOutputMultiRangeProof, len(proofbytes))
		return errhandler.NewPrivacyErr(errhandler.SetBytesProofErr, errors.New("Out of range aggregated range proof"))
	}
	if lenComOutputMultiRangeProof > 0 {
		bulletproof := &bulletproofs.AggregatedRangeProof{}
		bulletproof.Init()
		proof.aggregatedRangeProof = bulletproof
		err := proof.aggregatedRangeProof.SetBytes(proofbytes[offset : offset+lenComOutputMultiRangeProof])
		if err != nil {
			Logger.Log.Errorf("aggregated range proof setbytes error: %v\n", err)
			return errhandler.NewPrivacyErr(errhandler.SetBytesProofErr, err)
		}
		offset += lenComOutputMultiRangeProof
	}

	if offset >= len(proofbytes) {
		Logger.Log.Errorf("out of range input coin: %v >= %v\n", offset, len(proofbytes))
		return errhandler.NewPrivacyErr(errhandler.SetBytesProofErr, errors.New("Out of range input coins"))
	}
	lenInputCoinsArray := int(proofbytes[offset])
	offset++
	proof.inputCoins = make([]coin.PlainCoin, lenInputCoinsArray)
	var err error
	for i := 0; i < lenInputCoinsArray; i++ {
		// try get 1-byte for len
		if offset >= len(proofbytes) {
			Logger.Log.Errorf("out of range input coin: %v >= %v\n", offset, len(proofbytes))
			return errhandler.NewPrivacyErr(errhandler.SetBytesProofErr, errors.New("Out of range input coins"))
		}
		lenInputCoin := int(proofbytes[offset])
		offset++

		if offset+lenInputCoin > len(proofbytes) {
			Logger.Log.Errorf("out of range input coin: %v + %v >= %v\n", offset, lenInputCoin, len(proofbytes))
			return errhandler.NewPrivacyErr(errhandler.SetBytesProofErr, errors.New("Out of range input coins"))
		}
		proof.inputCoins[i], err = coin.NewPlainCoinFromByte(proofbytes[offset : offset+lenInputCoin])
		if err != nil {
			// 1-byte is wrong
			// try get 2-byte for len
			if offset+1 > len(proofbytes) {
				Logger.Log.Error("out of range input coin")
				return errhandler.NewPrivacyErr(errhandler.SetBytesProofErr, errors.New("Out of range input coins"))
			}
			lenInputCoin = common.BytesToInt(proofbytes[offset-1 : offset+1])
			offset++

			if offset+lenInputCoin > len(proofbytes) {
				Logger.Log.Error("out of range input coin")
				return errhandler.NewPrivacyErr(errhandler.SetBytesProofErr, errors.New("Out of range input coins"))
			}
			proof.inputCoins[i], err = coin.NewPlainCoinFromByte(proofbytes[offset : offset+lenInputCoin])
			if err != nil {
				Logger.Log.Errorf("input coin setbytes error: %v\n", err)
				return errhandler.NewPrivacyErr(errhandler.SetBytesProofErr, err)
			}
		}
		offset += lenInputCoin
	}

	if offset >= len(proofbytes) {
		Logger.Log.Error("out of range output coin: %v >= %v\n", offset, len(proofbytes))
		return errhandler.NewPrivacyErr(errhandler.SetBytesProofErr, errors.New("Out of range output coins"))
	}
	lenOutputCoinsArray := int(proofbytes[offset])
	offset++
	proof.outputCoins = make([]*coin.CoinV2, lenOutputCoinsArray)
	for i := 0; i < lenOutputCoinsArray; i++ {
		proof.outputCoins[i] = new(coin.CoinV2)
		// try get 1-byte for len
		if offset >= len(proofbytes) {
			Logger.Log.Error("out of range output coin: %v >= %v\n", offset, len(proofbytes))
			return errhandler.NewPrivacyErr(errhandler.SetBytesProofErr, errors.New("Out of range output coins"))
		}
		lenOutputCoin := int(proofbytes[offset])
		offset++

		if offset+lenOutputCoin > len(proofbytes) {
			Logger.Log.Error("out of range output coin: %v + %v >= %v\n", offset, lenOutputCoin, len(proofbytes))
			return errhandler.NewPrivacyErr(errhandler.SetBytesProofErr, errors.New("Out of range output coins"))
		}
		err := proof.outputCoins[i].SetBytes(proofbytes[offset : offset+lenOutputCoin])
		if err != nil {
			// 1-byte is wrong
			// try get 2-byte for len
			if offset+1 > len(proofbytes) {
				return errhandler.NewPrivacyErr(errhandler.SetBytesProofErr, errors.New("Out of range output coins"))
			}
			lenOutputCoin = common.BytesToInt(proofbytes[offset-1 : offset+1])
			offset++

			if offset+lenOutputCoin > len(proofbytes) {
				return errhandler.NewPrivacyErr(errhandler.SetBytesProofErr, errors.New("Out of range output coins"))
			}
			err1 := proof.outputCoins[i].SetBytes(proofbytes[offset : offset+lenOutputCoin])
			if err1 != nil {
				Logger.Log.Errorf("output coin setbytes error: %v\n", err1)
				return errhandler.NewPrivacyErr(errhandler.SetBytesProofErr, err1)
			}
		}
		offset += lenOutputCoin
	}

	return nil
}
