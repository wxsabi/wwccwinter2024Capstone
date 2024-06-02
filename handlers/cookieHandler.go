package handlers

import (
	"database/sql"
	"net/http"
	"wwccwinter2024Capstone/models"
)

func CookieHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request has a "Cookie" header
	if cookie, _ := r.Cookie("session_token"); cookie != nil {
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
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(""))
		return
	}

	// No cookie found in the request
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte(""))
}
