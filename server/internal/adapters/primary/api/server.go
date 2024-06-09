package api

import (
	"log"
	"net/http"

	"gerello/internal/core/services"
)

type Server interface {
	Run() error
}

type RestServer struct {
	services services.ProjectService
	port     string
}

func New(services services.ProjectService, port string) *RestServer {
	return &RestServer{
		services: services,
		port:     port,
	}
}

func (s *RestServer) Run() error {
	mux := http.NewServeMux()
	s.initRoutes(mux)

	log.Printf("listen on http://localhost:%s\n", s.port)

	return http.ListenAndServe(":"+s.port, mux)
}
