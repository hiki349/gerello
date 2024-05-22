package api

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"

	"gerello/internal/domain/models"
	response "gerello/pkg"
)

// projects/
func (s *RestServer) fetchAll(w http.ResponseWriter, r *http.Request) {
	projects, err := s.services.FetchAll(r.Context())
	if err != nil {
		response.NewInternalError(w)
		return
	}

	json, err := json.Marshal(projects)
	if err != nil {
		response.NewInternalError(w)
		return
	}

	response.New(w).
		SetStatusCode(http.StatusOK).
		SetBody(json).
		Build()
}

// projects/:id
func (s *RestServer) fetchByID(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r.PathValue("id"))
	if err != nil {
		response.NewInternalError(w)
		return
	}

	project, err := s.services.FetchByID(r.Context(), id)
	if err != nil {
		response.NewInternalError(w)
		return
	}

	json, err := json.Marshal(project)
	if err != nil {
		response.NewInternalError(w)
		return
	}

	response.New(w).
		SetStatusCode(http.StatusOK).
		SetBody(json).
		Build()
}

// projects/:id
func (s *RestServer) update(w http.ResponseWriter, r *http.Request) {
	var projectInput ProjectInput

	id, err := getID(r.PathValue("id"))
	if err != nil {
		response.NewInternalError(w)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&projectInput); err != nil {
		response.NewInternalError(w)
		return
	}

	project, err := s.services.Update(r.Context(), id, (*models.ProjectInput)(&projectInput))
	if err != nil {
		response.NewInternalError(w)
		return
	}

	json, err := json.Marshal(project)
	if err != nil {
		response.NewInternalError(w)
		return
	}

	response.New(w).
		SetStatusCode(http.StatusOK).
		SetBody(json).
		Build()
}

// projects/:id
func (s *RestServer) delete(w http.ResponseWriter, r *http.Request) {
	id, err := getID(r.PathValue("id"))
	if err != nil {
		response.NewInternalError(w)
		return
	}

	_, err = s.services.Delete(r.Context(), id)
	if err != nil {
		response.NewInternalError(w)
		return
	}
	response.New(w).
		SetStatusCode(http.StatusOK).
		Build()
}

// projects/
func (s *RestServer) create(w http.ResponseWriter, r *http.Request) {
	var projectInput ProjectInput

	if err := json.NewDecoder(r.Body).Decode(&projectInput); err != nil {
		response.NewInternalError(w)
		return
	}

	project, err := s.services.Add(r.Context(), (*models.ProjectInput)(&projectInput))
	if err != nil {
		response.NewInternalError(w)
		return
	}

	json, err := json.Marshal(project)
	if err != nil {
		response.NewInternalError(w)
		return
	}

	response.New(w).
		SetStatusCode(http.StatusCreated).
		SetBody(json).
		Build()
}

func getID(pathID string) (uuid.UUID, error) {
	if pathID == "" {
		return uuid.UUID{}, errors.New("")
	}

	id, err := uuid.FromBytes([]byte(pathID))
	if err != nil {
		return uuid.UUID{}, err
	}

	return id, nil
}
