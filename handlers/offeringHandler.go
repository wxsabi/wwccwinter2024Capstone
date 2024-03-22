package handlers

import(
	"net/http"
	"github.com/wxsabi/wwccwinter2024Capstone/offering"
	"encoding/json"
)

func offeringHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(map[string][]offering.Offering{"offerings": offerings})
	case http.MethodPost:
		var newItem offering.Offering
		json.NewDecoder(r.Body).Decode(&newItem)
		mu.Lock() // I think this is the best place to lock it...
		newItem.ID = len(offerings) + 1
		offerings = append(offerings, newItem)
		mu.Unlock()
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newItem)
	default:
		http.Error(w, "Jeromy Doesn't allow this method", http.StatusMethodNotAllowed)
	}
}
