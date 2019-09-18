package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/syahrul12345/Blocknalytics/node/modules/server/utils"

	"github.com/syahrul12345/Blocknalytics/node/modules/server/models"
)

//CreateAccount comment
var CreateAccount = func(writer http.ResponseWriter, request *http.Request) {
	account := &models.Account{}
	//check if the payload is in an account struct, if not throw error
	err := json.NewDecoder(request.Body).Decode(account)
	if err != nil {
		utils.Respond(writer, utils.Message(false, "Invalid request"))
	}
	resp := account.Create()
	utils.Respond(writer, resp)

}

// Authenticate comment
var Authenticate = func(writer http.ResponseWriter, request *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(request.Body).Decode(account)
	if err != nil {
		utils.Respond(writer, utils.Message(false, "Invalid request"))
	}
	resp := models.Login(account.Email, account.Password)
	utils.Respond(writer, resp)
}
