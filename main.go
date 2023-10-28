package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/login", Login)
	http.HandleFunc("/signup", Signup)
	http.HandleFunc("/login_user", LoginBtnPress)
	http.HandleFunc("/signup_user", SignupBtnPress)

	http.ListenAndServe(":8000", nil)
}
