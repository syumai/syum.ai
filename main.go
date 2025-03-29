package main

import (
	spinhttp "github.com/fermyon/spin/sdk/go/v2/http"
	"github.com/syumai/syum.ai/server"
)

func init() {
	spinhttp.Handle(server.NewHandler().ServeHTTP)
}
