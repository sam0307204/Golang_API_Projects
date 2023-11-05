package router

import (
	"sam0307/SignUp_API/conrtoller"

	"github.com/gorilla/mux"
)

var RegisterNewUser = func(router *mux.Router) {
	router.HandleFunc("/signup", conrtoller.CreateUser).Methods("POST")
	// router.HandleFunc("")
}
