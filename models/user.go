package models

import "time"

type User struct {
	ID        int
	Name      string
	LastName  string
	Email     string
	Password  string // not sure if a string is ok for this
	Photo     string // This will be a URL to a filesystem
	CreatedAt time.Time
}
