package load

import (
	"strings"
	"fmt"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"github.com/dustin/go-humanize"
)

//this is the payload we will send to our Ethereum node
type Payload struct {
	JsonRpc string 		
	Method string 			
	Params []interface{}
	Id int
}
//this is the result payload
type ParsedResponse struct {
	Id int
	JsonRpc string
	Result string
}

type ParsedResponseBool struct {
	Id int
	JsonRpc string
	Result bool
}

type ParsedResponseReceipt struct {
	Id int
	JsonRpc string
	Result json.RawMessage
}


type ErrorMessage struct {
	Message string
}

type TransactionStruct struct {
		BlockHash string
		BlockNumber string
		From string
		To string
		Gas string
		GasPrice string
		Hash string
}

func Start() (uint64,string,uint64,uint64,string,string,[]TransactionStruct,[]TransactionStruct){
	const ethRPC string = "https://adoring-snyder:humped-muster-device-mousy-bauble-appear@nd-806-802-183.p2pify.com"
	
	blockNumber,blockErr := Request(ethRPC,"eth_blockNumber",nil)
	if blockErr != nil {
		panic(blockErr)
	}
	currentBlockInfo, blockInfoErr := Request(ethRPC,"eth_getBlockByNumber",[]interface{}{blockNumber,true})
	if blockInfoErr != nil {
		panic(blockInfoErr)
	}
	currentMap := stringToMap(currentBlockInfo)
	currentBlock := toInt(currentMap["currentBlock"])
	
	prevBlock := "0x" + strconv.FormatInt(int64(currentBlock - 99), 16)
	prevBlockInfo , prevBlockErr := Request(ethRPC,"eth_getBlockByNumber",[]interface{}{prevBlock,true})
	if prevBlockErr != nil {
		panic(prevBlockErr)
	}
	

	networkID,networkErr := Request(ethRPC,"net_version",nil)
	if networkErr != nil {
		panic(networkErr)
	}

	peers,peersErr := Request(ethRPC,"net_peerCount",nil) 
	if peersErr != nil {
		panic(peersErr)
	}


	gasPrice, gasPriceErr := Request(ethRPC,"eth_gasPrice",nil)
	if gasPriceErr != nil {
		panic(gasPriceErr)
	}

	syncStatus, syncErr := Request(ethRPC,"eth_syncing",nil)
	if syncErr != nil {
		panic(syncErr)
	}
	//calculate hashrate
	var hashRate uint64;
	if blockInfoErr == nil && prevBlockErr == nil {
		prevMap := stringToMap(prevBlockInfo)
		//calculate the difficulties
		currentDifficulty := toInt(currentMap["currentDifficulty"])
		currentTimeStamp := toInt(currentMap["currentTimeStamp"])
		prevTimeStamp := toInt(prevMap["currentTimeStamp"])
		blocktime := (currentTimeStamp - prevTimeStamp)/99
		hashRate = currentDifficulty/blocktime
	}else{
		hashRate = 0
	}

	//get latest transactions in block
	txInCurrentBlock,txErr := GetTransactionsInBlock(ethRPC,"eth_getBlockByNumber",[]interface{}{blockNumber,true})
	if txErr != nil {
		panic(txErr)
	}

	pendingNodeTx,pendErr := GetTransactionsInBlock(ethRPC,"eth_pendingTransactions",nil)
	if pendErr != nil {
		panic(pendErr)
	}
	
	return toInt(blockNumber),networkID,toInt(peers),weiToGwei(toInt(gasPrice)),syncStatus,splitter(hashRate)+"H/s",txInCurrentBlock,pendingNodeTx
}

func weiToGwei(value uint64) uint64 {
	var gwei = value/1000000000
	return gwei
}

func splitter(value uint64) string {
	//split the string into an array
	runes := []rune(humanize.Bytes(value))
	return string(runes[0:(len(runes)-1)])
}

/**
@notice Converts all 0x123423 hexadecimal to the approriate ints
@dev Please use a string input, as json responses from the node is in string input
**/
func toInt(data string) uint64{
	runes := []rune(data)
	data = string(runes[2:len(runes)])
	b,err := strconv.ParseUint(data,16,64)
	if err != nil {
		panic(err)
	}
	return uint64(b)
}


