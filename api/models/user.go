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

func GetUserByID(id uint) *User {

	user := &User{}
	err := GetDB().Table("users").Where("id = ?", id).First(user).Error
	if err != nil {
		fmt.Println("Cannot get the user: ", err)
		return nil
	}
	return user
}

func (user *User) UpdateUserByID(id uint) map[string]interface{} {

	oldUser := GetUserByID(id)
	if oldUser == nil {
		return utils.Message(http.StatusNotFound, "Cannot find the user")
	}

	err := db.Model(&oldUser).Updates(User{Nickname: user.Nickname, Email: user.Email, Password: user.Password, CompanyID: user.CompanyID})
	if err.Error != nil {
		return utils.Message(http.StatusInternalServerError, "There was an issue while trying to update the user")
	}

	return utils.Message(http.StatusOK, "The user was updated")
}

func DeleteUserByID(id uint) map[string]interface{} {

	user := GetUserByID(id)
	if user == nil {
		return utils.Message(http.StatusNotFound, "Cannot find the user")
	}

	err := GetDB().Unscoped().Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "There was an issue trying to delete the user"+err.Error())
	}

	return utils.Message(http.StatusOK, "The user was deleted")
}
