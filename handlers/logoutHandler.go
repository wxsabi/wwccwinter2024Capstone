package handlers

import (
	"fmt"
	"net/http"
	"time"

	"wwccwinter2024Capstone/models"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Starting LogoutHandler.go:")

	fmt.Println("	Loggin out user...")
	// Retrieve the session token from the user's cookies
	cookie, err := r.Cookie("session_token")
	if err != nil {
		// Handle error (e.g., cookie not found)
		fmt.Println("	cookie not found, redirecting to homepage...")
		http.Redirect(w, r, "/index", http.StatusSeeOther)
		return
	}

	// Invalidate the session token (remove it from the map)
	fmt.Println("	Deleting sesion token...")
	delete(models.Sessions, cookie.Value)

	// Delete the session record from the database
	fmt.Println("	deleting from database...")
	_, err = models.Db.Exec("DELETE FROM Sessions WHERE SessionID = ?", cookie.Value)
	if err != nil {
		// Handle error (e.g., log it)
		fmt.Println("	Unable to delete from database")
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Expire the cookie (set expiration time to a past date)
	fmt.Println("	Expiring cookie...")
	expiration := time.Now().Add(-1 * time.Hour)
	deletedCookie := http.Cookie{
		Name:    "session_token",
		Value:   "",
		Path:    "/",
		Expires: expiration,
	}

	http.SetCookie(w, &deletedCookie)

	// Redirect to the login page
	fmt.Println("	redirecting to login...")
	http.Redirect(w, r, "/index?message=User+Logged+Out", http.StatusSeeOther)
	fmt.Println("Ending LogoutHandler.go:")
}
