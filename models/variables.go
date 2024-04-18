package models

import (
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

// This file is only here to define misc vars that don't
// belong elsewhere for separation of concerns.

var (
	Mu    sync.Mutex
	Items = make([]Item, 0)
	Db    *sql.DB
)
