package handlers

import (
	"html/template"
	"net/http"
	"wwccwinter2024Capstone/models"
)

// IndexHandler is a function that handles requests to the index page.
func PostHtmlHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the layout.html and index.html files. If there's an error, it will panic.
	tmpl, err := template.ParseFiles("html/layout.html", "html/post.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data := models.PageData{
		Title: "Britl!",
	}

	// Lock the mutex to prevent other goroutines from accessing the data at the same time.
	models.Mu.Lock()

	// Execute the template, writing the generated HTML to the http.ResponseWriter.
	// The second parameter is the data we want to pass into the template.
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Unlock the mutex to allow other goroutines to access the data.
	models.Mu.Unlock()
}
