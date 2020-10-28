package models

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"

	utils "project-demo/api/utils"
)

type User struct {
	gorm.Model
	Nickname  string `gorm:"size:255;not null;unique" json:"nickname"`
	Email     string `gorm:"size:100;not null;unique" json:"email"`
	Password  string `gorm:"size:100;not null;" json:"password"`
	CompanyID uint   `json:"company_id"`
}

func (user *User) Validate() (map[string]interface{}, bool) {

	if user.Nickname == "" {
		return utils.Message(http.StatusBadRequest, "User nickname is required"), false
	}

	if user.Email == "" {
		return utils.Message(http.StatusBadRequest, "Email is required"), false
	}

	if user.Password == "" {
		return utils.Message(http.StatusBadRequest, "Password is required"), false
	}

	if user.CompanyID <= 0 {
		return utils.Message(http.StatusBadRequest, "Comapany ID must be greater than 0"), false
	}

	return utils.Message(http.StatusOK, "success"), true
}

func (user *User) Create() map[string]interface{} {

	if resp, ok := user.Validate(); !ok {
		return resp
	}

	err := GetDB().Create(user).Error
	if err != nil {
		resp := utils.Message(http.StatusInternalServerError, "Error while creating user: "+err.Error())
		return resp
	}
	resp := utils.Message(http.StatusOK, "success")
	resp["user"] = user
	return resp

}

func GetUsers() []*User {

	users := make([]*User, 0)
	err := GetDB().Table("users").Find(&users).Error
	if err != nil {
		fmt.Println("Cannot get the users: ", err)
		return nil
	}

	return users
}
