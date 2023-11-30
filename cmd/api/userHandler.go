package main

import (
	"fmt"
	"net/http"

	"github.com/lafetz/trivia-go/internal/data"
)

func (app *application) signinHandler(w http.ResponseWriter, r *http.Request) {

}
func (app *application) signupHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name     string `json:"name" validate:"required,max=500"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	user := data.UserModel{
		Name:  input.Name,
		Email: input.Email,
	}
	err, duplicate := data.AddUser(app.db, &user)
	if duplicate {
		fmt.Println("duplicate")
		app.serverErrorResponse(w, r, err)
		return
	}
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusCreated, envelope{"user": user}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
