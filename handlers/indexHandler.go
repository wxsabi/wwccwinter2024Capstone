package handlers

import (
	"html/template"
	"net/http"
	"github.com/wxsabi/wwccwinter2024Capstone/models"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	models.Mu.Lock()
	tmpl.Execute(w, models.Items)
	models.Mu.Unlock()
}
