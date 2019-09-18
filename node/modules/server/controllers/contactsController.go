package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/syahrul12345/Blocknalytics/node/modules/server/utils"

	"github.com/syahrul12345/Blocknalytics/node/modules/server/models"
)

//CreateContact comment
var CreateContact = func(writer http.ResponseWriter, request *http.Request) {
	user := request.Context().Value("user").(uint)
	contact := &models.Contact{}
	err := json.NewDecoder(request.Body).Decode(contact)
	if err != nil {
		utils.Respond(writer, utils.Message(false, "Error while decoding request body"))
		return
	}
	contact.UserID = user
	resp := contact.Create()
	utils.Respond(writer, resp)
}

//GetContactsFor comment
var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//The passed path parameter is not an integer
		utils.Respond(w, utils.Message(false, "There was an error in your request"))
		return
	}

	data := models.GetContacts(uint(id))
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}
