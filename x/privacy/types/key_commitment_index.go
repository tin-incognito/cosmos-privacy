package types

import "encoding/binary"

var _ binary.ByteOrder

const (
	// CommitmentIndexKeyPrefix is the prefix to retrieve all CommitmentIndex
	CommitmentIndexKeyPrefix = "CommitmentIndex/value/"
)

// CommitmentIndexKey returns the store key to retrieve a CommitmentIndex from the index fields
func CommitmentIndexKey(
	index string,
) []byte {
	var key []byte

	indexBytes := []byte(index)
	key = append(key, indexBytes...)
	key = append(key, []byte("/")...)

	return key
}
