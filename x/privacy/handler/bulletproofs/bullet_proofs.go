package bulletproofs

import (
	"privacy/x/privacy/handler/operation"
)

// AggregatedRangeWitness contains the prover's secret data (the actual values to be proven & the generated random blinders)
// needed for creating a range proof.
type AggregatedRangeWitness struct {
	values []uint64
	rands  []*operation.Scalar
}

// AggregatedRangeProof is the struct for Bulletproof.
// The statement being proven is that output coins' values are in the uint64 range.
type AggregatedRangeProof struct {
	cmsValue          []*operation.Point
	a                 *operation.Point
	s                 *operation.Point
	t1                *operation.Point
	t2                *operation.Point
	tauX              *operation.Scalar
	tHat              *operation.Scalar
	mu                *operation.Scalar
	innerProductProof *InnerProductProof
}

type bulletproofParams struct {
	g  []*operation.Point
	h  []*operation.Point
	u  *operation.Point
	cs *operation.Point

	precomps []operation.PrecomputedPoint
}
