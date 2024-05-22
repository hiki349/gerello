package api

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrNotFoundID = errors.New("not found id")
	ErrNotValidID = errors.New("not valid id")
)

func getIDFromParams(params map[string]string) (uuid.UUID, error) {
	idStr := params["id"]
	if idStr == "" {
		return uuid.UUID{}, ErrNotFoundID
	}

	id, err := uuid.FromBytes([]byte(idStr))
	if err != nil {
		return uuid.UUID{}, ErrNotValidID
	}

	return id, nil
}
