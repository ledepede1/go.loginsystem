package main

import (
	users "LoginSystem/Users"
	"net/http"
	"text/template"
)

func Login(wr http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("./Frontend/login.html")).Execute(wr, nil)
}

func Signup(wr http.ResponseWriter, r *http.Request) {
	template.Must(template.ParseFiles("./Frontend/signup.html")).Execute(wr, nil)
}

func LoginBtnPress(wr http.ResponseWriter, r *http.Request) {
	var username = r.PostFormValue("username")
	var password = r.PostFormValue("Password")

	if users.GetUser(username, password) {
		template.Must(template.ParseFiles("./Frontend/index.html")).Execute(wr, nil)
	} else {
		template.Must(template.ParseFiles("./Frontend/error.html")).Execute(wr, nil)
	}
}

func SignupBtnPress(wr http.ResponseWriter, r *http.Request) {
	var username = r.PostFormValue("username")
	var password = r.PostFormValue("Password")

	if users.CreateNewUser(username, password) {
		template.Must(template.ParseFiles("./Frontend/login.html")).Execute(wr, nil)
	} else {
		template.Must(template.ParseFiles("./Frontend/alreadyexists.html")).Execute(wr, nil)
	}
}
