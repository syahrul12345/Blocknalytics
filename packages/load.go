package load

import (
	"strings"
	"fmt"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
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

func Start() (uint64,string,uint64,uint64,string,uint64){
	// var blockNumber uint64
	// var networkID string
	// var peers uint64
	// var hashRate string
	// var gasPrice uint64

	blockNumber,blockErr := request("eth_blockNumber",nil)
	if blockErr != nil {
		panic(blockErr)
	}
	currentBlockInfo, blockInfoErr := request("eth_getBlockByNumber",[]interface{}{blockNumber,true})
	if blockInfoErr != nil {
		panic(blockInfoErr)
	}
	currentMap := stringToMap(currentBlockInfo)
	currentBlock := toInt(currentMap["currentBlock"])
	
	prevBlock := "0x" + strconv.FormatInt(int64(currentBlock - 500), 16)
	prevBlockInfo , prevBlockErr := request("eth_getBlockByNumber",[]interface{}{prevBlock,true})
	if prevBlockErr != nil {
		panic(prevBlockErr)
	}
	


	networkID,networkErr := request("net_version",nil)
	if networkErr != nil {
		panic(networkErr)
	}

	peers,peersErr := request("net_peerCount",nil) 
	if peersErr != nil {
		panic(peersErr)
	}


	gasPrice, gasPriceErr := request("eth_gasPrice",nil)
	if gasPriceErr != nil {
		panic(gasPriceErr)
	}

	syncStatus, syncErr := request("eth_syncing",nil)
	if syncErr != nil {
		panic(syncErr)
	}
	//calculate hashrate
	prevMap := stringToMap(prevBlockInfo)
	//calculate the difficulties
	currentDifficulty := toInt(currentMap["currentDifficulty"])
	currentTimeStamp := toInt(currentMap["currentTimeStamp"])
	prevTimeStamp := toInt(prevMap["currentTimeStamp"])

	blocktime := (currentTimeStamp - prevTimeStamp)/500
	hashRate := currentDifficulty/blocktime

	return toInt(blockNumber),networkID,toInt(peers),weiToGwei(toInt(gasPrice)),syncStatus,hashRate
}

func weiToGwei(value uint64) uint64 {
	var gwei = value/1000000000
	return gwei
}

func splitter(value uint64) {
	names := make(map[string]uint64)

	names["kilo"] = 1000
	names["mega"] = 1e+6
	names["giga"] = 1e+9
	names["tera"] = 1e+12
	names["penta"] = 1e+15

	count := make(map[string]uint64)
	count["kilo"] = 0
	count["mega"] = 0
	count["giga"] = 0
	count["tera"] = 0
	count["penta"] = 0
	
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


func request(method string,params []interface{}) (string,error){
	const ethRPC string = "https://adoring-snyder:humped-muster-device-mousy-bauble-appear@nd-806-802-183.p2pify.com"
	
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
		var currentBlock = fmt.Sprintf("%v", params[0])
		var currentDifficulty = difficultyHex
		var currentTimeStamp = timeStampHex

		var returnValues = make(map[string]string)
		returnValues["currentBlock"] = currentBlock
		returnValues["currentDifficulty"] = currentDifficulty
		returnValues["currentTimeStamp"] = currentTimeStamp
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


