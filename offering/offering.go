package offering

import "time"

type Offering struct {
	ID          int
	Name        string
	Description string
	Price       float64
	ListedAt    time.Time
}

func (i Offering) DisplayName() string {
	return i.Name
}
