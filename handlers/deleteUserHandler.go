package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"wwccwinter2024Capstone/models"
)

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Starting DeleteUserHandler()")

	// Check if the request has a "Cookie" header
	if cookie, _ := r.Cookie("session_token"); cookie != nil {
		fmt.Println("	Cookie", cookie.Value, "Exists")
		// Verify the cookie value against a database
		var userID int
		err := models.Db.QueryRow("SELECT UserID FROM Sessions WHERE SessionID = ? AND ExpiresAt > NOW()", cookie.Value).Scan(&userID)
		if err != nil {
			if err == sql.ErrNoRows {
				// No session found with the given session token, or the session has expired
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(""))
				return
			}
			// An error occurred while querying the database
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		// The session token is valid and has not expired
		// Delete the items, item photos, sessions, and user for this user
		fmt.Println("	The session token is valid and has not expired")

		_, err = models.Db.Exec("DELETE FROM ItemPhotos WHERE ItemID IN (SELECT ItemID FROM Items WHERE UserID = ?)", userID)
		if err != nil {
			log.Fatal(err)
			return
		}

		_, err = models.Db.Exec("DELETE FROM Items WHERE UserID = ?", userID)
		if err != nil {
			log.Fatal(err)
			return
		}

		_, err = models.Db.Exec("DELETE FROM Sessions WHERE UserID = ?", userID)
		if err != nil {
			log.Fatal(err)
			return
		}

		_, err = models.Db.Exec("DELETE FROM Users WHERE ID = ?", userID)
		if err != nil {
			log.Fatal(err)
			return
		}

		// Redirect to the home page
		http.Redirect(w, r, "/", http.StatusSeeOther)
		fmt.Println("Ending DeleteUserHandler()")
		return
	}

	// No cookie found in the request
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Unauthorized: Invalid session token"))
	fmt.Println("Ending DeleteUserHandler()")
}
