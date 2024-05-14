package handlers

import (
	"log"
	"net/http"
	"time"

	"wwccwinter2024Capstone/models"

	"golang.org/x/crypto/bcrypt"
)

func SigninHandler(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	inputPassword := r.FormValue("password")

	// Retrieve user data from the database based on the provided email
	query := "SELECT password FROM users WHERE email = ?"
	var storedHashedPassword string
	err := models.Db.QueryRow(query, email).Scan(&storedHashedPassword)
	if err != nil {
		log.Printf("Error retrieving user data for email %s: %v", email, err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Compare the stored hashed password with the input password
	err = bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(inputPassword))
	if err != nil {
		log.Printf("Invalid credentials for email %s: %v", email, err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Authentication successful
	log.Printf("User %s logged in successfully", email)

	// Generate a new session token (use a UUID or other secure method)
	sessionToken, err := models.GenerateSessionToken()
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Set the expiry time for the session to 30 days
	expiresAt := time.Now().Add(30 * 24 * time.Hour)

	// Store the session token
	models.Sessions[sessionToken] = models.Session{
		Email:  email, // Use the email as the username
		Expiry: expiresAt,
	}

	// Create a new cookie with the session token
	cookie := http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Path:     "/", // Set the path where the cookie is valid (root path in this case)
		Expires:  expiresAt,
		HttpOnly: true, // Prevent JavaScript access to the cookie
	}

	http.SetCookie(w, &cookie)

	http.Redirect(w, r, "/welcome", http.StatusSeeOther)
}
