package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"wwccwinter2024Capstone/models"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:Capstone@/")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE DATABASE IF NOT EXISTS capDB`)
	if err != nil {
		log.Fatal(err)
	}

	db.Close()

	db, err = sql.Open("mysql", "root:Capstone@/capDB")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Items(
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
		rows, err := db.Query("SELECT * FROM Items")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var items []models.Item
		for rows.Next() {
			var item models.Item
			err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.ListedAt)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			items = append(items, item)
		}

		json.NewEncoder(w).Encode(map[string][]models.Item{"items": items})

	case http.MethodPost:
		var newItem models.Item //decode req body into item struct
		json.NewDecoder(r.Body).Decode(&newItem)

		_, err := db.Exec("INSERT INTO Items (ID, Name, Description, Price, ListedAt) VALUES (?, ?, ?, ?, ?)",
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

		_, err := db.Exec("DELETE FROM Items WHERE ID = ?", idToDelete)
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

		_, err = db.Exec("UPDATE Items SET Name = ?, Description = ?, Price = ? WHERE ID = ?",
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

/* Old  code
import(
	"net/http"
	"wwccwinter2024Capstone/models"
	"encoding/json"
)



// ItemHandler is a function that handles HTTP requests related to items.
func ItemHandler(w http.ResponseWriter, r *http.Request) {
	// The behavior of the handler depends on the HTTP method of the request.
	switch r.Method {
		// For GET requests, we return a JSON representation of all items.
		case http.MethodGet:
			json.NewEncoder(w).Encode(map[string][]models.Item{"items": models.Items})
			// For POST requests, we add a new item to our list.
		case http.MethodPost:
			// We decode the request body into an Item struct.
			var newItem models.Item
			json.NewDecoder(r.Body).Decode(&newItem)

			// We lock the mutex to prevent concurrent access to the items slice.
			models.Mu.Lock()

			// We assign an ID to the new item and append it to the items slice.
			newItem.ID = len(models.Items) + 1
			models.Items = append(models.Items, newItem)

			// We unlock the mutex.
			models.Mu.Unlock()

			// We respond with a status code indicating that the item was created.
			w.WriteHeader(http.StatusCreated)

			// We return a JSON representation of the new item.
			json.NewEncoder(w).Encode(newItem)
			// For all other HTTP methods, we respond with a 'method not allowed' error.
		default:
			http.Error(w, "Invalid Method", http.StatusMethodNotAllowed)
	}
}
*/
