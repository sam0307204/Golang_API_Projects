package models

import (
	dbconect "sam0307/SignUp_API/DB_conect"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type User struct {
	gorm.Model
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func init() {
	dbconect.Connect()
	db = dbconect.GetDB()
	db.AutoMigrate(&User{})
}

func (u *User) CreateUsers() *User {
	db.NewRecord(u)
	db.Create(u)
	return u
}
