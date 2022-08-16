package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// TxPrivacyDataKeyPrefix is the prefix to retrieve all TxPrivacyData
	TxPrivacyDataKeyPrefix = "TxPrivacyData/value/"
)

// TxPrivacyDataKey returns the store key to retrieve a TxPrivacyData from the index fields
func TxPrivacyDataKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
