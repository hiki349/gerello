package projects

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"

	"gerello/internal/core/domain/project"
)

type ProjectRepositoryInPostgres struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *ProjectRepositoryInPostgres {
	return &ProjectRepositoryInPostgres{
		db: db,
	}
}

func (r *ProjectRepositoryInPostgres) Add(ctx context.Context, input *project.ProjectInput) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}

func (r *ProjectRepositoryInPostgres) Delete(ctx context.Context, id uuid.UUID) (bool, error) {
	return false, nil
}

func (r *ProjectRepositoryInPostgres) Update(ctx context.Context, input *project.ProjectInput) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}

func (r *ProjectRepositoryInPostgres) GetAll(ctx context.Context) ([]project.Project, error) {
	return nil, nil
}

func (r *ProjectRepositoryInPostgres) GetByID(ctx context.Context, id uuid.UUID) (*project.Project, error) {
	return nil, nil
}
