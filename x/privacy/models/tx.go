package models

import (
	"math/big"
	"privacy/x/privacy/common"
	"privacy/x/privacy/common/base58"
	"privacy/x/privacy/repos"
	"privacy/x/privacy/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/protobuf/proto"
)

func FetchDataFromTx(ctx sdk.Context, txData []byte, outputCoinLength big.Int) (
	[]types.SerialNumber,
	map[string][]types.Commitment,
	map[string][]types.OutputCoin,
	map[string][]types.OTACoin,
	*big.Int,
	error,
) {
	txPrivacyData := types.TxPrivacyData{}
	err := proto.Unmarshal(txData, &txPrivacyData)
	if err != nil {
		ctx.Logger().Error("[incognito] fail to unmarshal by protobuf tx privacy data:")
		ctx.Logger().Error(err.Error())
		return nil, nil, nil, nil, nil, err
	}
	proof := repos.NewPaymentProof()
	err = proof.SetBytes(txPrivacyData.Proof)
	if err != nil {
		ctx.Logger().Error("[incognito] fail to unmarshal by protobuf proof:")
		ctx.Logger().Error(err.Error())
		return nil, nil, nil, nil, nil, err
	}

	serialNumbers, commitments, outCoins, otaCoins, newOutputCoinLength, err := parseDataFromProof(proof, outputCoinLength)
	if err != nil {
		ctx.Logger().Error("fail to parse info from privacy datav")
		ctx.Logger().Error(err.Error())
		return nil, nil, nil, nil, nil, err
	}

	return serialNumbers, commitments, outCoins, otaCoins, newOutputCoinLength, nil
}

func parseDataFromProof(proof *repos.PaymentProof, outputCoinLength big.Int) (
	[]types.SerialNumber,
	map[string][]types.Commitment,
	map[string][]types.OutputCoin,
	map[string][]types.OTACoin,
	*big.Int,
	error,
) {
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
		outputCoinLength.Add(&outputCoinLength, big.NewInt(1))
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
			SerialNumber:        outputCoinLength.Bytes(),
			IsConfidentialAsset: isConfidentialAsset,
			PubKey:              pubkey,
			Value:               outputCoinBytes,
		}
		acceptedOutputcoins[publicKeyStr] = append(acceptedOutputcoins[publicKeyStr], outputCoin)

		otaCoin := types.OTACoin{
			Index:           hash.String(),
			OutputCoinIndex: hash.String(),
		}
		acceptedOTACoins[publicKeyStr] = append(acceptedOTACoins[publicKeyStr], otaCoin)
	}

	return acceptedSerialNumbers, acceptedCommitments, acceptedOutputcoins, acceptedOTACoins, &outputCoinLength, nil
}
