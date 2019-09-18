package model

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/syahrul12345/Blocknalytics/node/modules/analytics/utils"
	"github.com/syahrul12345/Blocknalytics/node/modules/requests"
)

//Account stores the data of the account to be queried by the server
type Account struct {
	Address          string
	TransactionCount string
	Transactions     []Transaction
}

//Transaction represents all transactions that this acocunt has made to.
type Transaction struct {
	AddressTo string
	Value     string
}

//Validate ensures that the account being sent to the server is a valid account
func (acc *Account) Validate() (map[string]interface{}, bool) {
	pattern := "^0x[0-9a-fA-F]{40}$"
	match, _ := regexp.MatchString(pattern, acc.Address)
	//check if it's a correct format
	if !match {
		return utils.Message(false, "Account sent is not of a valid format"), false
	}
	fmt.Println(match)
	return utils.Message(false, "Account is a valid Account"), true
}

//Get calls for the transactional history of the account
func (acc *Account) Get() map[string]interface{} {
	response, ok := acc.Validate()
	if !ok {
		return response
	}
	jsonRaw := requests.Request("eth_getBlockByNumber", []interface{}{"latest", true})
	resultTransactions := &requests.ResultTransactions{}
	parseOK := json.Unmarshal(jsonRaw, resultTransactions)
	if parseOK != nil {
		return utils.Message(false, "Failed parsing block info obtained from the node")
	}
	response["block"] = resultTransactions
	return response
}
