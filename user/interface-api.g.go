package user

import "github.com/abdullin/omni/core/api"

// abdullin: generated from DSLv7

type Handler interface {
	Add(req *AddRequest) api.Response
	Delete(req *DeleteRequest) api.Response
	Rename(req *RenameRequest) api.Response
	Detail(req *DetailRequest) api.Response
}
