package api

import (
	"net/http"

	"gerello/internal/adapters/primary/api/projects"
)

func (s *RestServer) initRoutes(mux *http.ServeMux) {
	// /projects/
	mux.HandleFunc("GET /projects/", projects.FetchAll(s.services))

	// /projects/:id
	mux.HandleFunc("GET /projects/{id}", projects.FetchByID(s.services))

	// /projects/:id
	mux.HandleFunc("PUT /projects/{id}", projects.Update(s.services))

	// /projects/:id
	mux.HandleFunc("DELETE /projects/{id}", projects.Delete(s.services))

	// /projects/
	mux.HandleFunc("POST /projects/", projects.Create(s.services))
}
