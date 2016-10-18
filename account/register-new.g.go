package account

import "github.com/abdullin/omni/core/api"

// abdullin: generated from DSLv7

// RegisterNewRequest - register a new account
type RegisterNewRequest struct {
	AccountName string `json:"accountName"`
	UserName    string `json:"userName"`
	UserEmail   string `json:"userEmail"`
}

func (r *RegisterNewRequest) validate() (err error) {
	return nil
}
func registerNewRequest(h Handler, req *api.Request) api.Response {
	var request RegisterNewRequest

	if err := req.ParseBody(&request); err != nil {
		return api.BadRequestResponse(err.Error())
	}
	if err := request.validate(); err != nil {
		return api.BadRequestResponse(err.Error())
	}
	return h.RegisterNew(&request)
}
