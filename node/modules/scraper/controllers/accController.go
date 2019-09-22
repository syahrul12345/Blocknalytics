package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
	"github.com/syahrul12345/Blocknalytics/node/modules/scraper/models"
	"github.com/syahrul12345/Blocknalytics/node/modules/scraper/utils"
)

//Payload represents the struct of the incoming payload
type Payload struct {
	Address string
}

//GetTransactionsOfAccount retrieves all the transactions of an account
var GetTransactionsOfAccount = func(writer http.ResponseWriter, request *http.Request) {
	//First lets decode the request
	payload := &Payload{}
	err := json.NewDecoder(request.Body).Decode(payload)

	if err != nil {
		utils.Respond(writer, utils.Message(false, "The Payload is in the wrong format"))
	}
	account := &models.Account{}
	scanErr := models.GetDB().Table("accounts").Where("address = ?", payload.Address).First(account).Error
	if scanErr != nil {
		if scanErr == gorm.ErrRecordNotFound {
			utils.Respond(writer, utils.Message(false, "This Account has yet to be indexed....."))
		}
	}
	resp := utils.Message(true, "Sucess")
	resp["transactions"] = stringArrayToBytesArray(account.TxRaw)
	utils.Respond(writer, resp)
}

func stringArrayToBytesArray(transactionOut pq.ByteaArray) []models.Transaction {
	txResponse := []models.Transaction{}
	for _, item := range transactionOut {
		tx := &models.Transaction{}
		json.Unmarshal(item, tx)
		txResponse = append(txResponse, *tx)
	}
	return txResponse
}
