package mainservice

import (
	"app/users"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
)


var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)

var indexTemplete *template.Template
var userdata users.User
func Index(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	if r.Method == "GET" {
		// Check if user is authenticated
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			userdata=users.User{",","","",""}
		}else{
			username:=session.Values["username"].(string)
			useremail:=session.Values["email"].(string)

			userdata=users.User{",",username,"",useremail}

		}

		indexTemplete, _ = template.ParseFiles("mainservice/templates/index.tmpl","mainservice/templates/header.tmpl")

		indexTemplete.Execute(w, userdata)
	}

}





