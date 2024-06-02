package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"wwccwinter2024Capstone/models"
)

func GetNewestItems(w http.ResponseWriter, r *http.Request) {
	models.InitDb() // Initialize database
	fmt.Println("database initialized")
	rows, err := models.Db.Query(`
		SELECT Items.ItemID, Items.UserID, Items.Name, Items.Description, Items.Price, Items.ListedAt, Items.IsSold, ItemPhotos.PhotoURL 
		FROM Items 
		LEFT JOIN ItemPhotos ON Items.ItemID = ItemPhotos.ItemID 
		ORDER BY Items.ListedAt DESC 
		LIMIT 5 
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("error with models.Db.Query")
		return
	}
	var items []models.Item
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
		fmt.Println("error items = append(items, item)")
	}

	json.NewEncoder(w).Encode(items)
	fmt.Println("json.NewEncoder(w).Encode(items)")
	defer rows.Close()
	fmt.Println("rows closed")
}
