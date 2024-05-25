package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"wwccwinter2024Capstone/models"

	_ "github.com/go-sql-driver/mysql"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			// Log the panic
			log.Printf("Panic occurred: %v", r)
		}
	}()
	models.InitDb() // Initialize database
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
        INSERT INTO Users (Name, LastName, Email, Password, Photo, CreatedAt, IsAdmin)
        VALUES (?, ?, ?, ?, ?, ?, ?)
    `, user.Name, user.LastName, user.Email, user.Password, user.Photo, time.Now(), user.IsAdmin)

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
			DELETE FROM ItemPhotos WHERE ItemID IN (
				SELECT ItemID FROM Items WHERE UserID = ?
			)
		`, user.ID)
		if err != nil {
			log.Fatal(err)
			return
		}

		_, err = models.Db.Exec(`
			DELETE FROM Items WHERE UserID = ?
		`, user.ID)
		if err != nil {
			log.Fatal(err)
			return
		}

		_, err = models.Db.Exec(`
			DELETE FROM Sessions WHERE UserID = ?
		`, user.ID)
		if err != nil {
			log.Fatal(err)
			return
		}

		_, err = models.Db.Exec(`
			DELETE FROM Users WHERE ID = ?
		`, user.ID)
		if err != nil {
			log.Fatal(err)
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
		UPDATE Users
		SET Name=?, LastName=?, Email=?, Photo=?
		WHERE ID=?
	`, user.Name, user.LastName, user.Email, user.Photo, user.ID)

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
