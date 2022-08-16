package bulletproofs

import "privacy/x/privacy/handler/operation"

type InnerProductWitness struct {
	a []*operation.Scalar
	b []*operation.Scalar
	p *operation.Point
}

type InnerProductProof struct {
	l []*operation.Point
	r []*operation.Point
	a *operation.Scalar
	b *operation.Scalar
	p *operation.Point
}
