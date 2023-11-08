package conrtoller

import (
	"net/http"
	"sam0307/SignUp_API/models"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "index.html")
	}
	username := r.FormValue("logname")
	password := r.FormValue("logpass")
	res := models.FindUsername(username, password)

	// fmt.Println("One User created....")
	// CreateUser := &models.User{}
	// utils.ParseBody(r, CreateUser)
	// b := CreateUser.CreateUsers()
	// res, _ := json.Marshal(b)
	// w.WriteHeader(http.StatusOK)
	// w.Write(res)
}
