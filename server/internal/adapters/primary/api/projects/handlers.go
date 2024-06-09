package projects

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"

	"gerello/internal/core/domain/project"
	"gerello/internal/core/services"
	response "gerello/pkg"
)

func FetchAll(projectsSvc services.ProjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		projects, err := projectsSvc.FetchAll(r.Context())
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
}

func FetchByID(projectsSvc services.ProjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := getID(r.PathValue("id"))
		if err != nil {
			response.NewInternalError(w)
			return
		}

		project, err := projectsSvc.FetchByID(r.Context(), id)
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
}

func Update(projectsSvc services.ProjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		project, err := projectsSvc.Update(r.Context(), id, (*project.ProjectInput)(&projectInput))
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
}

func Delete(projectsSvc services.ProjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := getID(r.PathValue("id"))
		if err != nil {
			response.NewInternalError(w)
			return
		}

		_, err = projectsSvc.Delete(r.Context(), id)
		if err != nil {
			response.NewInternalError(w)
			return
		}
		response.New(w).
			SetStatusCode(http.StatusOK).
			Build()
	}
}

func Create(projectsSvc services.ProjectService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var projectInput ProjectInput

		if err := json.NewDecoder(r.Body).Decode(&projectInput); err != nil {
			response.NewInternalError(w)
			return
		}

		project, err := projectsSvc.Add(r.Context(), (*project.ProjectInput)(&projectInput))
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
