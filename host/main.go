package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/abdullin/omni/core/bus"
	"github.com/abdullin/omni/core/env"
	"github.com/abdullin/omni/core/hosting"

	"github.com/abdullin/dsl7/account"
	"github.com/abdullin/dsl7/product"
	"github.com/abdullin/dsl7/user"

	"github.com/op/go-logging"
)

var l = logging.MustGetLogger("main")

func main() {

	router := mux.NewRouter()
	memBus := bus.NewMem()

	var wrap = bus.WrapWithLogging(memBus)
	modules := factory(wrap)
	var host = hosting.New(modules)

	host.WireHttp(router)
	host.WireHandlers(memBus)

	memBus.Start()

	bind := ":8001"
	l.Info("Listening at %v", bind)
	l.Panic(http.ListenAndServe(bind, router))
}

func factory(pub env.Publisher) []env.Module {
	return []env.Module{
		product.NewModule(pub),
		account.NewModule(pub),
		user.NewModule(pub),
	}
}
