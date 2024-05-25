package services

import (
	"context"

	"github.com/google/uuid"

	"gerello/internal/core/domain/project"
	"gerello/internal/ports/db"
)

// TODO: Implement methods from ProjectService interface

type ProjectService struct {
	repo db.ProjectRepository
}

func New(repo db.ProjectRepository) *ProjectService {
	return &ProjectService{
		repo: repo,
	}
}

func (s *ProjectService) Add(ctx context.Context, input *project.ProjectInput) (*project.Project, error) {
	return nil, nil
}

func (s *ProjectService) Delete(ctx context.Context, id uuid.UUID) (bool, error) {
	return false, nil
}

func (s *ProjectService) Update(ctx context.Context, id uuid.UUID, input *project.ProjectInput) (*project.Project, error) {
	return nil, nil
}

func (s *ProjectService) FetchAll(ctx context.Context) ([]project.Project, error) {
	return nil, nil
}

func (s *ProjectService) FetchByID(ctx context.Context, id uuid.UUID) (*project.Project, error) {
	return nil, nil
}
