package models

import (
	"fmt"
	"math/big"
	"privacy/x/privacy/common"
	"privacy/x/privacy/repos"
	"privacy/x/privacy/repos/bulletproofs"
	"privacy/x/privacy/repos/coin"
	"privacy/x/privacy/repos/key"
	"privacy/x/privacy/repos/mlsag"
	"privacy/x/privacy/repos/operation"
	"privacy/x/privacy/types"

	"github.com/golang/protobuf/proto"
)

func BuildTransferTx(
	outcoins []types.OutputCoin, keySet key.KeySet, msg *types.MsgTransfer,
	lenOTACoins big.Int, mCoins map[string]types.OutputCoin, lCoins []types.OutputCoin,
) (*types.MsgCreateTx, error) {
	var amount uint64
	var err error
	for _, paymentInfo := range msg.PaymentInfos {
		amount, err = common.AddUint64(amount, paymentInfo.Amount)
		if err != nil {
			return nil, err
		}
	}
	coins, err := chooseCoinsByKeySet(outcoins, keySet, amount)
	if err != nil {
		return nil, err
	}
	return buildTransferTx(coins, keySet, msg, lenOTACoins, mCoins, lCoins)
}

func buildTransferTx(
	inputCoins []*coin.Coin, keySet key.KeySet, msg *types.MsgTransfer,
	lenOTACoins big.Int, mCoins map[string]types.OutputCoin, lCoins []types.OutputCoin,
) (*types.MsgCreateTx, error) {
	paymentInfos := make([]*key.PaymentInfo, len(msg.PaymentInfos))
	for index, info := range msg.PaymentInfos {
		paymentAddress := key.PaymentAddress{}
		if err := paymentAddress.FromString(info.PaymentAddress); err != nil {
			return nil, err
		}
		paymentInfos[index] = &key.PaymentInfo{
			Amount:         info.Amount,
			Message:        info.Info,
			PaymentAddress: paymentAddress,
		}
	}
	proof, outputCoins, err := Prove(inputCoins, paymentInfos)
	if err != nil {
		return nil, err
	}

	txPrivacyData, err := SignOnMessage(inputCoins, outputCoins, &keySet.PrivateKey, nil, lenOTACoins, mCoins, lCoins)
	if err != nil {
		return nil, err
	}
	txPrivacyData.Proof = proof.Bytes()

	txPrivacyDataBytes, err := proto.Marshal(txPrivacyData)
	if err != nil {
		return nil, err
	}
	res := &types.MsgCreateTx{
		Value: txPrivacyDataBytes,
	}

	return res, nil
}

func Prove(
	inputCoins []*coin.Coin,
	paymentInfos []*key.PaymentInfo,
) (*repos.PaymentProof, []*coin.Coin, error) {
	outputCoins, err := GenerateOutputCoinsByPaymentInfos(paymentInfos)
	if err != nil {
		return nil, nil, err
	}
	proof, err := prove(inputCoins, outputCoins, nil, false, paymentInfos)
	if err != nil {
		return nil, nil, err
	}

	return proof, outputCoins, nil
}

