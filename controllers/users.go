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

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	// click sign up on web
	fmt.Println("Inside the New method!!!")
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "this bug")
}
