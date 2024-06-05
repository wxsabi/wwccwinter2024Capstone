package main

import (
	"fmt"
	"log"
	"net/http"
	"wwccwinter2024Capstone/handlers"
	"wwccwinter2024Capstone/models"
)

func main() {
	// Initialize database
	models.InitDb()

	// This will handle the homepage
	http.HandleFunc("/html/index.html", handlers.IndexHandler)
	http.HandleFunc("/", handlers.IndexHandler)

	// This will handle the postings/items
	http.HandleFunc("/item", handlers.ItemHandler)
	http.HandleFunc("/html/post.html", handlers.PostHtmlHandler)

	// This will verify the cookie
	http.HandleFunc("/verify-cookie", handlers.CookieHandler)

	// This will handle the users
	http.HandleFunc("/user", handlers.UserHandler)

	// This will delete a user
	http.HandleFunc("/deleteUser", handlers.DeleteUserHandler)

	// This will show a specific item
	http.HandleFunc("/seeItem", handlers.SeeItemHandler)
	http.HandleFunc("/html/seeItem.html", handlers.SeeItemHtmlHandler)

	// This will handle the signup
	http.HandleFunc("/signup", handlers.SignupHandler)
	http.HandleFunc("/html/signup.html", handlers.SignupHtmlHandler)

	// This will handle the singin
	http.HandleFunc("/signin", handlers.SigninHandler)
	http.HandleFunc("/html/signin.html", handlers.SigninHtmlHandler)

	// This will log the user out
	http.HandleFunc("/logout", handlers.LogoutHandler)

	// This will handle the all-Items div on index.html
	http.HandleFunc("/allItems", handlers.GetNewestItems)

	// This will handle the all-Items div on index.html
	http.HandleFunc("/userItems", handlers.GetUserItems)

	// This will serve static files from the /images directory
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("images"))))

	// This will server css files from /css directory
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("./css"))))

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
