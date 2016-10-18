package kit

import "github.com/abdullin/dsl7/lang/kit"

// abdullin: generated from DSLv7

// Info - Product kit - a possible combination of items that we could sell
type Info struct {
	KitSku      string     `json:"kitSku"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	KitLines    []kit.Line `json:"kitLines"`
}

func (r *Info) validate() (err error) {
	return nil
}
func NewInfo(kitSku string, name string, description string, kitLines []kit.Line) *Info {
	return &Info{kitSku, name, description, kitLines}
}
