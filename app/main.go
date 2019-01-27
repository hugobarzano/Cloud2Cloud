package main

import (
	"app/mainservice"
	"app/users"
	"log"
	"net/http"
)



func main() {



	http.HandleFunc("/", mainservice.Index)
	http.HandleFunc("/login", users.LoginHandler)
	http.HandleFunc("/signin", users.SignInHandler)
	http.HandleFunc("/logout", users.LogOutHandler)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}