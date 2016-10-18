package size

import "github.com/abdullin/dsl7_retail/lang/size"

// abdullin: generated from DSLv7

// Box - Dimensions
type Box struct {
	SizeLength size.Length `json:"sizeLength"`
	SizeWidth  size.Width  `json:"sizeWidth"`
	SizeHeight size.Height `json:"sizeHeight"`
	SizeScale  size.Scale  `json:"sizeScale"`
}

func (r *Box) validate() (err error) {
	return nil
}
func NewBox(sizeLength size.Length, sizeWidth size.Width, sizeHeight size.Height, sizeScale size.Scale) *Box {
	return &Box{sizeLength, sizeWidth, sizeHeight, sizeScale}
}
