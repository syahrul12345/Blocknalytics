package main

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"fmt"
	"github.com/syahrul12345/Blocknalytics/packages/load"
	"strconv"
	"sort"

)

type TxCountStruct struct {
	BlockNumber uint64
	TxCount int
	GasUsed uint64
}



func main() {
	// Get some random points
	const ethRPC string = "https://adoring-snyder:humped-muster-device-mousy-bauble-appear@nd-806-802-183.p2pify.com"
	createTransactionCount99(ethRPC)
	
	// Create a new plot, set its title and
	// axis labels.
}

func createTransactionCount99(ethRPC string) {
	//latest blockNumber
	blockNumber,blockErr := load.Request(ethRPC,"eth_blockNumber",nil)
	if blockErr != nil {
		panic(blockErr)
	}
	txChan := make(chan TxCountStruct)
	for selectedBlock := toInt(blockNumber) - 99;selectedBlock<toInt(blockNumber);selectedBlock++ {
		blockHex := toHex(selectedBlock)
		go makeRequest(ethRPC,blockHex,txChan)
	}
	var transaction99Map = make(map[uint64]int)
	var gasUsed99Map = make(map[uint64]uint64)
	var txChanCount = 1
	for currentMap := range(txChan) {
		if txChanCount == 99 {
			transaction99Map[currentMap.BlockNumber] = currentMap.TxCount
			gasUsed99Map[currentMap.BlockNumber] = currentMap.GasUsed
			close(txChan)
		}else{
			txChanCount++
			transaction99Map[currentMap.BlockNumber] = currentMap.TxCount
			gasUsed99Map[currentMap.BlockNumber] = currentMap.GasUsed
		}
	}
	pts := make(plotter.XYs, len(transaction99Map))
	v := make(plotter.Values, 99)
	var keys []int
	//sorts the keys
	for key := range transaction99Map {
		keys = append(keys,int(key))
	}
	sort.Ints(keys)
	var ptsCount = 0
	for _,key := range keys {
		pts[ptsCount].X = float64(key)
		pts[ptsCount].Y = float64(transaction99Map[uint64(key)])
		v[ptsCount] = float64(gasUsed99Map[uint64(key)])
		ptsCount++
	}
	//create new plot
	p, err := plot.New()
	if err != nil {
		panic(err)
	}
	p.Title.Text = "Transaction Count in last 99 blocks"
	p.X.Label.Text = "Block Number"
	p.Y.Label.Text = "Transaction Count"
	// Draw a grid behind the data
	err = plotutil.AddLinePoints(p,pts)
	if err != nil {
		panic(err)
	}
	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
	for i := range v {
		fmt.Println(v[i])
	}
	
	//create new histogram
	p2, err := plot.New()
	if err != nil {
		panic(err)
	}
	p2.Title.Text = "Gas Used in last 99 blocks"
	p2.Y.Label.Text = "Gas used"
	w := vg.Points(99)
	barsA,err := plotter.NewBarChart(v,w)
	if err != nil {
		panic(err)
	}
	p2.Add(barsA)
	if err := p2.Save(20*vg.Inch, 20*vg.Inch, "barchart.png"); err != nil {
		panic(err)
	}
	// h, err := plotter.NewHist(v,99)
	// if err != nil {
	// 	panic(err)
	// }
	// // Normalize the area under the histogram to
	// // sum to one.
	// p2.Add(h)


	// // Save the plot to a PNG file.
	// if err := p2.Save(4*vg.Inch, 4*vg.Inch, "hist.png"); err != nil {
	// 	panic(err)
	// }
}

func makeRequest(ethRPC string,blockHex string,txChan chan TxCountStruct) {
	txInCurrentBlock,gasUsed,txErr := load.GetTransactionsInBlock(ethRPC,"eth_getBlockByNumber",[]interface{}{blockHex,true})
	if txErr != nil {
		panic(txErr)
	}
	var tempStruct = new(TxCountStruct)
	tempStruct.BlockNumber = toInt(blockHex)
	tempStruct.TxCount = len(txInCurrentBlock)
	tempStruct.GasUsed = toInt(gasUsed)
	txChan <- *tempStruct
}

func toInt(data string) uint64{
	runes := []rune(data)
	data = string(runes[2:len(runes)])
	b,err := strconv.ParseUint(data,16,64)
	if err != nil {
		panic(err)
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