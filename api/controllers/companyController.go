package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

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

var GetCompanyByID = func(writer http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(writer, http.StatusBadRequest, utils.Message(http.StatusBadRequest, "The id must be an integer greater than 0"))
		return
	}

	data := models.GetCompanyByID(uint(id))
	if data == nil {
		utils.Respond(writer, http.StatusNotFound, utils.Message(http.StatusNotFound, "Cannot find the company"))
		return
	}
	resp := utils.Message(http.StatusOK, "success")
	resp["data"] = data
	utils.Respond(writer, http.StatusOK, resp)
}

var UpdateCompanyByID = func(writer http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(writer, http.StatusBadRequest, utils.Message(http.StatusBadRequest, "The id must be an integer greater than 0"))
		return
	}

	company := &models.Company{}
	err = json.NewDecoder(request.Body).Decode(company)
	if err != nil {
		utils.Respond(writer, http.StatusBadRequest, utils.Message(http.StatusBadRequest, "Error while decoding request body"))
		return
	}

	resp := company.UpdateCompanyByID(uint(id))
	if resp["status"] != 200 && resp["status"] == 404 {
		utils.Respond(writer, http.StatusNotFound, resp)
		return
	} else if resp["status"] != 200 {
		utils.Respond(writer, http.StatusInternalServerError, resp)
		return
	}
	utils.Respond(writer, http.StatusOK, resp)
}

var DeleteCompanyByID = func(writer http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(writer, http.StatusBadRequest, utils.Message(http.StatusBadRequest, "The id must be an integer greater than 0"))
		return
	}

	resp := models.DeleteCompanyByID(uint(id))
	if resp["status"] != 200 {
		utils.Respond(writer, http.StatusInternalServerError, resp)
		return
	}
	utils.Respond(writer, http.StatusOK, resp)
}
