package user

import (
	"github.com/abdullin/dsl7_retail/lang/author"
	"github.com/abdullin/omni/core/api"
)

// abdullin: generated from DSLv7

// RenameRequest - change user's name
type RenameRequest struct {
	AuthorInfo author.Info `json:"authorInfo"`
	UserName   string      `json:"userName"`
}

func (r *RenameRequest) validate() (err error) {
	return nil
}
func renameRequest(h Handler, req *api.Request) api.Response {
	var request RenameRequest

	if err := req.ParseBody(&request); err != nil {
		return api.BadRequestResponse(err.Error())
	}
	if err := request.validate(); err != nil {
		return api.BadRequestResponse(err.Error())
	}
	return h.Rename(&request)
}
