package plot

import (
	"fmt"
	"strconv"

	"github.com/syahrul12345/Blocknalytics/desktop/packages/load"
)

type TxCountStruct struct {
	BlockNumber uint64
	TxCount     int
	GasUsed     uint64
}

func Plot() (map[uint64]int, map[uint64]uint64, error) {
	// Get some random points
	const ethRPC string = "https://hopeful-fermat:snout-quail-embark-update-halved-yonder@nd-369-037-227.p2pify.com"
	transaction99MapVal, gasUsed99MapVal, err := create(ethRPC)
	if err != nil {
		return nil, nil, err
	}
	return transaction99MapVal, gasUsed99MapVal, err
	// Create a new plot, set its title and
	// axis labels.
}

func create(ethRPC string) (transaction99Map map[uint64]int, gasUsed99Map map[uint64]uint64, err error) {
	//latest blockNumber
	blockNumber, blockErr := load.Request(ethRPC, "eth_blockNumber", nil)
	if blockErr != nil {
		return nil, nil, blockErr
	}
	txChan := make(chan TxCountStruct)
	for selectedBlock := toInt(blockNumber) - 99; selectedBlock < toInt(blockNumber); selectedBlock++ {
		blockHex := toHex(selectedBlock)
		go makeRequest(ethRPC, blockHex, txChan)
	}
	transaction99Map = make(map[uint64]int)
	gasUsed99Map = make(map[uint64]uint64)
	var txChanCount = 1
	for currentMap := range txChan {
		if txChanCount == 99 {
			transaction99Map[currentMap.BlockNumber] = currentMap.TxCount
			gasUsed99Map[currentMap.BlockNumber] = currentMap.GasUsed
			close(txChan)
		} else {
			txChanCount++
			transaction99Map[currentMap.BlockNumber] = currentMap.TxCount
			gasUsed99Map[currentMap.BlockNumber] = currentMap.GasUsed
		}
	}
	return transaction99Map, gasUsed99Map, nil
	//if we want to use a plotter
	// pts := make(plotter.XYs, len(transaction99Map))
	// v := make(plotter.Values, 99)
	// var keys []int
	// //sorts the keys
	// for key := range transaction99Map {
	// 	keys = append(keys,int(key))
	// }
	// sort.Ints(keys)
	// var ptsCount = 0
	// for _,key := range keys {
	// 	pts[ptsCount].X = float64(key)
	// 	pts[ptsCount].Y = float64(transaction99Map[uint64(key)])
	// 	v[ptsCount] = float64(gasUsed99Map[uint64(key)])/(8*10e6)*100/9.5
	// 	ptsCount++
	// }

}

func makeRequest(ethRPC string, blockHex string, txChan chan TxCountStruct) {
	txInCurrentBlock, gasUsed, txErr := load.GetTransactionsInBlock(ethRPC, "eth_getBlockByNumber", []interface{}{blockHex, true})
	if txErr != nil {
		fmt.Println(txErr)
		fmt.Println("try again...")
	}
	var tempStruct = new(TxCountStruct)
	tempStruct.BlockNumber = toInt(blockHex)
	tempStruct.TxCount = len(txInCurrentBlock)
	tempStruct.GasUsed = toInt(gasUsed)
	txChan <- *tempStruct
}

func toInt(data string) uint64 {
	runes := []rune(data)
	data = string(runes[2:len(runes)])
	b, err := strconv.ParseUint(data, 16, 64)
	if err != nil {
		fmt.Println(err)
		fmt.Print("try again")
	}
	return uint64(b)
}

func toHex(block uint64) string {
	h := fmt.Sprintf("%x", block)
	return "0x" + h
}

// randomPoints returns some random x, y points.
// func randomPoints(n int) plotter.XYs {
// 	pts := make(plotter.XYs, n)
// 	for i := range pts {
// 		if i == 0 {
// 			pts[i].X = rand.Float64()
// 		} else {
// 			pts[i].X = pts[i-1].X + rand.Float64()
// 		}
// 		pts[i].Y = pts[i].X + 10*rand.Float64()
// 	}
// 	return pts
// }
