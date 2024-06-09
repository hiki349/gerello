package db

import (
	"context"

	"github.com/google/uuid"

	"gerello/internal/core/domain/project"
)

type ProjectRepository interface {
	Add(ctx context.Context, input *project.ProjectInput) (uuid.UUID, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, error)
	Update(ctx context.Context, input *project.ProjectInput, id uuid.UUID) (uuid.UUID, error)
	GetAll(ctx context.Context) ([]project.Project, error)
	GetByID(ctx context.Context, id uuid.UUID) (*project.Project, error)
}
