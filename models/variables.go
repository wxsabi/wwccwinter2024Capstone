package models

import (
	"sync"
)

var (
	Mu    sync.Mutex
	Items = make([]Item, 0)
)


