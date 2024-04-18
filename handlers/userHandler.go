package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"wwccwinter2024Capstone/models"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	var err error
	models.Db, err = sql.Open("mysql", "root:Capstone@/")
	if err != nil {
		log.Fatal(err)
	}

	_, err = models.Db.Exec(`CREATE DATABASE IF NOT EXISTS capDB`)
	if err != nil {
		log.Fatal(err)
	}

	models.Db.Close()

	models.Db, err = sql.Open("mysql", "root:Capstone@/capDB")
	if err != nil {
		log.Fatal(err)
	}

	_, err = models.Db.Exec(`
CREATE TABLE IF NOT EXISTS Users(
	ID INT AUTO_INCREMENT PRIMARY KEY,
	Name VARCHAR(255),
	LastName VARCHAR(255),
	Email VARCHAR(255) UNIQUE,
	Password VARCHAR(255),
	Photo VARCHAR(255),
	CreatedAt TIMESTAMP
)`)

	if err != nil {
		log.Fatal(err)
	}
}

func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rows, err := models.Db.Query("SELECT * FROM Users")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var users []models.User
		for rows.Next() {
			var user models.User
			var createdAtStr string
			err := rows.Scan(&user.ID, &user.Name, &user.LastName, &user.Email, &user.Password, &user.Photo, &createdAtStr)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Parse the CreatedAt string into a time.Time
			layout := "2006-01-02 15:04:05" // adjust this layout to match the format of your timestamp
			createdAtTime, err := time.Parse(layout, createdAtStr)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			user.CreatedAt = createdAtTime // assign the time.Time value directly

			users = append(users, user)
		}

		json.NewEncoder(w).Encode(map[string][]models.User{"users": users})

	case http.MethodPost:
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = models.Db.Exec(`
			INSERT INTO Users (Name, LastName, Email, Password, Photo, CreatedAt)
			VALUES (?, ?, ?, ?, ?, ?)
		`, user.Name, user.LastName, user.Email, user.Password, user.Photo, time.Now())

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)

	case http.MethodDelete:
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = models.Db.Exec(`
			DELETE FROM Users WHERE ID = ?
		`, user.ID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})

	case http.MethodPut:
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = models.Db.Exec(`
			UPDATE Users SET Name = ?, LastName = ?, Email = ?, Password = ?, Photo = ? WHERE ID = ?
		`, user.Name, user.LastName, user.Email, user.Password, user.Photo, user.ID)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(user)

	default:
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
	}
}
