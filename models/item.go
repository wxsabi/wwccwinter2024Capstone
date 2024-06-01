package models

import "time"

type Item struct {
	ItemID      int
	UserID      int
	Name        string
	Description string
	Price       float64
	ListedAt    time.Time
	IsSold      bool
}
