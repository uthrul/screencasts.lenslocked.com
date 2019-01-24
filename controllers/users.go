package controllers

import (
	"fmt"
	"net/http"

	"lenslocked.com/views"
)

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

type Users struct {
	NewView *views.View
}

// GET / signup
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	// click sign up on web
	fmt.Println("Inside the New method!!!")
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

// POST / signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	// r.PostForm = map[string][]string
	fmt.Fprintln(w, r.PostForm["email"])
	fmt.Fprintln(w, r.PostFormValue("email"))
	fmt.Fprintln(w, r.PostForm["password"])
	fmt.Fprintln(w, r.PostFormValue("password"))
}
