package main

import (
	"bshop/app/users/controllers"
	"bshop/app/users/middleware"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var (
		domen = "http://"
		host  = "127.0.0.1"
		port  = ":8001"
	)
	mux := http.NewServeMux()
	fmt.Println("Server started")
	fmt.Println(domen + host + port)
	// URLS
	mux.Handle("/api/user/registration",
		middleware.CheckJSONMiddleware(http.HandlerFunc(controllers.UserRegistrationController)))
	//
	mux.Handle("/api/user/update", middleware.CheckJSONMiddleware(nil))
	//
	mux.Handle("/api/user/delete", middleware.CheckJSONMiddleware(nil))
	// Run Server
	err := http.ListenAndServe(host+port, mux)
	if err != nil {
		fmt.Printf("Server starting error %s\n", err)
		log.Fatal(err)
	}
}
