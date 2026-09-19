package domain

import "context"

type SessionRepository interface {
	ByID(context.Context, *Session)
	Save(context.Context, *Session) error
}

type ProjectRepository interface {
	ByID(context.Context, *Project)
	Save(context.Context, *Project) error
}
