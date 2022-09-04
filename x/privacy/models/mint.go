package models

import (
	"math/big"
	"privacy/x/privacy/common"
	"privacy/x/privacy/repos"
	"privacy/x/privacy/repos/coin"
	"privacy/x/privacy/repos/key"
	"privacy/x/privacy/repos/operation"
	"privacy/x/privacy/repos/schnorr"
	"privacy/x/privacy/types"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/protobuf/proto"
)

func BuildMintTx(ctx sdk.Context, privateKey *key.PrivateKey, otaReceiver coin.OTAReceiver, amount big.Int, info []byte, hashedMessage common.Hash) (*types.MsgCreateTx, error) {
	outputCoin, err := GenerateOutputCoin(amount, info, otaReceiver)
	if err != nil {
		return nil, err
	}

	proof := repos.NewPaymentProof()
	proof.SetOutputCoins([]*coin.Coin{outputCoin})

	sig, sigPubKey, err := SignNoPrivacy(privateKey, hashedMessage.Bytes())
	if err != nil {
		return nil, err
	}

	txPrivacyData := &types.TxPrivacyData{
		Proof:     proof.Bytes(),
		SigPubKey: sigPubKey,
		Sig:       sig,
		LockTime:  uint64(time.Now().Unix()),
		Info:      info,
	}
	txPrivacyDataBytes, err := proto.Marshal(txPrivacyData)
	if err != nil {
		return nil, err
	}
	res := &types.MsgCreateTx{
		Value: txPrivacyDataBytes,
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
