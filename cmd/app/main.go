package main

import (
	"fmt"
	"log"
	"os"

	"github.com/syumai/syum.ai/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = fmt.Sprintf("localhost:%s", port)
	}

	log.Printf("Listening on port: %s, host: %s", port, host)

	s := server.NewServer(port, host)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	defer s.Close()
}
