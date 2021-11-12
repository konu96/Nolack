package main

import (
	"github.com/konu96/Nolack/internal/external/server"
	"log"
)

func main() {
	s := server.NewServer()

	if err := s.Run(); err != nil {
		log.Fatal(err)
	}
}
