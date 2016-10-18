package product

// abdullin: generated from DSLv7

// Attrib - product attrib
type Attrib struct {
	AttribId    uint64 `json:"attribId"`
	AttribName  string `json:"attribName"`
	AttribValue string `json:"attribValue"`
}

func (r *Attrib) validate() (err error) {
	return nil
}
func NewAttrib(attribId uint64, attribName string, attribValue string) *Attrib {
	return &Attrib{attribId, attribName, attribValue}
}
