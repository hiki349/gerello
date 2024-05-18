package repositories

import (
	"context"

	"github.com/google/uuid"
)

type ProjectRepository interface {
	Add(ctx context.Context, input *ProjectInput) (uuid.UUID, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, error)
	Update(ctx context.Context, input *ProjectInput) (uuid.UUID, error)
	GetAll(ctx context.Context) ([]Project, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Project, error)
}
