package coin

import (
	"fmt"
	"privacy/x/privacy/handler/operation"
)

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func parsePointForSetBytes(coinBytes *[]byte, offset *int) (*operation.Point, error) {
	b := *coinBytes
	if *offset >= len(b) {
		return nil, fmt.Errorf("offset is larger than len(bytes), cannot parse point")
	}
	var point *operation.Point
	var err error
	lenField := b[*offset]
	*offset++
	if lenField != 0 {
		if *offset+int(lenField) > len(b) {
			return nil, fmt.Errorf("offset+curLen is larger than len(bytes), cannot parse point for set bytes")
		}
		data := b[*offset : *offset+int(lenField)]
		point, err = new(operation.Point).FromBytesS(data)
		if err != nil {
			return nil, err
		}
		*offset += int(lenField)
	}
	return point, nil
}

func parseInfoForSetBytes(coinBytes *[]byte, offset *int) ([]byte, error) {
	b := *coinBytes
	if *offset >= len(b) {
		return []byte{}, fmt.Errorf("offset is larger than len(bytes), cannot parse info")
	}
	info := []byte{}
	lenField := b[*offset]
	*offset++
	if lenField != 0 {
		if *offset+int(lenField) > len(b) {
			return []byte{}, fmt.Errorf("offset+curLen is larger than len(bytes), cannot parse info for set bytes")
		}
		info = make([]byte, lenField)
		copy(info, b[*offset:*offset+int(lenField)])
		*offset += int(lenField)
	}
	return info, nil
}

func parseScalarForSetBytes(coinBytes *[]byte, offset *int) (*operation.Scalar, error) {
	b := *coinBytes
	if *offset >= len(b) {
		return nil, fmt.Errorf("offset is larger than len(bytes), cannot parse scalar")
	}
	var sc *operation.Scalar
	lenField := b[*offset]
	*offset++
	if lenField != 0 {
		if *offset+int(lenField) > len(b) {
			return nil, fmt.Errorf("offset+curLen is larger than len(bytes), cannot parse scalar for set bytes")
		}
		data := b[*offset : *offset+int(lenField)]
		sc = new(operation.Scalar).FromBytesS(data)
		*offset += int(lenField)
	}
	return sc, nil
}
