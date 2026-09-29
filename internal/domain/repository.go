package domain

import "context"

type SessionRepository interface {
	ByID(context.Context, SessionID) (*Session, error)
	Save(context.Context, *Session) error
}

type ProjectRepository interface {
	ByID(context.Context, ProjectID) (*Project, error)
	Save(context.Context, *Project) error
}
