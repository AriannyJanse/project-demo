package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

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
	if resp["status"] != 200 && resp["status"] != 400 {
		utils.Respond(writer, http.StatusInternalServerError, resp)
		return
	} else if resp["status"] == 400 {
		utils.Respond(writer, http.StatusBadRequest, resp)
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

var GetUserByID = func(writer http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(writer, http.StatusBadRequest, utils.Message(http.StatusBadRequest, "The id must be an integer greater than 0"))
		return
	}

	data := models.GetUserByID(uint(id))
	if data == nil {
		utils.Respond(writer, http.StatusNotFound, utils.Message(http.StatusNotFound, "Cannot find the user"))
		return
	}
	resp := utils.Message(http.StatusOK, "success")
	resp["data"] = data
	utils.Respond(writer, http.StatusOK, resp)
}

var UpdateUserByID = func(writer http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(writer, http.StatusBadRequest, utils.Message(http.StatusBadRequest, "The id must be an integer greater than 0"))
		return
	}

	user := &models.User{}
	err = json.NewDecoder(request.Body).Decode(user)
	if err != nil {
		utils.Respond(writer, http.StatusBadRequest, utils.Message(http.StatusBadRequest, "Error while decoding request body"))
		return
	}

	resp := user.UpdateUserByID(uint(id))
	if resp["status"] != 200 && resp["status"] == 404 {
		utils.Respond(writer, http.StatusNotFound, resp)
		return
	} else if resp["status"] != 200 {
		utils.Respond(writer, http.StatusInternalServerError, resp)
		return
	}
	utils.Respond(writer, http.StatusOK, resp)
}

var DeleteUserByID = func(writer http.ResponseWriter, request *http.Request) {

	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(writer, http.StatusBadRequest, utils.Message(http.StatusBadRequest, "The id must be an integer greater than 0"))
		return
	}

	resp := models.DeleteUserByID(uint(id))
	if resp["status"] != 200 {
		utils.Respond(writer, http.StatusInternalServerError, resp)
		return
	}
	utils.Respond(writer, http.StatusOK, resp)
}
