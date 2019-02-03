package main

import (
	"app/api"
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/handlers"

)



func main() {

	port := os.Getenv("PORT")
	port="80" //HARDCODE
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	fmt.Println("API End-points::")
	router := api.NewRouter()

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
	fmt.Println("Listenig on ::"+port)

	// Launch server with CORS validations
	log.Fatal(http.ListenAndServe(":" + port, handlers.CORS(allowedOrigins, allowedMethods)(router)))

}