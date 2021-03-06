package main

import (
	"log"
	"os"

	"github.com/syumai/syum.ai/server"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on port %s", port)

	s := server.NewServer(port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	defer s.Close()
}
