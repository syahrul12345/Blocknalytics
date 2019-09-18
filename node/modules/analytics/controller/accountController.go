package controller

import (
	"encoding/json"
	"net/http"

	model "github.com/syahrul12345/Blocknalytics/node/modules/analytics/models"
	"github.com/syahrul12345/Blocknalytics/node/modules/analytics/utils"
)

//GetAccountInfo gets the transaction records of the account
var GetAccountInfo = func(writer http.ResponseWriter, request *http.Request) {
	//lets get the address from the request
	account := &model.Account{}
	err := json.NewDecoder(request.Body).Decode(account)
	if err != nil {
		utils.Respond(writer, utils.Message(false, "Invalid JSON request"))
	}
	resp := account.Get()
	utils.Respond(writer, resp)
}
