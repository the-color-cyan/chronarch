package storage

import (
	"chronarch/internal/domain"
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type ProjectRepository struct {
	db *sql.DB
}

func Open(ctx context.Context, path string) (*sql.DB, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", path+"?_foreign_keys=on")
	if err != nil {
		return nil, err
	}
	if err := db.PingContext(ctx); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}

func (p *ProjectRepository) Save(ctx context.Context, project *domain.Project) error {
	query := `
		INSERT INTO projects (id, name, description)
		VALUES (?, ?, ?)
	`

	p.db.ExecContext(ctx, query, project.ID, project.Name, project.Description)

	return nil
}

func (p *ProjectRepository) ByID(ctx context.Context, id domain.ProjectID) (*domain.Project, error) {
	query := `
		SELECT id, name, description
		FROM projects
		WHERE id = ?
	`

	rows := p.db.QueryRowContext(ctx, query, id)

	var name, description string

	err := rows.Scan(id, name, description)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("no entries found for project id: %s", id)
	case err != nil:
		return nil, err
	}

	return domain.NewProjectWithID(id, name, description), nil
}

func InitSchema(ctx context.Context, db *sql.DB) error {
	for _, query := range []string{
		`CREATE TABLE IF NOT EXISTS projects (
           id TEXT PRIMARY KEY,
           name TEXT NOT NULL,
           description TEXT NOT NULL
        )`,
		`CREATE TABLE IF NOT EXISTS sessions (
           id TEXT PRIMARY KEY,
           project_id TEXT NOT NULL REFERENCES projects(id),
           start TEXT NOT NULL,
           end TEXT
        )`,
	} {
		if _, err := db.ExecContext(ctx, query); err != nil {
			return err
		}
	}

	return nil
}
