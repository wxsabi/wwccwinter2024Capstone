package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"wwccwinter2024Capstone/models"
)

// SeeItemHtmlHandler is a function that handles requests to the see item page.
func SeeItemHtmlHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Starting SeeItemHtmlHandler()")
	// Parse the layout.html and item.html files. If there's an error, it will panic.
	tmpl, err := template.ParseFiles("html/layout.html", "html/item.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the itemId from the query parameters
	itemId := r.URL.Query().Get("itemId")
	if itemId == "" {
		http.Error(w, "No itemId provided", http.StatusBadRequest)
		return
	}

	data := models.PageData{
		Title: "Britl!",
		// Add the itemId to the data
		ItemId: itemId,
	}

	// Lock the mutex to prevent other goroutines from accessing the data at the same time.
	models.Mu.Lock()

	// Execute the template, writing the generated HTML to the http.ResponseWriter.
	// The second parameter is the data we want to pass into the template.
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Unlock the mutex to allow other goroutines to access the data.
	models.Mu.Unlock()
	fmt.Println("Ending SeeItemHtmlHandler()")
}
