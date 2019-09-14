package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

const (
	ethRPC string = "https://mainnet.infura.io/v3/2a3f078d3755444b8777a0204e5f694a"
)

//Payload is the onject to be sent to the ethereum node.
type Payload struct {
	JSONRPC string
	Method  string
	Params  []interface{}
	ID      int
}

//ParsedResponse is the response when a payload is sent to the server. This will be used for all responses
type ParsedResponse struct {
	ID      int
	JSONRPC string
	Result  json.RawMessage
}

//Important functions are
// 1) eth.GetBlockNumber
// 2) eth.GetBalance
// 3) eth.GetBlockByNumber
// 4) eth.GetBlockByHash
//a struct is created for each of these API calls as they return different responses

//ResultString is the struct when calling eth.getBlockNumbe or eth.getBalance
//It will return the hexadecimal value
type ResultString struct {
	Result string
}

//ResultTransactions is the struct used when eth.getBlockByNumber is called. It is the same for eth.GetBlockByHash
//It contains an array of TransactionStructs and details of the block such as result
type ResultTransactions struct {
	Number           string
	Hash             string
	ParentHash       string
	Nonce            string
	Sha3Uncles       string
	LogsBloom        string
	TransactionsRoot string
	StateRoot        string
	Miner            string
	Difficulty       string
	TotalDifficulty  string
	ExtraData        string
	Size             string
	GasUsed          string
	TimeStamp        string
	Transactions     []TransactionStruct
}

// TransactionStruct is the struct for each transaction in a block
type TransactionStruct struct {
	BlockHash   string
	BlockNumber string
	From        string
	To          string
	Gas         string
	GasPrice    string
	Hash        string
}

// Request makes a request and returns the raw JSON of the RESULT only
// @Params method : The method to be called to the Ethereum node example: "eth_getBalance"
// @Params params: The params array to be sent to the ethereum node example: []interface{}{"0x44427aa6c87dabb4cea29a3ad111345d895d1d8b", "latest"}.
// Boolean values must not be stringified.
func Request(method string, params []interface{}) json.RawMessage {
	//examples of params
	payload := &Payload{
		JSONRPC: "2.0",
		Method:  method,
		Params:  params,
		ID:      1,
	}
	requestBody, reqestBodyErr := json.Marshal(payload)
	if reqestBodyErr != nil {
		fmt.Println(reqestBodyErr)
	}

	response, responseErr := http.Post(ethRPC, "application/json", bytes.NewBuffer(requestBody))
	if responseErr != nil {
		fmt.Println(responseErr)
	}
	//close payload to prevent leakages
	defer response.Body.Close()

	// read the response
	body, bodyErr := ioutil.ReadAll(response.Body)
	if bodyErr != nil {
		fmt.Println(bodyErr)
	}
	//parse the response
	var parsedResponse = new(ParsedResponse)
	parsedErr := json.Unmarshal(body, &parsedResponse)
	if parsedErr != nil {
		fmt.Print(parsedErr)
	}
	return parsedResponse.Result

	//example of marshalling responses
	// if method == "eth_blockNumber" || method == "eth_getBalance" {
	// 	var resultString = new(ResultString)
	// 	err := json.Unmarshal(parsedResponse.Result, &resultString.Result)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	//the result is in responseStruct.Result
	// 	fmt.Println(resultString.Result)
	// } else {
	// 	var resultTransactions = new(ResultTransactions)
	// 	err := json.Unmarshal(parsedResponse.Result, &resultTransactions)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	//the answer is in resultTransactions, it represents the data in each block
	// 	fmt.Println(resultTransactions.Number)
	// }

}
func toInt(data string) uint64 {
	runes := []rune(data)
	data = string(runes[2:len(runes)])
	b, err := strconv.ParseUint(data, 16, 64)
	if err != nil {
		fmt.Println(err)
	}
	return uint64(b)
}
