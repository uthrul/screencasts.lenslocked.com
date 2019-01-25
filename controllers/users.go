package controllers

import (
	"fmt"
	"net/http"

	"lenslocked.com/views"
)

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
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

type SignupForm struct {
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

// POST / signup
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form)

	// 	r.PostForm = map[string][]string
	// 	fmt.Fprintln(w, r.PostForm["email"])
	// 	fmt.Fprintln(w, r.PostFormValue("email"))
	// 	fmt.Fprintln(w, r.PostForm["password"])
	// 	fmt.Fprintln(w, r.PostFormValue("password"))
}

// POST / Update
func (u *Users) Update(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	fmt.Fprintln(w, form)
}
