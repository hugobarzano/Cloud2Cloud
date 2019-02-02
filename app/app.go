package main

import (
	"app/mainservice"
	"app/minerset"
	"app/users"
	"fmt"
	"log"
	"net/http"
	"os"
)



func main() {

	port := os.Getenv("PORT")
	port="80"
	if port == "" {
		log.Fatal("$PORT must be set")
	}


	http.HandleFunc("/", mainservice.Index)
	http.HandleFunc("/login", users.LoginHandler)
	http.HandleFunc("/signin", users.SignInHandler)
	http.HandleFunc("/logout", users.LogOutHandler)
	http.HandleFunc("/create", minerset.CreateMinersetHandler)
	http.HandleFunc("/list", minerset.ListMinersetHandler)
	http.HandleFunc("/push", minerset.PushToMinerSetHandler)



	fmt.Println("ListenAndServe on :"+port)
	err := http.ListenAndServe(":"+ port, nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}