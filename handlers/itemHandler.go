package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
	"wwccwinter2024Capstone/models"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	var err error
	models.Db, err = sql.Open("mysql", "capstone_user:capstone@/")
	if err != nil {
		log.Fatal(err)
	}

	_, err = models.Db.Exec(`CREATE DATABASE IF NOT EXISTS capDB`)
	if err != nil {
		log.Fatal(err)
	}

	models.Db.Close()

	models.Db, err = sql.Open("mysql", "capstone_user:capstone@/ capDB")
	if err != nil {
		log.Fatal(err)
	}

	_, err = models.Db.Exec(`CREATE TABLE IF NOT EXISTS Items(
		ID INT PRIMARY KEY,
		Name VARCHAR(255),
		Description VARCHAR(255),
		Price DECIMAL(10, 2),
		ListedAt DATETIME
	)`)
	if err != nil {
		log.Fatal(err)
	}
}

func ItemHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rows, err := models.Db.Query("SELECT * FROM Items")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var items []models.Item
		for rows.Next() {
			var item models.Item
			var listedAtStr string
			err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &listedAtStr)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			// Parse the ListedAt string into a time.Time
			layout := "2006-01-02 15:04:05" // adjust this layout to match the format of your timestamp
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

		_, err := models.Db.Exec("INSERT INTO Items (ID, Name, Description, Price, ListedAt) VALUES (?, ?, ?, ?, ?)",
			newItem.ID, newItem.Name, newItem.Description, newItem.Price, newItem.ListedAt)
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
			itemToUpdate.Name, itemToUpdate.Description, itemToUpdate.Price, itemToUpdate.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

	default:
		http.Error(w, "Invalid Method", http.StatusMethodNotAllowed) // for all other http methods
	}
}
