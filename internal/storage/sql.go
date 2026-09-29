package storage

import (
	"chronarch/internal/domain"
	"context"
	"database/sql"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

type ProjectRepository struct{}

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
	panic("not yet implemented")
}

func (p *ProjectRepository) ByID(ctx context.Context, id domain.ProjectID) (*domain.Project, error) {
	panic("not yet implemented")
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
