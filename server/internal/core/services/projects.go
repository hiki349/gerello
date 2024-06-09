package services

import (
	"context"

	"github.com/google/uuid"

	"gerello/internal/core/domain/project"
	"gerello/internal/ports/db"
)

type ProjectService struct {
	repo db.ProjectRepository
}

func New(repo db.ProjectRepository) *ProjectService {
	return &ProjectService{
		repo: repo,
	}
}

func (s *ProjectService) Add(ctx context.Context, input *project.ProjectInput) (*project.Project, error) {
	id, err := s.repo.Add(ctx, input)
	if err != nil {
		return nil, err
	}

	project, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (s *ProjectService) Delete(ctx context.Context, id uuid.UUID) (bool, error) {
	ok, err := s.repo.Delete(ctx, id)
	if err != nil {
		return false, err
	}

	return ok, nil
}

func (s *ProjectService) Update(ctx context.Context, id uuid.UUID, input *project.ProjectInput) (*project.Project, error) {
	id, err := s.repo.Update(ctx, input, id)
	if err != nil {
		return nil, err
	}

	project, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return project, nil
}

func (s *ProjectService) FetchAll(ctx context.Context) ([]project.Project, error) {
	projects, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (s *ProjectService) FetchByID(ctx context.Context, id uuid.UUID) (*project.Project, error) {
	project, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return project, nil
}
