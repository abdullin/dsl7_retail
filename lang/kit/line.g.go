package kit

import "github.com/abdullin/dsl7/lang/product"

// abdullin: generated from DSLv7

// Line - An entry in the kit
type Line struct {
	Quantity    uint64                  `json:"quantity"`
	Description string                  `json:"description"`
	ProductIds  map[product.Id]struct{} `json:"productIds"`
}

func (r *Line) validate() (err error) {
	return nil
}
func NewLine(quantity uint64, description string, productIds map[product.Id]struct{}) *Line {
	return &Line{quantity, description, productIds}
}
