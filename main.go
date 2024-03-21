// Create a simple web server that listens on port 8888. The server should have
// the following endpoints:
//
// GET /index.html
//
//	serves the index.html template file. The contents should include all
//	items in the grocery list.
//
// GET /item
//
//	returns the contents of the grocery list in JSON format. The JSON
//	response should look like this:
//
//	{"items": [{"id": 1, "name": "milk"}, {"id": 2, "name": "eggs"}]}
//
// POST /item
//
//	adds an item to the grocery list. The item should be added to the list in
//	the order it was received. The request body should be in JSON format. The
//	JSON request should look like this:
//
//	{"name": "milk"}
//
//	The response should be a 201 Created status code. The response body should
//	be in JSON format. The JSON response should look like this:
//
//	{"id": 1, "name": "milk"}
//
//	The ID should be the next available ID in the list. The ID should be unique
//	across all items in the list.
//
// DELETE /item/{id}
//
//	removes an item from the grocery list. The response should be a 204 No
//	Content status code.
//
// PUT /item/{id}
//
//	updates an item in the grocery list. The request body should be in JSON
//	format. The JSON request should look like this:
//
//	{"name": "milk"}
//
//	The response should be a 200 OK status code. The response body should be in
//	JSON format. The JSON response should look like this:
//
//	{"id": 1, "name": "milk"}
//
// In summary:
//  1. Create a web server that listens on port 8888.
//  2. Create a handler function named "indexHandler" that serves the
//     index.html template file.
//  3. Create a handler function named "itemHandler" that handles requests to
//     the "/item" endpoint.
//  4. Create a handler function named "itemIDHandler" that handles requests to
//     the "/item/{id}" endpoint.
//  5. The "/item" endpoint should support the following HTTP methods:
//     GET and POST.
//  6. The "/item/{id}" endpoint should support the following HTTP methods:
//     DELETE and PUT.
//
// Hints:
//  1. Store the grocery list in a mutex-protected variable.
//  2. Use the "html/template" package to serve the index.html template file.
//  3. Use the "encoding/json" package to encode and decode JSON.
//  4. Use the "net/http" package to create the web server.
//
// Bonus:
//  1. Add the ability to add, update and remove items in the grocery list using
//     your webpage.
package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Item struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var (
	mu    sync.Mutex
	items = make([]Item, 0)
)

func main() {
	http.HandleFunc("/index.html", indexHandler)
	http.HandleFunc("/item", itemHandler)
	http.HandleFunc("/item/", itemIDHandler)
	fmt.Println("Web Server Started") // this has to go before ListenAndServe
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	mu.Lock()
	tmpl.Execute(w, items)
	mu.Unlock()
}

func itemHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(map[string][]Item{"items": items})
	case http.MethodPost:
		var newItem Item
		json.NewDecoder(r.Body).Decode(&newItem)
		mu.Lock() // I think this is the best place to lock it...
		newItem.ID = len(items) + 1
		items = append(items, newItem)
		mu.Unlock()
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newItem)
	default:
		http.Error(w, "Jeromy Doesn't allow this method", http.StatusMethodNotAllowed)
	}
}

func itemIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/item/"):]
	id, err := strconv.Atoi(idStr) // convert to int for use later
	if err != nil {
		http.Error(w, "This ID is soooo invalid it made Jeromy cry", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodDelete:
		mu.Lock()
		for i, item := range items {
			if item.ID == id {
				items = append(items[:i], items[i+1:]...)
				break
			}
		}
		mu.Unlock()
		w.WriteHeader(http.StatusNoContent)
	case http.MethodPut:
		var updatedItem Item
		json.NewDecoder(r.Body).Decode(&updatedItem)
		updatedItem.ID = id
		mu.Lock()
		found := false
		for i, item := range items {
			if item.ID == id {
				items[i] = updatedItem
				found = true
				break
			}
		}
		mu.Unlock()
		if !found {
			http.Error(w, "Jeromy DENIETH you this item... because it doesn't exist.", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(updatedItem)
	}
}
