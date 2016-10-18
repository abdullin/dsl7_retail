package product

import (
	"github.com/abdullin/dsl7/lang/author"
	"github.com/abdullin/omni/core/api"
)

// abdullin: generated from DSLv7

// AddRequest - add a new product
type AddRequest struct {
	AuthorInfo author.Info `json:"authorInfo"`
	ProductSku string      `json:"productSku"`
}

func (r *AddRequest) validate() (err error) {
	return nil
}
func addRequest(h Handler, req *api.Request) api.Response {
	var request AddRequest

	if err := req.ParseBody(&request); err != nil {
		return api.BadRequestResponse(err.Error())
	}
	if err := request.validate(); err != nil {
		return api.BadRequestResponse(err.Error())
	}
	return h.Add(&request)
}
