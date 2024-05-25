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

	// This will handle the singin
	http.HandleFunc("/signin", handlers.SigninHandler)

	// This will start the web server
	go func() {
		err := http.ListenAndServe(":8888", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()
	fmt.Println("Server is running on port: 8888")
	select {}
}

/*
This doesn't use https because there is no CA to get a cert from
and it's running locally only for the purposes of showicasing this
project. Otherwise, the function calls would need to be changed to this:

http.ListenAndServeTLS(":8888", "cert.pem", "key.pem", nil)
*/
