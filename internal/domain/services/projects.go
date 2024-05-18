package services

import "gerello/internal/repositories"

// TODO: Implement methods from ProjectService interface

type ProjectServiceImpl struct {
	repo repositories.ProjectRepository
}

func New(repo repositories.ProjectRepository) *ProjectServiceImpl {
	return &ProjectServiceImpl{
		repo: repo,
	}
}
