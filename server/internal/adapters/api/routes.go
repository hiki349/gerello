package api

import (
	"net/http"

	"gerello/internal/adapters/api/projects"
)

func (s *RestServer) initRoutes(mux *http.ServeMux) {
	// api/projects/
	mux.HandleFunc("GET projects/", projects.FetchAll(s.services))

	// api/projects/:id
	mux.HandleFunc("GET projects/{id}", projects.FetchByID(s.services))

	// api/projects/:id
	mux.HandleFunc("PUT projects/{id}", projects.Update(s.services))

	// api/projects/:id
	mux.HandleFunc("DELETE projects/{id}", projects.Delete(s.services))

	// api/projects/
	mux.HandleFunc("POST projects/", projects.Create(s.services))
}
