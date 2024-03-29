package models

import "time"

type Item struct {
	ID          int
	Name        string
	Description string
	Price       float64
	ListedAt    time.Time
}

func (i Item) DisplayName() string {
	return i.Name
}
