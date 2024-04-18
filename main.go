package main

import (
	"fmt"
	"log"
	"net/http"
	"wwccwinter2024Capstone/handlers"
)

func main() {
	// This will handle the homepage
	http.HandleFunc("/index.html", handlers.IndexHandler)

	// This will handle the postings/items
	http.HandleFunc("/item", handlers.ItemHandler)

	// This will start the web server
	fmt.Println("Server is running on port: 8888") // this has to go before ListenAndServe
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err)
	}
}

/*
Old Code:


import (
	// "encoding/json"
	"fmt"
	// "html/template"
	"log"
	"net/http"

	"wwccwinter2024Capstone/handlers"
	// "strconv"
	// "sync"
	// When ready, add this to mod.go "require golang.org/x/crypto v0.21.0"
	// "golang.org/x/crypto/bcrypt" // crypto/bcrypt package to hash and salt
	// "github.com/wxsabi/wwccwinter2024Capstone/offering"
	// "github.com/wxsabi/wwccwinter2024Capstone/user"

)



func main() {
	// This will handle the homepage
	http.HandleFunc("/index.html", handlers.IndexHandler)

	// This will handle the postings/items
	http.HandleFunc("/item", handlers.ItemHandler)

	// This will handle the search results
	// http.HandleFunc("/search", searchHandler)

	// This will redirect to the login page
	// http.HandleFunc("/login", loginHandler)

	// This will redirect to the register page
	// http.HandleFunc("/register", registerHandler)

	// This will redirect to the report page
	// http.HandleFunc("/report", reportHandler)

	// This will take to the shopping cart
	// http.HandleFunc("/cart", cartHandler)

	// This will take to the checkout
	// http.HandleFunc("/cart/checkout", chekoutHandler)

	// This is only done as an Idea for now but it will handle the DB
	//http.HandleFunc("/offering/", offeringIDHandler)

	// This will start the web server
	fmt.Println("Web Server Started") // this has to go before ListenAndServe
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err)
	}

	// // This is a temporary code snippet that will manage the way user passwords will
	// // be obfuscated. It may be moved to another area of the code later on:
	// password := "userPassword"
 //
	// // Hashing the password with the default cost of 10
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// if err != nil {
	// 	panic(err)
	// }
 //
	// fmt.Println(string(hashedPassword))
}




// func itemIDHandler(w http.ResponseWriter, r *http.Request) {
// 	idStr := r.URL.Path[len("/item/"):]
// 	id, err := strconv.Atoi(idStr) // convert to int for use later
// 	if err != nil {
// 		http.Error(w, "Invalid ID", http.StatusBadRequest)
// 		return
// 	}
//
// 	switch r.Method {
// 	case http.MethodDelete:
// 		mu.Lock()
// 		for i, item := range items {
// 			if item.ID == id {
// 				items = append(items[:i], items[i+1:]...)
// 				break
// 			}
// 		}
// 		mu.Unlock()
// 		w.WriteHeader(http.StatusNoContent)
// 	case http.MethodPut:
// 		var updatedItem Item
// 		json.NewDecoder(r.Body).Decode(&updatedItem)
// 		updatedItem.ID = id
// 		mu.Lock()
// 		found := false
// 		for i, item := range items {
// 			if item.ID == id {
// 				items[i] = updatedItem
// 				found = true
// 				break
// 			}
// 		}
// 		mu.Unlock()
// 		if !found {
// 			http.Error(w, "Item doesn't exist", http.StatusNotFound)
// 			return
// 		}
// 		json.NewEncoder(w).Encode(updatedItem)
// 	}
// }

*/
