package users

import (
	"github.com/gorilla/sessions"
	"github.com/satori/go.uuid"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)


func LoginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")

	if r.Method == "GET" {
		t, _ := template.ParseFiles("users/templates/login.gtpl")
		w.Header().Set("Content-Type", "text/html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		user,err := getUser(strings.Join(r.Form["username"],""))

		if err!=nil{
			log.Fatal(err)
		}
		session.Values["authenticated"] = true
		session.Values["username"] = user.Username
		session.Values["email"] = user.Email

		session.Save(r, w)
		http.Redirect(w, r, "/", 302)
	}
}


func SignInHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		t, _ := template.ParseFiles("users/templates/signin.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		user:=&User{
			uuid.NewV4().String(),
			strings.Join(r.Form["username"],""),
			strings.Join(r.Form["password"],""),
			strings.Join(r.Form["email"],"")}

		user,err:=createUser(user)
		if err!=nil{
			log.Fatal(err)
		}

		http.Redirect(w, r, "/login", 302)

	}
}


func LogOutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "cookie-name")
	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Values["username"] = nil
	session.Values["email"] = nil
	session.Save(r, w)
	http.Redirect(w, r, "/", 302)

}