func Request(ethRPC string,method string,params []interface{}) (string,error){
	
	if params == nil {
		params = []interface{}{}
	}

	payload := &Payload{
		JsonRpc: "2.0",
		Method: method,
		Params: params,
		Id: 1,
	}
	//marshal the payload
	requestBody,error := json.Marshal(payload)
	if error != nil {
		return "nil",error
	}
	//send the payload
	response,error := http.Post(ethRPC,"application/json",bytes.NewBuffer(requestBody))
	if error!= nil {
		return "nil",error
	}

	//close the payload
	defer response.Body.Close()

	//read the response
	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		return "nil",error
	}
	
	//convert the response to the desired struct
	if method == "eth_syncing" {
		var parsedResponse = new(ParsedResponseBool)
		err := json.Unmarshal(body, &parsedResponse)
		if(error != nil){
	        return "nil",err
	    }
	    if(parsedResponse.Result == true){
	    	return "Yes",nil
	    }else{
	    	return "No",nil
	    }

	} else if method == "eth_getBlockByNumber" {
		var parsedResponse = new(ParsedResponseReceipt)
		err := json.Unmarshal(body, &parsedResponse)
		if(error != nil){
	        return "nil",err
	    }

	    result := make(map[string]json.RawMessage)
	    err2 := json.Unmarshal(parsedResponse.Result, &result)
	    if err2 != nil {
	    	return "nil",err2
	    }
	    var difficultyHex string
		if err := json.Unmarshal(result["difficulty"], &difficultyHex); err != nil {
			return "nil", err
		}
		var timeStampHex string
		if err := json.Unmarshal(result["timestamp"], &timeStampHex); err != nil {
			return "nil", err
		}
		var returnValues = make(map[string]string)
		returnValues["currentBlock"] = fmt.Sprintf("%v", params[0])
		returnValues["currentDifficulty"] = difficultyHex
		returnValues["currentTimeStamp"] = timeStampHex
		return createKeyValuePairs(returnValues),nil
	}else {
		var parsedResponse = new(ParsedResponse)
		err := json.Unmarshal(body, &parsedResponse)
		if(error != nil){
	        return "nil",err
	    }
	    return parsedResponse.Result,nil
	}
    
}
func GetTransactionsInBlock(ethRPC string,method string,params []interface{}) ([]TransactionStruct,error) {
	if params == nil {
		params = []interface{}{}
	}

	payload := &Payload{
		JsonRpc: "2.0",
		Method: method,
		Params: params,
		Id: 1,
	}
	//marshal the payload
	requestBody,error := json.Marshal(payload)
	if error != nil {
		return []TransactionStruct{},error
	}
	//send the payload
	response,error := http.Post(ethRPC,"application/json",bytes.NewBuffer(requestBody))
	if error!= nil {
		return []TransactionStruct{},error
	}

	//close the payload
	defer response.Body.Close()

	//read the response
	body, error := ioutil.ReadAll(response.Body)
	if error != nil {
		return []TransactionStruct{},error
	}
	var parsedResponse = new(ParsedResponseReceipt)
	err := json.Unmarshal(body, &parsedResponse)
	if(error != nil){
        return []TransactionStruct{},err
    }

    if method == "eth_pendingTransactions" {
    	var transactions = []TransactionStruct{}
    	err2 := json.Unmarshal(parsedResponse.Result, &transactions)
	    if err2 != nil {
	    	return []TransactionStruct{},err2
	    }
	    return transactions,nil

    }else{
    	result := make(map[string]json.RawMessage)
	    err2 := json.Unmarshal(parsedResponse.Result, &result)
	    if err2 != nil {
	    	return []TransactionStruct{},err2
	    }

	    //lets initialzie an array of structs
	    var transactions = []TransactionStruct{}
		err3 := json.Unmarshal(result["transactions"],&transactions)
		if err3 != nil {
			return []TransactionStruct{},err3
		}
	    return transactions,nil
    }
    
}


func createKeyValuePairs(m map[string]string) string {
    b := new(bytes.Buffer)
    for key, value := range m {
        fmt.Fprintf(b, "%s=%s&", key, value)
    }
    return b.String()
}
func stringToMap(s string) map[string]string {
	var ss []string
	ss = strings.Split(s, "&")
	var m = make(map[string]string)
	for _,pair := range ss {
		z := strings.Split(pair, "=")
		//exclude the empty last index
		if(z[0] != ""){
			m[z[0]] = z[1]
		}
	}
	return m
}


