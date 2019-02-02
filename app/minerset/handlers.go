package minerset

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/satori/go.uuid"
	"html/template"
	"log"
	"net/http"
	"strings"
	"github.com/gorilla/mux"
	"time"
)

var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key = []byte("super-secret-key")
	store = sessions.NewCookieStore(key)
)


type Context struct {
	List      []Minerset

}

func CreateMinersetHandler(w http.ResponseWriter, r *http.Request) {


	if r.Method == "GET" {
		t, _ := template.ParseFiles("minerset/templates/create.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		session, _ := store.Get(r, "cookie-name")

		mine:=&Minerset{ID:uuid.NewV4().String(),
			Title:strings.Join(r.Form["title"],""),
			Owner:session.Values["username"].(string),
			Url:"/push/?title="+strings.Join(r.Form["title"],""),
			Date:time.Now()	}




		mineset,err:=createMinerset(mine)
		fmt.Print("Created mine set:")
		fmt.Print(mineset)
		if err!=nil{
			log.Fatal(err)
		}

		http.Redirect(w, r, "/", 302)

	}
}

func ListMinersetHandler(w http.ResponseWriter, r *http.Request) {


	if r.Method == "GET" {
		mine,err:=listMinersets()
		fmt.Print("LIST:")
		fmt.Print(mine)
		if err!=nil{
			log.Fatal(err)
		}

		t, _ := template.ParseFiles("minerset/templates/list.gtpl")
		t.Execute(w, Context{List: mine})
	} else {

		http.Redirect(w, r, "/", 302)

	}
}


func PushToMinerSetHandler(w http.ResponseWriter, r *http.Request) {


	if r.Method == "GET" {
		vars := mux.Vars(r)
		fmt.Print(vars)

		minetitle, _ := r.URL.Query()["title"]

		mine,err:=getMinerset(string(minetitle[0]))
		fmt.Print(mine)
		if err!=nil{
			log.Fatal(err)
		}

		t, _ := template.ParseFiles("minerset/templates/push.gtpl")
		t.Execute(w, mine)
	} else {
	//
		http.Redirect(w, r, "/", 302)
	//
	}
}


