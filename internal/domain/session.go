package domain

import (
	"errors"
	"fmt"
	"time"
)

var ErrInvalidTimeRange = errors.New("invalid time range")

type SessionID string

type Session struct {
	start   time.Time
	end     time.Time
	project *Project
	id      SessionID
}

func NewSession(project *Project, at time.Time) (*Session, error) {
	var end time.Time

	id, err := newSessionID()
	if err != nil {
		return nil, err
	}

	return &Session{
		start:   at,
		end:     end,
		project: project,
		id:      id,
	}, nil
}

func NewRetrospectiveSession(
	project *Project,
	start, end time.Time,
) (*Session, error) {
	if err := validateTimeRange(start, end); err != nil {
		return nil, err
	}

	id, err := newSessionID()
	if err != nil {
		return nil, err
	}

	return &Session{
		start,
		end,
		project,
		id,
	}, nil
}

func (s *Session) StartTime() time.Time { return s.start }

func (s *Session) EndTime() time.Time { return s.end }

func (s *Session) Project() *Project { return s.project }

func (s *Session) ID() SessionID { return s.id }

func (s *Session) Duration() time.Duration {
	return s.end.Sub(s.start)
}

func (s *Session) Stop(at time.Time) error {
	if err := validateTimeRange(s.start, at); err != nil {
		return err
	}

	s.end = at
	return nil
}

func newSessionID() (SessionID, error) {
	id, err := newID()
	if err != nil {
		return "", err
	}

	return SessionID(id), nil
}

func validateTimeRange(start, end time.Time) error {
	if end.Before(start) {
		return fmt.Errorf(
			"%w: start time %s is after end time %s",
			ErrInvalidTimeRange,
			start,
			end,
		)
	}

	return nil
}
