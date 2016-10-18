package user

import (
	"github.com/abdullin/dsl7/lang/author"
	"github.com/abdullin/omni/core/api"
)

// abdullin: generated from DSLv7

// AddRequest - adds a new user into the account
type AddRequest struct {
	AuthorInfo author.Info `json:"authorInfo"`
	UserId     uint64      `json:"userId"`
	UserName   string      `json:"userName"`
	UserEmail  string      `json:"userEmail"`
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
