package project

import (
	"time"

	"github.com/google/uuid"
)

type Project struct {
	ID          uuid.UUID
	Title       string
	Description string
	CreatedAt   time.Time
	UpdateAt    time.Time
}

type ProjectInput struct {
	Title       string
	Description string
}
