package main

import (
	"fmt"
	"log"
	"net/http"
	"sam0307/SignUp_API/router"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	router.RegisterNewUser(r)
	http.Handle("/", r)
	fmt.Println("Welcome to  signup ....")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", r))
}
