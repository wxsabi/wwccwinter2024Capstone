package models

import "time"

type User struct {
	ID        int
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
}

func (u User) DisplayUsername() string {
	return u.Username
}
