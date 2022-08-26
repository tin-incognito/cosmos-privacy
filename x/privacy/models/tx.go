package models

import (
	"math/big"
	"privacy/x/privacy/common"
	"privacy/x/privacy/common/base58"
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

func FetchDataFromTx(ctx sdk.Context, txData []byte) (
	[]types.SerialNumber, map[string][]types.Commitment, map[string][]types.OutputCoin, map[string][]types.OTACoin, error,
) {
	txPrivacyData := types.TxPrivacyData{}
	err := proto.Unmarshal(txData, &txPrivacyData)
	if err != nil {
		ctx.Logger().Error("[incognito] fail to unmarshal by protobuf tx privacy data:")
		ctx.Logger().Error(err.Error())
		return nil, nil, nil, nil, err
	}
	proof := repos.NewPaymentProof()
	err = proof.SetBytes(txPrivacyData.Proof)
	if err != nil {
		ctx.Logger().Error("[incognito] fail to unmarshal by protobuf proof:")
		ctx.Logger().Error(err.Error())
		return nil, nil, nil, nil, err
	}

	serialNumbers, commitments, outCoins, otaCoins, err := parseDataFromProof(proof)
	if err != nil {
		ctx.Logger().Error("fail to parse info from privacy datav")
		ctx.Logger().Error(err.Error())
		return nil, nil, nil, nil, err
	}

	return serialNumbers, commitments, outCoins, otaCoins, nil
}

func parseDataFromProof(proof *repos.PaymentProof) ([]types.SerialNumber, map[string][]types.Commitment, map[string][]types.OutputCoin, map[string][]types.OTACoin, error) {
	acceptedSerialNumbers := make([]types.SerialNumber, 0)
	acceptedCommitments := make(map[string][]types.Commitment)
	acceptedOutputcoins := make(map[string][]types.OutputCoin)
	acceptedOTACoins := make(map[string][]types.OTACoin)

	inputCoins := proof.InputCoins()
	for _, item := range inputCoins {
		isConfidentialAsset := item.AssetTag != nil
		serialNum := item.GetKeyImage().ToBytesS()
		hash := common.HashH(append([]byte{common.BoolToByte(isConfidentialAsset)}, serialNum...))
		serialNumber := types.SerialNumber{
			Value:               serialNum,
			Index:               hash.String(),
			IsConfidentialAsset: isConfidentialAsset,
		}
		acceptedSerialNumbers = append(acceptedSerialNumbers, serialNumber)
	}
	outputCoins := proof.OutputCoins()
	for _, item := range outputCoins {
		isConfidentialAsset := item.AssetTag != nil
		pubkey := item.GetPublicKey().ToBytesS()
		publicKeyStr := base58.Base58Check{}.Encode(pubkey, common.ZeroByte)
		if acceptedCommitments[publicKeyStr] == nil {
			acceptedCommitments[publicKeyStr] = make([]types.Commitment, 0)
		}
		c := item.GetCommitment().ToBytesS()
		hash := common.HashH(append([]byte{common.BoolToByte(isConfidentialAsset)}, c...))

		// get data for commitments
		commiment := types.Commitment{
			Index:               hash.String(),
			Value:               c,
			IsConfidentialAsset: isConfidentialAsset,
		}
		acceptedCommitments[publicKeyStr] = append(acceptedCommitments[publicKeyStr], commiment)
		// get data for output coin
		if acceptedOutputcoins[publicKeyStr] == nil {
			acceptedOutputcoins[publicKeyStr] = make([]types.OutputCoin, 0)
		}
		outputCoinBytes := item.Bytes()
		temp := append([]byte{common.BoolToByte(isConfidentialAsset)}, pubkey...)
		hash = common.HashH(append(temp, outputCoinBytes...))
		outputCoin := types.OutputCoin{
			Index:               hash.String(),
			IsConfidentialAsset: isConfidentialAsset,
			PubKey:              pubkey,
			Value:               outputCoinBytes,
		}
		acceptedOutputcoins[publicKeyStr] = append(acceptedOutputcoins[publicKeyStr], outputCoin)

		otaCoin := types.OTACoin{
			OutputCoinIndex: hash.String(),
		}
		acceptedOTACoins[publicKeyStr] = append(acceptedOTACoins[publicKeyStr], otaCoin)
	}

	return acceptedSerialNumbers, acceptedCommitments, acceptedOutputcoins, acceptedOTACoins, nil
}

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
