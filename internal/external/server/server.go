package server

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/konu96/Nolack/internal/interfaces"
	"log"
	"net/http"
	"os"
)

var PORT = ":8080"

type Server struct {
	controller interfaces.Controller
}

func NewServer(controller interfaces.Controller) Server {
	return Server{
		controller: controller,
	}
}

func (s *Server) Run() error {
	env := os.Getenv("GO_ENV")
	if err := godotenv.Load(fmt.Sprintf("./%s.env", env)); err != nil {
		panic("not found .env file")
	}

	s.register()

	log.Println("[INFO] Server listening " + PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		return err
	}
	return nil
}

func (s *Server) register() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := s.controller.Exec(w, r); err != nil {
			log.Println(err)
		}
	})
}
