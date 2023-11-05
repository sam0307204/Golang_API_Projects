package conrtoller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sam0307/SignUp_API/models"
	"sam0307/SignUp_API/utils"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("One User created....")
	CreateUser := &models.User{}
	utils.ParseBody(r, CreateUser)
	b := CreateUser.CreateUsers()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
