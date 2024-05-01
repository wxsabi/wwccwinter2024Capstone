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

	// This will handle the users
	http.HandleFunc("/user", handlers.UserHandler)

	// This will handle the signup
	http.HandleFunc("/signup", handlers.SignupHandler)

	// This will start the web server
	fmt.Println("Server is running on port: 8888") // this has to go before ListenAndServe
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err)
	}
}
