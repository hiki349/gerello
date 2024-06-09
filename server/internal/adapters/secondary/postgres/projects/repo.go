package projects

import (
	"context"
	"time"

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
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*500)
	defer cancel()

	var projectID uuid.UUID
	sql := `INSERT INTO projects (title, description) VALUES ($1, $2) RETURNING id;`

	err := r.db.QueryRow(ctx, sql, input.Title, input.Description).Scan(&projectID)
	if err != nil {
		return uuid.UUID{}, err
	}

	return projectID, nil
}

func (r *ProjectRepositoryInPostgres) Delete(ctx context.Context, id uuid.UUID) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*500)
	defer cancel()

	sql := `DELETE FROM projects WHERE id = $1;`
	_, err := r.db.Exec(ctx, sql, id)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *ProjectRepositoryInPostgres) Update(ctx context.Context, input *project.ProjectInput, id uuid.UUID) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*500)
	defer cancel()

	var projectID uuid.UUID
	sql := `UPDATE projects SET title = $1, description = $2 WHERE id = $3 RETURNING id;`

	err := r.db.QueryRow(ctx, sql, input.Title, input.Description, id).Scan(&projectID)
	if err != nil {
		return uuid.UUID{}, err
	}

	return projectID, nil
}

func (r *ProjectRepositoryInPostgres) GetAll(ctx context.Context) ([]project.Project, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*500)
	defer cancel()

	var projects []project.Project
	sql := `SELECT id, title, description, updated_at, created_at FROM projects;`

	rows, err := r.db.Query(ctx, sql)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var project project.Project

		err := rows.Scan(
			&project.ID,
			&project.Title,
			&project.Description,
			&project.UpdateAt,
			&project.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}

	rows.Close()
	if rows.Err() != nil {
		return nil, err
	}

	return projects, nil
}

func (r *ProjectRepositoryInPostgres) GetByID(ctx context.Context, id uuid.UUID) (*project.Project, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*500)
	defer cancel()

	var project project.Project
	sql := `SELECT id, title, description, updated_at, created_at FROM projects WHERE id = $1;`

	err := r.db.QueryRow(ctx, sql, id).Scan(
		&project.ID,
		&project.Title,
		&project.Description,
		&project.UpdateAt,
		&project.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &project, nil
}
