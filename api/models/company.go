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

func GetCompanyByID(id uint) *Company {

	company := &Company{}
	err := GetDB().Table("companies").Where("id = ?", id).First(company).Error
	if err != nil {
		fmt.Println("Cannot get the company: ", err)
		return nil
	}
	return company
}

func (company *Company) UpdateCompanyByID(id uint) map[string]interface{} {

	oldCompany := GetCompanyByID(id)
	if oldCompany == nil {
		return utils.Message(http.StatusNotFound, "Cannot find the company")
	}

	err := GetDB().Model(&oldCompany).Updates(Company{Name: company.Name}).Error
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "There was an issue while trying to update the company")
	}

	return utils.Message(http.StatusOK, "The company was updated")
}

func DeleteCompanyByID(id uint) map[string]interface{} {

	company := GetCompanyByID(id)
	if company == nil {
		return utils.Message(http.StatusNotFound, "Cannot find the company")
	}

	err := GetDB().Unscoped().Where("id = ?", id).Delete(&Company{}).Error
	if err != nil {
		return utils.Message(http.StatusInternalServerError, "There was an issue trying to delete the company"+err.Error())
	}

	return utils.Message(http.StatusOK, "The company was deleted")
}
