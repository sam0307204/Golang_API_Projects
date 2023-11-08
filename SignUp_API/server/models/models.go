package models

import (
	"fmt"
	dbconect "sam0307/SignUp_API/DB_conect"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golang.org/x/crypto/bcrypt"
)

var db *gorm.DB
var user string

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

func (u *User) FindUsername(U_name string, P_word string) *gorm.DB {
	err := db.Where("username=?", U_name).Find(&u).Error

	switch {
	case err == gorm.ErrRecordNotFound:
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(P_word), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println("Server Error!")
			return
		}

		user = User{Username: U_name}
	}
}
