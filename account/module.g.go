package account

import (
	"github.com/abdullin/omni/core/api"
	"github.com/abdullin/omni/core/env"
)

// abdullin: generated from DSLv7

// type Module struct { }
// func NewModule() *Module {return &Module{} }
// func (m *Module) Register(r env.Registrar) {}
func registerRequests(m *Module, e env.Registrar) {
	e.HandleHttp("POST", "/account/register-new", func(r *api.Request) api.Response { return registerNewRequest(m, r) })
}
