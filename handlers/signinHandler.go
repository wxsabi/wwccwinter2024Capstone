package handlers

import (
	//"database/sql"
	"log"
	"wwccwinter2024Capstone/models"

	"golang.org/x/crypto/bcrypt"
	// Other necessary imports
)

func AuthenticateUser(email, inputPassword string) (bool, error) {
	// Retrieve user data from the database based on the provided email
	// (Assuming you have a 'users' table with columns: 'email' and 'password')
	query := "SELECT password FROM users WHERE email = ?"
	var storedHashedPassword string
	err := models.Db.QueryRow(query, email).Scan(&storedHashedPassword)
	if err != nil {
		log.Println("Error retrieving user data:", err)
		return false, err
	}

	// Compare the stored hashed password with the input password
	err = bcrypt.CompareHashAndPassword([]byte(storedHashedPassword), []byte(inputPassword))
	if err != nil {
		log.Println("Invalid credentials:", err)
		return false, nil
	}

	// Authentication successful
	return true, nil
}
