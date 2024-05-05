package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"wwccwinter2024Capstone/models"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	models.InitDb() // Initialize database
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash the user's password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	// Insert the user into the database
	_, err = models.Db.Exec(`
        INSERT INTO Users (Name, LastName, Email, Password, Photo, CreatedAt, IsAdmin, SessionID, LastLogin, IsLoggedIn, RememberToken)
        VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `, user.Name, user.LastName, user.Email, hashedPassword, user.Photo, time.Now(), user.IsAdmin, user.SessionID, user.LastLogin, user.IsLoggedIn, user.RememberToken)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
