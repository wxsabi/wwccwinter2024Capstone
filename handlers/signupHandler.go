package handlers

import (
	"fmt"

	"crypto/rand"
	"encoding/hex"
	"io"
	"net/http"
	"os"

	// "os"
	//  "strconv"
	"time"
	"wwccwinter2024Capstone/models"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	models.InitDb() // Initialize database

	fmt.Println(os.Getwd())

	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20) // limit your maxMultipartMemory
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the form fields
	name := r.FormValue("name")
	lastName := r.FormValue("lastname")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Create a new User instance
	user := models.User{
		Name:     name,
		LastName: lastName,
		Email:    email,
		// Password will be filled in after hashing
		// Photo will be filled in after saving the file
		CreatedAt: time.Now(),
		IsAdmin:   false,
	}

	// Get the uploaded file
	file, header, err := r.FormFile("photo")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// Generate a unique random name for the file
	randBytes := make([]byte, 16)
	_, err = rand.Read(randBytes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fileName := hex.EncodeToString(randBytes) + "_" + header.Filename

	// TODO: Save the file to disk or a database here

	// Check if email already exists
	var existingUser models.User
	err = models.Db.QueryRow("SELECT * FROM Users WHERE Email = ?", email).Scan(&existingUser)
	if err == nil {
		http.Error(w, "Email already exists", http.StatusBadRequest)
		return
	}

	// Create the file on disk
	filePath := "/Users/jaycm/OneDrive/Documents/wwccwinter2024Capstone/images/user_photos/" + fileName
	newFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	// Write the uploaded file's content to the new file
	_, err = io.Copy(newFile, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError) //THIS
		return
	}
	user.Photo = filePath

	// Close the file
	err = newFile.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Hash the user's password using bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Initializing -- Set CreatedAt and LastLogin to the current time
	// createdAt := time.Now()

	// Insert the user into the database
	res, err := models.Db.Exec(`
INSERT INTO Users (Name, LastName, Email, Password, Photo, CreatedAt, IsAdmin)
VALUES (?, ?, ?, ?, ?, ?, ?)
`, user.Name, user.LastName, user.Email, user.Password, user.Photo, user.CreatedAt, user.IsAdmin)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the ID field of the user
	id, err := res.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.ID = int(id)

	// Update the name of the image file to prepend the user's ID
	newFilePath := fmt.Sprintf("/Users/jaycm/OneDrive/Documents/wwccwinter2024Capstone/images/user_photos/%d_%s", user.ID, fileName)
	err = os.Rename(user.Photo, newFilePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user.Photo = newFilePath

	// Update the photo path in the database
	_, err = models.Db.Exec("UPDATE Users SET Photo = ? WHERE ID = ?", user.Photo, user.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Print("user ", user.Name, " created sucessfully!")

	// Redirect to the login page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
