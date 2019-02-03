package api

import (
	"app/mainservice"
	"app/minerset"
	"app/users"
	"log"
	"net/http"
	"github.com/gorilla/mux"

)



// Route defines a route
type Route struct {
	Name        string
	Method      []string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route


var routes = Routes {
	//Route {
	//	"Authentication",
	//	"POST",
	//	"/get-token",
	//	controller.GetToken,
	//},
	Route {
		"Index",
		[]string{"GET"},
		"/",
		mainservice.Index,
	},
	Route {
		"login",
		[]string{"GET","POST"},
		"/login",
		users.LoginHandler,
	},
	Route {
		"signin",
		[]string{"GET","POST"},
		"/signin",
		users.SignInHandler,
	},
	Route {
		"logout",
		[]string{"GET"},
		"/logout",
		users.LogOutHandler,
	},
	Route {
		"create",
		[]string{"GET","POST"},
		"/create",
		minerset.CreateMinersetHandler,
	},
	Route {
		"list",
		[]string{"GET"},
		"/list",
		minerset.ListMinersetHandler,
	},
	Route {
	"push",
[]string{"GET"},
	"/push/{title}",
	minerset.PushToMinerSetHandler,
}}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method... ).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}

