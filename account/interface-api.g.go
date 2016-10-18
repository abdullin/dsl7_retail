package account

import "github.com/abdullin/omni/core/api"

// abdullin: generated from DSLv7

type Handler interface {
	RegisterNew(req *RegisterNewRequest) api.Response
}
