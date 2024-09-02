package models

import "time"

type Task struct {
	ID       int
	Name     string
	Desc     string
	Comment  string
	Deadline time.Time
	Priority int
}
