package models

import "time"

type User struct {
	ID        int
	Name      string
	LastName  string
	Email     string
	Password  string
	Photo     string // This will be a URL to a filesystem
	CreatedAt time.Time
}

// func (u User) DisplayUsername() string {
// 	return u.Username
// }
