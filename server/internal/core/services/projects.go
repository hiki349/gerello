package services

import "gerello/internal/ports/db"

// TODO: Implement methods from ProjectService interface

type ProjectServiceImpl struct {
	repo db.ProjectRepository
}

func New(repo db.ProjectRepository) *ProjectServiceImpl {
	return &ProjectServiceImpl{
		repo: repo,
	}
}
