package tag

import "github.com/abdullin/dsl7_retail/lang/tag"

// abdullin: generated from DSLv7

type Set map[tag.Value]struct{}

func (r *Set) validate() (err error) {
	return nil
}

type Value string

func (r *Value) validate() (err error) {
	return nil
}
