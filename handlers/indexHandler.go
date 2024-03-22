package handlers

import (
	"html/template"
	"net/http"
	"sync"
	"github.com/wxsabi/wwccwinter2024Capstone/offering"
)

var (
	mu    sync.Mutex
	offerings = make([]offering.Offering, 0)
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	mu.Lock()
	tmpl.Execute(w, offerings)
	mu.Unlock()
}
