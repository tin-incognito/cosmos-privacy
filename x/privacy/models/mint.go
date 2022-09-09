package models

import (
	"privacy/x/privacy/common"
	"privacy/x/privacy/repos"
	"privacy/x/privacy/repos/coin"
	"privacy/x/privacy/repos/key"
	"privacy/x/privacy/repos/operation"
	"privacy/x/privacy/repos/schnorr"
	"privacy/x/privacy/types"
	"time"
)

func BuildMintTx(
	otaReceiver coin.OTAReceiver,
	amount uint64,
	info []byte,
	md []byte,
	hashedMessage common.Hash,
) (*types.MsgPrivacyData, error) {
	privateKey := GeneratePrivateKey()
	outputCoin, err := GenerateOutputCoin(amount, info, otaReceiver)
	if err != nil {
		return nil, err
	}

	proof := repos.NewPaymentProof()
	proof.SetOutputCoins([]*coin.Coin{outputCoin})

	sig, sigPubKey, err := SignNoPrivacy(&privateKey, hashedMessage.Bytes())
	if err != nil {
		return nil, err
	}

	res := &types.MsgPrivacyData{
		Proof:     proof.Bytes(),
		SigPubKey: sigPubKey,
		Sig:       sig,
		LockTime:  uint64(time.Now().Unix()),
		Info:      info,
		TxType:    TxMintType,
		Metadata:  md,
	}
	return res, nil
}

func SignNoPrivacy(privateKey *key.PrivateKey, hashedMessage []byte) (signatureBytes []byte, sigPubKey []byte, err error) {
	/****** using Schnorr signature *******/
	// sign with sigPrivKey
	// prepare private key for Schnorr
	sk := new(operation.Scalar).FromBytesS(*privateKey)
	r := new(operation.Scalar).FromUint64(0)
	sigKey := new(schnorr.SchnorrPrivateKey)
	sigKey.Set(sk, r)
	signature, err := sigKey.Sign(hashedMessage)
	if err != nil {
		return nil, nil, err
	}

	signatureBytes = signature.Bytes()
	sigPubKey = sigKey.GetPublicKey().GetPublicKey().ToBytesS()
	return signatureBytes, sigPubKey, nil
}
