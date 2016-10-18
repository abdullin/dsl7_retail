package kit

import "github.com/abdullin/dsl7/lang/kit"

// abdullin: generated from DSLv7

type Id uint64

func (r *Id) validate() (err error) {
	return nil
}

type Lines []kit.Line

func (r *Lines) validate() (err error) {
	return nil
}

type Ids map[kit.Id]struct{}

func (r *Ids) validate() (err error) {
	return nil
}

type Sku string

func (r *Sku) validate() (err error) {
	return nil
}
