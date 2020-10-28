package models

import (
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"

	utils "project-demo/api/utils"
)

type Company struct {
	gorm.Model
	Name string `gorm:"size:255;not null;unique" json:"name"`
}

func (company *Company) Validate() (map[string]interface{}, bool) {

	if company.Name == "" {
		return utils.Message(http.StatusBadRequest, "Company name is required"), false
	}

	return utils.Message(http.StatusOK, "success"), true
}

func (company *Company) Create() map[string]interface{} {

	if resp, ok := company.Validate(); !ok {
		return resp
	}

	err := GetDB().Create(company).Error
	if err != nil {
		resp := utils.Message(http.StatusInternalServerError, "Error while creating company: "+err.Error())
		return resp
	}
	resp := utils.Message(http.StatusOK, "success")
	resp["company"] = company
	return resp

}

func GetCompanies() []*Company {

	companies := make([]*Company, 0)
	err := GetDB().Table("companies").Find(&companies).Error
	if err != nil {
		fmt.Println("Cannot get the companies: ", err)
		return nil
	}

	return companies
}
