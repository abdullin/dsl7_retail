package taxonomy

import "github.com/abdullin/dsl7/lang/attrib"

// abdullin: generated from DSLv7

// Info - information about the taxonomy
type Info struct {
	TaxonomyId   uint64        `json:"taxonomyId"`
	TaxonomyName string        `json:"taxonomyName"`
	AttribList   []attrib.Info `json:"attribList"`
}

func (r *Info) validate() (err error) {
	return nil
}
func NewInfo(taxonomyId uint64, taxonomyName string, attribList []attrib.Info) *Info {
	return &Info{taxonomyId, taxonomyName, attribList}
}
