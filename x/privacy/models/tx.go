package models

import (
	"privacy/x/privacy/common"
	"privacy/x/privacy/common/base58"
	"privacy/x/privacy/repos"
	"privacy/x/privacy/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/protobuf/proto"
)

func FetchDataFromTx(ctx sdk.Context, msg *types.MsgCreateTx) ([]types.SerialNumber, map[string][]types.Commitment, map[string][]types.OutputCoin, error) {

	txPrivacyData := types.TxPrivacyData{}

	err := proto.Unmarshal(msg.Value, &txPrivacyData)
	if err != nil {
		ctx.Logger().Error("fail to unmarshal by protobuf tx privacy data %v", err)
		return nil, nil, nil, err
	}

	proof := repos.NewPaymentProof()

	err = proto.Unmarshal(txPrivacyData.Proof, proof)
	if err != nil {
		ctx.Logger().Error("fail to unmarshal by protobuf proof %v", err)
		return nil, nil, nil, err
	}

	serialNumbers, commitments, outCoins, err := parseDataFromProof(proof)
	if err != nil {
		ctx.Logger().Error("fail to parse info from privacy data %v", err)
		return nil, nil, nil, err
	}

	return serialNumbers, commitments, outCoins, nil
}

func parseDataFromProof(proof *repos.PaymentProof) ([]types.SerialNumber, map[string][]types.Commitment, map[string][]types.OutputCoin, error) {
	acceptedSerialNumbers := make([]types.SerialNumber, 0)
	acceptedCommitments := make(map[string][]types.Commitment)
	acceptedOutputcoins := make(map[string][]types.OutputCoin)

	inputCoins := proof.InputCoins
	for _, item := range inputCoins {
		isConfidentialAsset := item.AssetTag != nil
		serialNum := item.GetKeyImage().ToBytesS()
		hash, err := common.NewHashFromBytes(append([]byte{common.BoolToByte(isConfidentialAsset)}, serialNum...))
		if err != nil {
			return nil, nil, nil, err
		}
		serialNumber := types.SerialNumber{
			Value:               serialNum,
			Index:               hash.String(),
			IsConfidentialAsset: isConfidentialAsset,
		}
		acceptedSerialNumbers = append(acceptedSerialNumbers, serialNumber)
	}
	outputCoins := proof.OutputCoins
	for _, item := range outputCoins {
		isConfidentialAsset := item.AssetTag != nil
		pubkey := item.GetPublicKey().ToBytesS()
		publicKeyStr := base58.Base58Check{}.Encode(pubkey, common.ZeroByte)
		if acceptedCommitments[publicKeyStr] == nil {
			acceptedCommitments[publicKeyStr] = make([]types.Commitment, 0)
		}
		c := item.GetCommitment().ToBytesS()
		hash, err := common.NewHashFromBytes(append([]byte{common.BoolToByte(isConfidentialAsset)}, c...))
		if err != nil {
			return nil, nil, nil, err
		}
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
		outputCoinBytes, err := proto.Marshal(item)
		if err != nil {
			return nil, nil, nil, err
		}
		temp := append([]byte{common.BoolToByte(isConfidentialAsset)}, pubkey...)
		hash, err = common.NewHashFromBytes(append(temp, outputCoinBytes...))
		outputCoin := types.OutputCoin{
			Index:               hash.String(),
			IsConfidentialAsset: isConfidentialAsset,
			Value:               outputCoinBytes,
		}
		acceptedOutputcoins[publicKeyStr] = append(acceptedOutputcoins[publicKeyStr], outputCoin)
	}

	return acceptedSerialNumbers, acceptedCommitments, acceptedOutputcoins, nil
}
