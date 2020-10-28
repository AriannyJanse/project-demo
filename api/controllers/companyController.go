package controllers

import (
	"encoding/json"
	"net/http"

	"project-demo/api/models"
	utils "project-demo/api/utils"
)

var CreateCompany = func(writer http.ResponseWriter, request *http.Request) {

	company := &models.Company{}

	err := json.NewDecoder(request.Body).Decode(company)
	if err != nil {
		utils.Respond(writer, http.StatusBadRequest, utils.Message(http.StatusBadRequest, "Error while decoding request body"))
		return
	}

	resp := company.Create()
	if resp["status"] != 200 {
		utils.Respond(writer, http.StatusInternalServerError, resp)
		return
	}
	utils.Respond(writer, http.StatusOK, resp)
}

var GetAllCompanies = func(writer http.ResponseWriter, request *http.Request) {

	data := models.GetCompanies()
	if data == nil {
		utils.Respond(writer, http.StatusInternalServerError, utils.Message(http.StatusInternalServerError, "Cannot get the companies"))
		return
	}
	resp := utils.Message(http.StatusOK, "success")
	resp["data"] = data
	utils.Respond(writer, http.StatusOK, resp)
}