func prove(
	inputCoins, outputCoins []*coin.Coin,
	sharedSecrets []*operation.Point,
	hasConfidentialAsset bool,
	paymentInfos []*key.PaymentInfo,
) (*repos.PaymentProof, error) {
	proof := repos.NewPaymentProof()
	err := proof.SetInputCoins(inputCoins)
	if err != nil {
		return nil, err
	}
	err = proof.SetOutputCoins(outputCoins)
	if err != nil {
		return nil, err
	}

	// Prepare range proofs
	n := len(outputCoins)
	outputValues := make([]uint64, n)
	outputRands := make([]*operation.Scalar, n)
	for i := 0; i < n; i++ {
		outputValues[i] = outputCoins[i].GetValue()
		outputRands[i] = outputCoins[i].GetRandomness()
	}

	wit := new(bulletproofs.AggregatedRangeWitness)
	wit.Set(outputValues, outputRands)
	if hasConfidentialAsset {
		/*blinders := make([]*operation.Scalar, len(sharedSecrets))*/
		/*for i := range sharedSecrets {*/
		/*if sharedSecrets[i] == nil {*/
		/*blinders[i] = new(operation.Scalar).FromUint64(0)*/
		/*} else {*/
		/*blinders[i], err = coin.ComputeAssetTagBlinder(sharedSecrets[i])*/
		/*if err != nil {*/
		/*return nil, err*/
		/*}*/
		/*}*/
		/*}*/
		/*var err error*/
		/*wit, err = bulletproofs.TransformWitnessToCAWitness(wit, blinders)*/
		/*if err != nil {*/
		/*return nil, err*/
		/*}*/

		/*theBase, err := bulletproofs.GetFirstAssetTag(outputCoins)*/
		/*if err != nil {*/
		/*return nil, err*/
		/*}*/
		/*proof.aggregatedRangeProof, err = wit.ProveUsingBase(theBase)*/

		/*outputCommitments := make([]*operation.Point, n)*/
		/*for i := 0; i < n; i++ {*/
		/*com, err := outputCoins[i].ComputeCommitmentCA()*/
		/*if err != nil {*/
		/*return nil, err*/
		/*}*/
		/*outputCommitments[i] = com*/
		/*}*/
		/*proof.aggregatedRangeProof.SetCommitments(outputCommitments)*/
		/*if err != nil {*/
		/*return nil, err*/
		/*}*/
	} else {
		aggregatedRangeProof, err := wit.Prove()
		if err != nil {
			return nil, err
		}
		proof.SetAggregatedRangeProof(aggregatedRangeProof)
	}

	// After Prove, we should hide all information in coin details.
	for i, outputCoin := range proof.OutputCoins() {
		//if !common.IsPublicKeyBurningAddress(outputCoin.GetPublicKey().ToBytesS()) {
		if err = outputCoin.ConcealOutputCoin(paymentInfos[i].PaymentAddress.GetPublicView()); err != nil {
			return nil, err
		}

		// OutputCoin.GetKeyImage should be nil even though we do not have it
		// Because otherwise the RPC server will return the Bytes of [1 0 0 0 0 ...] (the default byte)
		outputCoins[i].SetKeyImage(nil)
		proof.SetOutputCoinAtIndex(i, &outputCoin)
		//}
	}

	for i, inputCoin := range proof.InputCoins() {
		inputCoin.ConcealInputCoin()
		proof.SetInputCoinAtIndex(i, &inputCoin)
	}

	return proof, nil
}

func SignOnMessage(
	inputCoins, outputCoins []*coin.Coin,
	privateKey *key.PrivateKey, hashedMessage []byte,
	lenOTACoins big.Int, dbCoins map[string]types.OutputCoin, lCoins []types.OutputCoin,
) (*types.TxPrivacyData, error) {
	tx := new(types.TxPrivacyData)

	// Generate Ring
	piBig, err := common.RandBigIntMaxRange(big.NewInt(int64(RingSize)))
	if err != nil {
		return nil, err
	}
	var pi int = int(piBig.Int64())
	ring, indexes, commitmentToZero, err := generateMlsagRingWithIndexes(inputCoins, outputCoins, pi, lenOTACoins, dbCoins, lCoins)
	if err != nil {
		return nil, err
	}

	// Set SigPubKey
	txSigPubKey := new(SigPubKey)
	txSigPubKey.Indexes = indexes
	tx.SigPubKey, err = txSigPubKey.Bytes()
	if err != nil {
		return nil, err
	}

	// Set sigPrivKey
	privKeysMlsag, err := createPrivKeyMlsag(inputCoins, outputCoins, privateKey, commitmentToZero)
	if err != nil {
		return nil, err
	}
	sag := mlsag.NewMlsag(privKeysMlsag, ring, pi)

	// Set Signature
	mlsagSignature, err := sag.Sign(hashedMessage)
	if err != nil {
		return nil, err
	}

	// inputCoins already hold keyImage so set to nil to reduce size
	mlsagSignature.SetKeyImages(nil)
	tx.Sig, err = mlsagSignature.ToBytes()

	return &types.TxPrivacyData{}, nil
}

