package product

import "github.com/abdullin/dsl7_retail/lang/product"

// abdullin: generated from DSLv7

type Ids map[product.Id]struct{}

func (r *Ids) validate() (err error) {
	return nil
}

type Sku string

func (r *Sku) validate() (err error) {
	return nil
}

type ReorderPoint uint32

func (r *ReorderPoint) validate() (err error) {
	return nil
}

type Id uint64

func (r *Id) validate() (err error) {
	return nil
}

type Attribs []product.Attrib

func (r *Attribs) validate() (err error) {
	return nil
}

type Description string

func (r *Description) validate() (err error) {
	return nil
}
