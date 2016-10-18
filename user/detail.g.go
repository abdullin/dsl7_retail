package user

import (
	"github.com/abdullin/dsl7/lang/author"
	"github.com/abdullin/omni/core/api"
)

// abdullin: generated from DSLv7

// DetailRequest - query user detail given her id
type DetailRequest struct {
	AuthorInfo author.Info `json:"authorInfo"`
	UserId     uint64      `json:"userId"`
}

func (r *DetailRequest) validate() (err error) {
	return nil
}
func detailRequest(h Handler, req *api.Request) api.Response {
	var request DetailRequest

	if err := req.ParseBody(&request); err != nil {
		return api.BadRequestResponse(err.Error())
	}
	if err := request.validate(); err != nil {
		return api.BadRequestResponse(err.Error())
	}
	return h.Detail(&request)
}
