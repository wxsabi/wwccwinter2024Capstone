package handlers

import (
	// "database/sql"
	"encoding/json"
	// "log"
	"net/http"
	"time"
	"wwccwinter2024Capstone/models"

	_ "github.com/go-sql-driver/mysql"
)

func ItemHandler(w http.ResponseWriter, r *http.Request) {
	models.InitDb() // Initialize database
	switch r.Method {
	case http.MethodGet:
		// This implicitly sets the value of rows to be the
		// rows from the items table of the database
		// if there are none, an error will be thrown.
		rows, err := models.Db.Query("SELECT * FROM Items")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close() // This closes the rows once the function is done.

		var items []models.Item
		for rows.Next() {
			var item models.Item
			var listedAtStr string
			err := rows.Scan(&item.ItemID, &item.UserID, &item.Name, &item.Description, &item.Price, &listedAtStr)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Parse the ListedAt string into a time.Time
			layout := "2006-01-02 15:04:05" // adjust this layout to match the format of timestamp
			listedAtTime, err := time.Parse(layout, listedAtStr)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			item.ListedAt = listedAtTime // assign the time.Time value directly

			items = append(items, item)
		}

		json.NewEncoder(w).Encode(map[string][]models.Item{"items": items})

	case http.MethodPost:
		var newItem models.Item //decode req body into item struct
		json.NewDecoder(r.Body).Decode(&newItem)

		_, err := models.Db.Exec("INSERT INTO Items (ItemID, UserID, Name, Description, Price, ListedAt) VALUES (?, ?, ?, ?, ?, ?)",
			newItem.ItemID, newItem.UserID, newItem.Name, newItem.Description, newItem.Price, time.Now())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated) // return JSON of new line
		json.NewEncoder(w).Encode(newItem)

	case http.MethodDelete:
		var idToDelete int
		json.NewDecoder(r.Body).Decode(&idToDelete)

		_, err := models.Db.Exec("DELETE FROM Items WHERE ID = ?", idToDelete)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

	case http.MethodPut:
		var itemToUpdate models.Item
		err := json.NewDecoder(r.Body).Decode(&itemToUpdate)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = models.Db.Exec("UPDATE Items SET Name = ?, Description = ?, Price = ? WHERE ID = ?",
			itemToUpdate.Name, itemToUpdate.Description, itemToUpdate.Price, itemToUpdate.ItemID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed) // for all other http methods
	}
}
