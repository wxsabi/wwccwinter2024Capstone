package models

import (
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// This file is only here to define misc vars that don't
// belong elsewhere for separation of concerns.

var (
	Mu    sync.Mutex
	Items = make([]Item, 0)
	Db    *sql.DB
)

// Map to store session information
var Sessions = map[string]Session{}

// Session struct to hold session data
type Session struct {
	Email  string
	Expiry time.Time
}

type TemplateHandler struct {
	Filename string
}

func InitDb() {
	var err error
	Db, err = sql.Open("mysql", "capstone_user:capstone@/")
	if err != nil {
		log.Fatal(err)
	}

	_, err = Db.Exec(`CREATE DATABASE IF NOT EXISTS capstone_database`)
	if err != nil {
		log.Fatal(err)
	}

	Db.Close()

	Db, err = sql.Open("mysql", "capstone_user:capstone@/capstone_database")
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
	CREATE TABLE IF NOT EXISTS Sessions (
		SessionID VARCHAR(255) PRIMARY KEY,
		UserID INT,
		ExpiresAt TIMESTAMP,
		LastLogin TIMESTAMP,    
		IsLoggedIn BOOLEAN DEFAULT FALSE, 
		RememberToken BOOLEAN DEFAULT FALSE,
		FOREIGN KEY(UserID) REFERENCES Users(ID)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = Db.Exec(`
	CREATE TABLE IF NOT EXISTS Items (
		ItemID INT PRIMARY KEY,
		UserID INT,
		Name VARCHAR(255),
		Description VARCHAR(255),
		Price DECIMAL(10, 2),
		ListedAt DATETIME,
		IsSold BOOLEAN DEFAULT TRUE,
		FOREIGN KEY(UserID) REFERENCES Users(ID)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = Db.Exec(`
	CREATE TABLE IF NOT EXISTS ItemPhotos (
		PhotoID INT PRIMARY KEY AUTO_INCREMENT,
		ItemID INT,
		PhotoURL VARCHAR(255),
		FOREIGN KEY (ItemID) REFERENCES Items(ItemID)
	);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

// generateSessionToken generates a new session token
func GenerateSessionToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

type PageData struct {
	Title string
	// Template string
	// TemplateName string
	// Template     *template.Template
}
