package models

import "time"

type Task struct {
	ID         uint64
	Name       string
	Desc       string
	Deadline   time.Time
	Priority   int
	End_status bool
}
