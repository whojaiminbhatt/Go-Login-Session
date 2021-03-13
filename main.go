package main

import (
	"net/http"

	"controllers/accountcontroller"
)

func main() {

	http.HandleFunc("/account", accountcontroller.Index)
	http.HandleFunc("/account/index", accountcontroller.Index)
	http.HandleFunc("/account/login", accountcontroller.Login)
	http.HandleFunc("/account/welcome", accountcontroller.Welcome)
	http.HandleFunc("/account/logout", accountcontroller.Logout)
	http.ListenAndServe(":3000", nil)
}
