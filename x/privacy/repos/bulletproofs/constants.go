package bulletproofs

import "privacy/x/privacy/common"

const (
	precompPedGValIndex = iota
	precompPedGRandIndex
	precompUIndex
	precompGIndex
)

const (
	aggParamNMax = common.MaxOutputCoin * common.MaxExp
)

func precompHIndex(paramN int) int {
	return precompGIndex + paramN
}
