package controllers

import (
	"encoding/json"
	"net/http"

	"project-demo/api/models"
	utils "project-demo/api/utils"
)

var CreateUser = func(writer http.ResponseWriter, request *http.Request) {

	user := &models.User{}

	err := json.NewDecoder(request.Body).Decode(user)
	if err != nil {
		utils.Respond(writer, http.StatusBadRequest, utils.Message(http.StatusBadRequest, "Error while decoding request body"))
		return
	}

	resp := user.Create()
	if resp["status"] != 200 {
		utils.Respond(writer, http.StatusInternalServerError, resp)
		return
	}
	utils.Respond(writer, http.StatusOK, resp)
}

var GetAllUsers = func(writer http.ResponseWriter, request *http.Request) {

	data := models.GetUsers()
	if data == nil {
		utils.Respond(writer, http.StatusInternalServerError, utils.Message(http.StatusInternalServerError, "Cannot get the users"))
		return
	}
	resp := utils.Message(http.StatusOK, "success")
	resp["data"] = data
	utils.Respond(writer, http.StatusOK, resp)
}
