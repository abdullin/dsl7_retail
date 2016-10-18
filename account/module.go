package account

import (
	"github.com/abdullin/dsl7_retail/lang/author"
	"github.com/abdullin/omni/core/api"
	"github.com/abdullin/omni/core/env"
)

type Module struct{}

func NewModule(pub env.Publisher) *Module { return &Module{} }

func (m *Module) Register(e env.Registrar) {
	registerRequests(m, e)
}
func (m *Module) RegisterNew(req *RegisterNewRequest) api.Response {
	return api.NewJSON(NewRegistered(*author.NewInfo(1, 2), req.AccountName))
}
