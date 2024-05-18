package services

import "gerello/internal/repositories"

type ProjectServiceImpl struct {
	repo repositories.ProjectRepository
}

func New(repo repositories.ProjectRepository) *ProjectServiceImpl {
	return &ProjectServiceImpl{
		repo: repo,
	}
}
