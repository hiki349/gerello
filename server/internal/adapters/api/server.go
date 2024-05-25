package api

import (
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

	return http.ListenAndServe(":"+s.port, mux)
}
