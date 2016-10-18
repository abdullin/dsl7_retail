package weight

import "github.com/abdullin/dsl7/lang/weight"

// abdullin: generated from DSLv7

// Mass - Weight with a dimension
type Mass struct {
	WeightScale weight.Scale `json:"weightScale"`
	WeightValue float64      `json:"weightValue"`
}

func (r *Mass) validate() (err error) {
	return nil
}
func NewMass(weightScale weight.Scale, weightValue float64) *Mass {
	return &Mass{weightScale, weightValue}
}
