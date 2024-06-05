package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"wwccwinter2024Capstone/models"
)

func SeeItemHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Starting SeeItemHandler()")
	fmt.Println("	received: ", r)

	// Parse the URL query parameters
	params := r.URL.Query()

	// Get the itemId parameter
	itemId := params.Get("itemId")
	fmt.Println("	", itemId)

	// Query the database for the item with the specific itemId and UserID = ID from users table
	row := models.Db.QueryRow("SELECT * FROM Items WHERE Items.ItemID = ?", itemId)

	var item models.Item
	var listedAt string
	err := row.Scan(&item.ItemID, &item.UserID, &item.Name, &item.Description, &item.Price, &listedAt, &item.IsSold)
	if err != nil {
		if err == sql.ErrNoRows {
			// No rows were returned
			http.Error(w, "No item found with the provided ID", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	row = models.Db.QueryRow("SELECT PhotoURL FROM ItemPhotos WHERE ItemID = ?", itemId)
	err = row.Scan(&item.PhotoURL)
	if err != nil {
		if err == sql.ErrNoRows {
			// No rows were returned
			http.Error(w, "No item found with the provided ID", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
	// Convert listedAt to a string and parse it into a time.Time value.
	item.ListedAt, err = time.Parse("2006-01-02 15:04:05", string(listedAt))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("	error with time.Parse:", err)
		return
	}

	// Convert the item to JSON and write it to the response
	json.NewEncoder(w).Encode(item)

	fmt.Println("Ending SeeItemHandler()")
}
