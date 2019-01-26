package app

import (
	"fmt"
	"net/http"
)


// init is run before the application starts serving.
func init() {
	// Handle all requests with path /hello with the helloHandler function.
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/test", testInterfaceHandler)
}


func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "CesarCorp2Test GCP!")
}


func testInterfaceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is a testing inferface, are you expecting something funny or what?")
}
