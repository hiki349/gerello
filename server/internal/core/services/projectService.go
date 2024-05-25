package services

import (
	"context"

	"github.com/google/uuid"

	"gerello/internal/core/domain/project"
)

type ProjectService interface {
	Add(ctx context.Context, input *project.ProjectInput) (*project.Project, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, error)
	Update(ctx context.Context, id uuid.UUID, input *project.ProjectInput) (*project.Project, error)
	FetchAll(ctx context.Context) ([]project.Project, error)
	FetchByID(ctx context.Context, id uuid.UUID) (*project.Project, error)
}
