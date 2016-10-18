package user

import (
	"github.com/abdullin/omni/core/api"
	"github.com/abdullin/omni/core/env"
)

// abdullin: generated from DSLv7

// type Module struct { }
// func NewModule() *Module {return &Module{} }
// func (m *Module) Register(r env.Registrar) {}
func registerRequests(m *Module, e env.Registrar) {
	e.HandleHttp("POST", "/user/add", func(r *api.Request) api.Response { return addRequest(m, r) })
	e.HandleHttp("POST", "/user/delete", func(r *api.Request) api.Response { return deleteRequest(m, r) })
	e.HandleHttp("POST", "/user/rename", func(r *api.Request) api.Response { return renameRequest(m, r) })
	e.HandleHttp("POST", "/user/detail", func(r *api.Request) api.Response { return detailRequest(m, r) })
}
