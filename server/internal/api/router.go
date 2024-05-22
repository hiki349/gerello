package api

import (
	"net/http"
)

func (s *RestServer) createProjectsRouter(mux *http.ServeMux) {
	// api/projects/
	mux.HandleFunc("GET projects/", s.fetchAll)

	// api/projects/:id
	mux.HandleFunc("GET projects/{id}", s.fetchByID)

	// api/projects/:id
	mux.HandleFunc("PUT projects/{id}", s.update)

	// api/projects/:id
	mux.HandleFunc("DELETE projects/{id}", s.delete)

	// api/projects/
	mux.HandleFunc("POST projects/", s.create)
}
