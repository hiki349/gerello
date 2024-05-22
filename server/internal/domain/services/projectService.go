package services

import (
	"context"

	"github.com/google/uuid"

	"gerello/internal/domain/models"
)

type ProjectService interface {
	Add(ctx context.Context, input *models.ProjectInput) (*models.Project, error)
	Delete(ctx context.Context, id uuid.UUID) (bool, error)
	Update(ctx context.Context, id uuid.UUID, input *models.ProjectInput) (*models.Project, error)
	FetchAll(ctx context.Context) ([]models.Project, error)
	FetchByID(ctx context.Context, id uuid.UUID) (*models.Project, error)
}
