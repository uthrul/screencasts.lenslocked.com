package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	var username string = "sauth"
	w.Header().Set("Content-Type", "text/html")
	if username == "" {
			fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
	} else {
			fmt.Fprint(w, "<h1>Welcome to my awesome site, "+username"!</h1>")
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "To get in touch, please send an email to <a href=\"mailto:sauthrully@gmail.com\">sauthrully@gmail.com</a>.")
}

func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Sorry :( 404</h1>")
}

func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFound)
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":3000", r)
}
