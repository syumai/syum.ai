package main

import (
	"log"
	"os"

	"github.com/syumai/syum.ai/redirects/redirect"
)

const location = "https://github.com/syumai"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)

	if err := redirect.Serve(location, port); err != nil {
		log.Fatal(err)
	}
}
