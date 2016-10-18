package attrib

import "github.com/abdullin/dsl7/lang/attrib"

// abdullin: generated from DSLv7

type List []attrib.Info

func (r *List) validate() (err error) {
	return nil
}

type Values map[attrib.Value]struct{}

func (r *Values) validate() (err error) {
	return nil
}

type Name string

func (r *Name) validate() (err error) {
	return nil
}

type Value string

func (r *Value) validate() (err error) {
	return nil
}

type Id uint64

func (r *Id) validate() (err error) {
	return nil
}
