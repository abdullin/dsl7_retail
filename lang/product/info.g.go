package product

import (
	"github.com/abdullin/dsl7/lang/media"
	"github.com/abdullin/dsl7/lang/product"
	"github.com/abdullin/dsl7/lang/size"
	"github.com/abdullin/dsl7/lang/tag"
	"github.com/abdullin/dsl7/lang/weight"
)

// abdullin: generated from DSLv7

// Info - product data
type Info struct {
	ProductId           uint64                 `json:"productId"`
	ProductSku          string                 `json:"productSku"`
	TaxonomyId          uint64                 `json:"taxonomyId"`
	ProductDescription  string                 `json:"productDescription"`
	MediaPictures       []media.Picture        `json:"mediaPictures"`
	SupplierId          uint64                 `json:"supplierId"`
	ProductAttribs      []product.Attrib       `json:"productAttribs"`
	ProductReorderPoint uint32                 `json:"productReorderPoint"`
	SizeBox             size.Box               `json:"sizeBox"`
	WeightMass          weight.Mass            `json:"weightMass"`
	TagSet              map[tag.Value]struct{} `json:"tagSet"`
}

func (r *Info) validate() (err error) {
	return nil
}
func NewInfo(productId uint64, productSku string, taxonomyId uint64, productDescription string, mediaPictures []media.Picture, supplierId uint64, productAttribs []product.Attrib, productReorderPoint uint32, sizeBox size.Box, weightMass weight.Mass, tagSet map[tag.Value]struct{}) *Info {
	return &Info{productId, productSku, taxonomyId, productDescription, mediaPictures, supplierId, productAttribs, productReorderPoint, sizeBox, weightMass, tagSet}
}
