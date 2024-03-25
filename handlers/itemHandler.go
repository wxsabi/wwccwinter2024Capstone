package handlers

import(
	"net/http"
	"github.com/wxsabi/wwccwinter2024Capstone/models"
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