func generateMlsagRingWithIndexes(
	inputCoins, outputCoins []*coin.Coin,
	pi int,
	lenOTACoins big.Int, mCoins map[string]types.OutputCoin, lCoins []types.OutputCoin,
) (*mlsag.Ring, [][]*big.Int, *operation.Point, error) {
	outputCoinsAsGeneric := make([]*coin.Coin, len(outputCoins))
	for i := 0; i < len(outputCoins); i++ {
		outputCoinsAsGeneric[i] = outputCoins[i]
	}
	var fee uint64
	sumOutputsWithFee := CalculateSumOutputsWithFee(outputCoinsAsGeneric, fee)
	indexes := make([][]*big.Int, RingSize)
	ring := make([][]*operation.Point, RingSize)
	var commitmentToZero *operation.Point
	attempts := 0
	for i := 0; i < RingSize; i++ {
		sumInputs := new(operation.Point).Identity()
		sumInputs.Sub(sumInputs, sumOutputsWithFee)

		row := make([]*operation.Point, len(inputCoins))
		rowIndexes := make([]*big.Int, len(inputCoins))
		if i == pi {
			for j := 0; j < len(inputCoins); j++ {
				row[j] = inputCoins[j].GetPublicKey()
				isConfidentialAsset := inputCoins[j].AssetTag != nil
				outputCoinBytes := inputCoins[j].Bytes()
				temp := append([]byte{common.BoolToByte(isConfidentialAsset)}, row[j].ToBytesS()...)
				hash := common.HashH(append(temp, outputCoinBytes...))
				rowIndexes[j] = big.NewInt(0).SetBytes(mCoins[hash.String()].SerialNumber)
				sumInputs.Add(sumInputs, inputCoins[j].GetCommitment())
			}
		} else {
			for j := 0; j < len(inputCoins); j++ {
				coinDB := new(coin.Coin)
				for attempts < coin.MaxAttempts { // The chance of infinite loop is negligible
					rowIndexes[j], _ = common.RandBigIntMaxRange(&lenOTACoins)
					c := lCoins[rowIndexes[j].Uint64()]
					if err := coinDB.SetBytes(c.Value); err != nil {
						return nil, nil, nil, err
					}

					/*// we do not use burned coins since they will reduce the privacy level of the transaction.*/
					/*if !common.IsPublicKeyBurningAddress(coinDB.GetPublicKey().ToBytesS()) {*/
					/*break*/
					/*}*/
					attempts++
				}
				if attempts == coin.MaxAttempts {
					return nil, nil, nil, fmt.Errorf("cannot form decoys")
				}

				row[j] = coinDB.GetPublicKey()
				sumInputs.Add(sumInputs, coinDB.GetCommitment())
			}
		}
		row = append(row, sumInputs)
		if i == pi {
			commitmentToZero = sumInputs
		}
		ring[i] = row
		indexes[i] = rowIndexes
	}
	return mlsag.NewRing(ring), indexes, commitmentToZero, nil
}

func createPrivKeyMlsag(
	inputCoins, outputCoins []*coin.Coin,
	privateKey *key.PrivateKey, commitmentToZero *operation.Point,
) ([]*operation.Scalar, error) {
	sumRand := new(operation.Scalar).FromUint64(0)
	for _, in := range inputCoins {
		sumRand.Add(sumRand, in.GetRandomness())
	}
	for _, out := range outputCoins {
		sumRand.Sub(sumRand, out.GetRandomness())
	}

	privKeyMlsag := make([]*operation.Scalar, len(inputCoins)+1)
	for i := 0; i < len(inputCoins); i++ {
		var err error
		privKeyMlsag[i], err = inputCoins[i].ParsePrivateKeyOfCoin(*privateKey)
		if err != nil {
			return nil, err
		}
	}
	commitmentToZeroRecomputed := new(operation.Point).ScalarMult(operation.PedCom.G[operation.PedersenRandomnessIndex], sumRand)
	match := operation.IsPointEqual(commitmentToZeroRecomputed, commitmentToZero)
	if !match {
		return nil, fmt.Errorf("error : asset tag sum or commitment sum mismatch")
	}
	privKeyMlsag[len(inputCoins)] = sumRand
	return privKeyMlsag, nil
}
