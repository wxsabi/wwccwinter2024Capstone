package handlers

import (
	"net/http"
)

func CookieHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request has a "Cookie" header
	if cookie, ok := r.Cookie("session_token"); ok {
		// Verify the cookie value (e.g., check against a database)
		// For this example, assume the cookie is valid
		http.WriteResponse(w, http.StatusOK, "")
		return
	}
	http.WriteResponse(w, http.StatusUnauthorized, "")
}
