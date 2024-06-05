package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"wwccwinter2024Capstone/models"
)

func GetUserItems(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Starting GetUserItems()")
	// Get the user ID from the session
	//fmt.Println("Response recieved by /userItems:", r)

	cookie, err := r.Cookie("session_token")
	if err != nil {
		http.Error(w, "No session token in cookie", http.StatusBadRequest)
		return
	}
	sessionToken := cookie.Value

	// Query the sessions table for the user ID
	var userID string
	err = models.Db.QueryRow("SELECT UserID FROM Sessions WHERE SessionID = ?", sessionToken).Scan(&userID)
	if err != nil {
		http.Error(w, "Failed to get user ID from session token", http.StatusInternalServerError)
		return
	}

	rows, err := models.Db.Query(`
        SELECT Items.ItemID, Items.UserID, Items.Name, Items.Description, Items.Price, Items.ListedAt, Items.IsSold, ItemPhotos.PhotoURL 
        FROM Items 
        LEFT JOIN ItemPhotos ON Items.ItemID = ItemPhotos.ItemID 
        WHERE Items.UserID = ?
        ORDER BY Items.ListedAt DESC 
        LIMIT 5 
    `, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("error with models.Db.Query")
		return
	}
	var items []models.Item
	fmt.Println("created var items")
	for rows.Next() {
		var item models.Item
		var listedAt []byte // Temporary variable to hold the ListedAt value

		err = rows.Scan(&item.ItemID, &item.UserID, &item.Name, &item.Description, &item.Price, &listedAt, &item.IsSold, &item.PhotoURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("error with rows.Scan:", err)
			return
		}
		// Convert listedAt to a string and parse it into a time.Time value.
		item.ListedAt, err = time.Parse("2006-01-02 15:04:05", string(listedAt))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("error with time.Parse:", err)
			return
		}
		items = append(items, item)
	}
	//fmt.Println("items = append(items, item)", items)
	json.NewEncoder(w).Encode(items)
	defer rows.Close()

	// jsonData, err := json.Marshal(items)
	// if err != nil {
	// 	// handle error
	// 	fmt.Println("error with json.Marshal:", err)
	// 	return
	// }
	//fmt.Println("JSON data to send: ", string(jsonData))

	fmt.Println("Ending GetUserItems()")
}
