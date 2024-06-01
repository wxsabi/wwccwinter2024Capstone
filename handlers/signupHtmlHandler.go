package handlers

import (
	"html/template"
	"net/http"
	"wwccwinter2024Capstone/models"
)

func SignupHtmlHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, htmlerr := template.ParseFiles("html/layout.html", "signup.html")
	if htmlerr != nil {
		http.Error(w, htmlerr.Error(), http.StatusInternalServerError)
	}
	data := models.PageData{
		Title: "Create User - Britl!",
	}
	models.Mu.Lock()
	htmlerr = tmpl.Execute(w, data)
	if htmlerr != nil {
		http.Error(w, htmlerr.Error(), http.StatusInternalServerError)
	}
	models.Mu.Unlock()

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
}
