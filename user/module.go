package user

import (
	"github.com/abdullin/omni/core/api"
	"github.com/abdullin/omni/core/env"
)

type Module struct{}

func NewModule(pub env.Publisher) *Module { return &Module{} }

func (m *Module) Register(e env.Registrar) {
	registerRequests(m, e)
}

// handlers
func (m *Module) Add(req *AddRequest) api.Response {
	return api.NewJSON(&addUserResponse{
		Name: req.UserName,
	})
}
func (m *Module) Delete(req *DeleteRequest) api.Response {
	return api.NewJSON(&addUserResponse{})
}
func (m *Module) Rename(req *RenameRequest) api.Response {
	return api.NewJSON(&addUserResponse{})
}
func (m *Module) Detail(req *DetailRequest) api.Response {
	return api.NewJSON(&addUserResponse{})
}

// to generate
type addUserResponse struct {
	Name string `json:"name"`
}
