package product

import (
	"github.com/abdullin/omni/core/api"
	"github.com/abdullin/omni/core/env"
)

type Module struct{}

func NewModule(pub env.Publisher) *Module { return &Module{} }

func (m *Module) Register(e env.Registrar) {
	registerRequests(m, e)
}

func (m *Module) Add(req *AddRequest) api.Response {
	return api.NewJSON(1)
}
