package models

import (
	"database/sql"
	"log"
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

func InitDb() {
	var err error
	Db, err = sql.Open("mysql", "capstone_user:capstone@/")
	if err != nil {
		log.Fatal(err)
	}

	_, err = Db.Exec(`CREATE DATABASE IF NOT EXISTS capDB`)
	if err != nil {
		log.Fatal(err)
	}

	Db.Close()

	Db, err = sql.Open("mysql", "capstone_user:capstone@/capDB")
	if err != nil {
		log.Fatal(err)
	}

	_, err = Db.Exec(`
	CREATE TABLE IF NOT EXISTS Users (
		ID INT AUTO_INCREMENT PRIMARY KEY,
		Name VARCHAR(255),
		LastName VARCHAR(255),
		Email VARCHAR(255) UNIQUE,
		Password VARCHAR(255),
		Photo VARCHAR(255),
		CreatedAt TIMESTAMP,
		IsAdmin BOOLEAN DEFAULT FALSE
	);
	`)

	if err != nil {
		log.Fatal(err)
	}

	_, err = Db.Exec(`
	CREATE TABLE IF NOT EXISTS Items (
		ID INT PRIMARY KEY,
		Name VARCHAR(255),
		Description VARCHAR(255),
		Price DECIMAL(10, 2),
		ListedAt DATETIME
	);
	`)

	if err != nil {
		log.Fatal(err)
	}
}
