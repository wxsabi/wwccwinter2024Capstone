package handlers

import (
	"net/http"
	"time"

	"wwccwinter2024Capstone/models"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the session token from the user's cookies
	cookie, err := r.Cookie("session_token")
	if err != nil {
		// Handle error (e.g., cookie not found)
		http.Redirect(w, r, "/index", http.StatusSeeOther)
		return
	}

	// Invalidate the session token (remove it from the map)
	delete(models.Sessions, cookie.Value)

	// Delete the session record from the database
	_, err = models.Db.Exec("DELETE FROM Sessions WHERE SessionID = ?", cookie.Value)
	if err != nil {
		// Handle error (e.g., log it)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Expire the cookie (set expiration time to a past date)
	expiration := time.Now().Add(-1 * time.Hour)
	deletedCookie := http.Cookie{
		Name:    "session_token",
		Value:   "",
		Path:    "/",
		Expires: expiration,
	}

	http.SetCookie(w, &deletedCookie)

	// Redirect to the login page
	http.Redirect(w, r, "/index", http.StatusSeeOther)
}
