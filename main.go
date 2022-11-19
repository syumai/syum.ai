package main

import (
	"github.com/syumai/syum.ai/server"
	"github.com/syumai/workers"
)

func main() {
	workers.Serve(server.NewHandler())
}
