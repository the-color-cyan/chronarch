package domain

import "time"

type Session struct {
	start   time.Time
	end     time.Time
	project Project
}
