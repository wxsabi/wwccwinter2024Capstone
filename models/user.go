package models

import "time"

type User struct {
	ID            int
	Name          string
	LastName      string
	Email         string
	Password      string
	Photo         string
	CreatedAt     time.Time
	IsAdmin       bool
	SessionID     string
	LastLogin     time.Time
	IsLoggedIn    bool
	RememberToken bool
}
