package repos

import (
	"fmt"
	"privacy/x/privacy/common"
	"privacy/x/privacy/repos/bulletproofs"
	"privacy/x/privacy/repos/coin"

	proto "github.com/gogo/protobuf/proto"
)

type PaymentProof struct {
	AggregatedRangeProof *bulletproofs.AggregatedRangeProof
	InputCoins           []*coin.Coin
	OutputCoins          []*coin.Coin
}

func NewPaymentProof() *PaymentProof {
	proof := NewPaymentProof()
	proof.AggregatedRangeProof = bulletproofs.NewAggregatedRangeProof()
	proof.InputCoins = []*coin.Coin{}
	proof.OutputCoins = []*coin.Coin{}
	return proof
}

// Bytes does byte serialization for this payment proof
func (proof PaymentProof) Bytes() []byte {
	var bytes []byte

	comOutputMultiRangeProof := proof.AggregatedRangeProof.Bytes()
	var rangeProofLength uint32 = uint32(len(comOutputMultiRangeProof))
	bytes = append(bytes, common.Uint32ToBytes(rangeProofLength)...)
	bytes = append(bytes, comOutputMultiRangeProof...)

	// InputCoins
	bytes = append(bytes, byte(len(proof.InputCoins)))
	for i := 0; i < len(proof.InputCoins); i++ {
		inputCoins := proof.InputCoins[i].Bytes()
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
	bytes = append(bytes, byte(len(proof.OutputCoins)))
	for i := 0; i < len(proof.OutputCoins); i++ {
		outputCoins := proof.OutputCoins[i].Bytes()
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
func (proof *PaymentProof) SetBytes(proofbytes []byte) error {
	if len(proofbytes) == 0 {
		return fmt.Errorf("Proof bytes is zero")
	}
	offset := 0

	// ComOutputMultiRangeProofSize *aggregatedRangeProof
	if offset+common.Uint32Size >= len(proofbytes) {
		return fmt.Errorf("Out of range aggregated range proof")
	}
	lenComOutputMultiRangeUint32, _ := common.BytesToUint32(proofbytes[offset : offset+common.Uint32Size])
	lenComOutputMultiRangeProof := int(lenComOutputMultiRangeUint32)
	offset += common.Uint32Size

	if offset+lenComOutputMultiRangeProof > len(proofbytes) {
		return fmt.Errorf("Out of range aggregated range proof")
	}
	if lenComOutputMultiRangeProof > 0 {
		bulletproof := bulletproofs.NewAggregatedRangeProof()
		proof.AggregatedRangeProof = bulletproof
		err := proof.AggregatedRangeProof.SetBytes(proofbytes[offset : offset+lenComOutputMultiRangeProof])
		if err != nil {
			return err
		}
		offset += lenComOutputMultiRangeProof
	}

	if offset >= len(proofbytes) {
		return fmt.Errorf("Out of range input coins")
	}
	lenInputCoinsArray := int(proofbytes[offset])
	offset++
	proof.InputCoins = make([]*coin.Coin, lenInputCoinsArray)
	var err error
	for i := 0; i < lenInputCoinsArray; i++ {
		// try get 1-byte for len
		if offset >= len(proofbytes) {
			return fmt.Errorf("Out of range input coins")
		}
		lenInputCoin := int(proofbytes[offset])
		offset++

		if offset+lenInputCoin > len(proofbytes) {
			return fmt.Errorf("Out of range input coins")
		}
		proof.InputCoins[i], err = coin.NewCoinFromBytes(proofbytes[offset : offset+lenInputCoin])
		if err != nil {
			// 1-byte is wrong
			// try get 2-byte for len
			if offset+1 > len(proofbytes) {
				return fmt.Errorf("Out of range input coins")
			}
			lenInputCoin = common.BytesToInt(proofbytes[offset-1 : offset+1])
			offset++

			if offset+lenInputCoin > len(proofbytes) {
				return fmt.Errorf("Out of range input coins")
			}
			proof.InputCoins[i], err = coin.NewCoinFromBytes(proofbytes[offset : offset+lenInputCoin])
			if err != nil {
				return err
			}
		}
		offset += lenInputCoin
	}

	if offset >= len(proofbytes) {
		return fmt.Errorf("Out of range output coins")
	}
	lenOutputCoinsArray := int(proofbytes[offset])
	offset++
	proof.OutputCoins = make([]*coin.Coin, lenOutputCoinsArray)
	for i := 0; i < lenOutputCoinsArray; i++ {
		proof.OutputCoins[i] = new(coin.Coin)
		// try get 1-byte for len
		if offset >= len(proofbytes) {
			return fmt.Errorf("Out of range output coins")
		}
		lenOutputCoin := int(proofbytes[offset])
		offset++

		if offset+lenOutputCoin > len(proofbytes) {
			return fmt.Errorf("Out of range output coins")
		}
		err := proof.OutputCoins[i].SetBytes(proofbytes[offset : offset+lenOutputCoin])
		if err != nil {
			// 1-byte is wrong
			// try get 2-byte for len
			if offset+1 > len(proofbytes) {
				return fmt.Errorf("Out of range output coins")
			}
			lenOutputCoin = common.BytesToInt(proofbytes[offset-1 : offset+1])
			offset++

			if offset+lenOutputCoin > len(proofbytes) {
				return fmt.Errorf("Out of range output coins")
			}
			e := proof.OutputCoins[i].SetBytes(proofbytes[offset : offset+lenOutputCoin])
			if e != nil {
				return e
			}
		}
		offset += lenOutputCoin
	}

	return nil
}

func (p *PaymentProof) Reset()         { p = NewPaymentProof() }
func (p *PaymentProof) String() string { return proto.CompactTextString(p) }
func (p *PaymentProof) ProtoMessage()  {}
