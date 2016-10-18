package user

import (
	"github.com/abdullin/dsl7/lang/author"
	"github.com/abdullin/omni/core/api"
)

// abdullin: generated from DSLv7

// DeleteRequest - drops a user from this account
type DeleteRequest struct {
	AuthorInfo author.Info `json:"authorInfo"`
	UserId     uint64      `json:"userId"`
}

func (r *DeleteRequest) validate() (err error) {
	return nil
}
func deleteRequest(h Handler, req *api.Request) api.Response {
	var request DeleteRequest

	if err := req.ParseBody(&request); err != nil {
		return api.BadRequestResponse(err.Error())
	}
	if err := request.validate(); err != nil {
		return api.BadRequestResponse(err.Error())
	}
	return h.Delete(&request)
}
