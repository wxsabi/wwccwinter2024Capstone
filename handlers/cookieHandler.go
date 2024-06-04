package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"wwccwinter2024Capstone/models"
)

func CookieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Starting CookieHandler()")

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
		// Query the items for this user
		fmt.Println("	The session token is valid and has not expired")
		rows, err := models.Db.Query("SELECT * FROM items WHERE UserID = ?", userID)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		// Create the response data
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		fmt.Println("Ending CookieHandler()")
		return
	}

	// No cookie found in the request
	w.WriteHeader(http.StatusUnauthorized)
	w.Write([]byte("Unauthorized: Invalid session token"))
	fmt.Println("Ending CookieHandler()")
}

// func CookieHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Starting CookieHandler()")

// 	// Check if the request has a "Cookie" header
// 	if cookie, _ := r.Cookie("session_token"); cookie != nil {
// 		fmt.Println("Cookie", cookie.Value, "Exists")
// 		// Verify the cookie value against a database
// 		var userID int
// 		err := models.Db.QueryRow("SELECT UserID FROM Sessions WHERE SessionID = ? AND ExpiresAt > NOW()", cookie.Value).Scan(&userID)
// 		if err != nil {
// 			if err == sql.ErrNoRows {
// 				// No session found with the given session token, or the session has expired
// 				w.WriteHeader(http.StatusUnauthorized)
// 				w.Write([]byte(""))
// 				return
// 			}
// 			// An error occurred while querying the database
// 			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 			return
// 		}

// 		// The session token is valid and has not expired
// 		// Query the items for this user
// 		fmt.Println("The session token is valid and has not expired")
// 		rows, err := models.Db.Query("SELECT * FROM items WHERE UserID = ?", userID)
// 		if err != nil {
// 			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 			return
// 		}
// 		defer rows.Close()

// 		// Scan the rows into a slice of items
// 		fmt.Println("Scanning the rows into a slice of items")
// 		items := make([]models.Item, 0)
// 		for rows.Next() {
// 			var item models.Item
// 			if err := rows.Scan(&item.ItemID, &item.Name, &item.Description, &item.Price, &item.IsSold, &item.PhotoURL); err != nil {
// 				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 				fmt.Println("Slicing failed")
// 				return
// 			}
// 			items = append(items, item)
// 		}

// 		// Convert the slice to JSON and write it to the response
// 		fmt.Println("This is going to be sent with w using json.NewEncoder: ", items)
// 		// Query for the username
// 		row := models.Db.QueryRow("SELECT Name FROM users WHERE ID = ?", userID)
// 		var username string
// 		err = row.Scan(&username)
// 		if err != nil {
// 			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
// 			return
// 		}

// 		// Create the response data
// 		responseData := models.ResponseData{
// 			Items:    items,
// 			UserName: username,
// 		}

// 		fmt.Println("This will be sent as a response for w: ", responseData)

// 		// Convert the struct to JSON and write it to the response
// 		json.NewEncoder(w).Encode(responseData)
// 		fmt.Println("Ending CookieHandler()")
// 		return
// 	}

// 	// No cookie found in the request
// 	w.WriteHeader(http.StatusUnauthorized)
// 	w.Write([]byte(""))
// 	fmt.Println("Ending CookieHandler()")
// }
