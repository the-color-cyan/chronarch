package domain

import (
	"errors"
	"fmt"
	"time"
)

var ErrInvalidTimeRange = errors.New("invalid time range")

type Session struct {
	start   time.Time
	end     time.Time
	project *Project
}

func NewSession(project *Project, at time.Time) *Session {
	var end time.Time

	return &Session{
		start:   at,
		end:     end,
		project: project,
	}
}

func NewRetrospectiveSession(
	project *Project,
	start, end time.Time,
) (*Session, error) {
	if err := validateTimeRange(start, end); err != nil {
		return nil, err
	}

	return &Session{
		start,
		end,
		project,
	}, nil
}

func (s *Session) Stop(at time.Time) error {
	if err := validateTimeRange(s.start, at); err != nil {
		return err
	}

	s.end = at
	return nil
}

func (s *Session) Duration() time.Duration {
	return s.end.Sub(s.start)
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
