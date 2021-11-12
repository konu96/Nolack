package server

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/konu96/Nolack/internal/interfaces"
	"log"
	"net/http"
	"os"
)

type Server struct {
	controller interfaces.Controller
}

func NewServer() Server {
	return Server{
		controller: interfaces.NewController(),
	}
}

func (s *Server) Run() error {
	env := os.Getenv("GO_ENV")
	if err := godotenv.Load(fmt.Sprintf("./%s.env", env)); err != nil {
		return fmt.Errorf("%s.env not found", env)
	}

	s.register()

	log.Println("[INFO] Server listening")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		return err
	}
	return nil
}

func (s *Server) register() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		s.controller.Exec(w, r)
	})
}
