package api

import (
	"log"
	"net/http"

	"gerello/internal/domain/services"
)

type Server interface {
	Start()
}

type RestServer struct {
	services services.ProjectService
}

func New(services services.ProjectService) *RestServer {
	return &RestServer{
		services: services,
	}
}

func (s *RestServer) Start() {
	mux := http.NewServeMux()
	s.createProjectsRouter(mux)

	log.Fatal(http.ListenAndServe(":7070", mux))
}
