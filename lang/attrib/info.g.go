package attrib

import "github.com/abdullin/dsl7/lang/attrib"

// abdullin: generated from DSLv7

// Info - product attribs
type Info struct {
	AttribId     uint64                    `json:"attribId"`
	AttribName   string                    `json:"attribName"`
	AttribValues map[attrib.Value]struct{} `json:"attribValues"`
}

func (r *Info) validate() (err error) {
	return nil
}
func NewInfo(attribId uint64, attribName string, attribValues map[attrib.Value]struct{}) *Info {
	return &Info{attribId, attribName, attribValues}
}
