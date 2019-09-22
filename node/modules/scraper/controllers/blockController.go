package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/syahrul12345/Blocknalytics/node/modules/requests"
	"github.com/syahrul12345/Blocknalytics/node/modules/scraper/models"
)

//GetLatest returns the latest block number
func GetLatest() {
	fmt.Println("Initiating block controller...")
	fmt.Println("Getting lastest block...")
	resp := requests.Request("eth_blockNumber", []interface{}{})
	resultString := &models.ResultString{}
	err := json.Unmarshal(resp, &resultString.Result)
	if err != nil {
		fmt.Println("Error getting block information")
	}
	latestBlock, err := toInt(resultString.Result)
	if err != nil {
		fmt.Println("Managed to get the block, but failed to converted to int")
	}
	currentBlock := 0
	for currentBlock < int(latestBlock) {
		fmt.Println("Saving transactions for block: ", currentBlock)
		GetBlock(toHex(int64(currentBlock)))
		currentBlock++
		fmt.Println("SAVED for block: ", currentBlock)
	}

}

//GetBlock information of given block number
func GetBlock(blockNumber string) {
	resp := requests.Request("eth_getBlockByNumber", []interface{}{blockNumber, true})
	blockStruct := &models.ResultTransactions{}
	err := json.Unmarshal(resp, blockStruct)
	if err != nil {
		fmt.Println("Failed to get transaction info from block")
	}
	//blockStruct.Transactions represent all Transactions in the DB
	for i := range blockStruct.Transactions {
		currentTx := blockStruct.Transactions[i]
		currentTx.Store()
	}

}

func toInt(blockNumber string) (int64, error) {
	runes := []rune(blockNumber)
	runes = runes[2:]
	return strconv.ParseInt(string(runes), 16, 64)
}

func toHex(blockNumber int64) string {
	h := fmt.Sprintf("%x", blockNumber)
	return "0x" + h
}
