package main

import (
	// "image/color"
	// "math/rand"
	// "gonum.org/v1/plot"
	// "gonum.org/v1/plot/plotter"
	// "gonum.org/v1/plot/vg"
	// "gonum.org/v1/plot/vg/draw"
	"fmt"
	"github.com/syahrul12345/Blocknalytics/packages/load"
	"strconv"
	
)


func main() {
	// Get some random points
	const ethRPC string = "https://adoring-snyder:humped-muster-device-mousy-bauble-appear@nd-806-802-183.p2pify.com"
	
	//latest blockNumber
	blockNumber,blockErr := load.Request(ethRPC,"eth_blockNumber",nil)
	if blockErr != nil {
		panic(blockErr)
	}
	fmt.Println("the latest block is: ",toInt(blockNumber))
	txChan := make(chan int)
	blockChan := make(chan uint64)
	for selectedBlock := toInt(blockNumber) - 99;selectedBlock<toInt(blockNumber);selectedBlock++ {
		blockHex := toHex(selectedBlock)
		go makeRequest(ethRPC,blockHex,txChan,blockChan)
	}
	//wait for all of it to be done
	for i := range(txChan) {
		fmt.Println(i)
	}
	for i := range(blockChan) {
		fmt.Println(i)
	}
	// txInCurrentBlock,txErr := load.GetTransactionsInBlock(ethRPC,"eth_getBlockByNumber",[]interface{}{blockNumber,true})
	// if txErr != nil {
	// 	panic(txErr)
	// }
	// fmt.Println(txInCurrentBlock)
	
	// rand.Seed(int64(0))
	// n := 15
	// scatterData := randomPoints(n)
	// lineData := randomPoints(n)
	// linePointsData := randomPoints(n)

	// // Create a new plot, set its title and
	// // axis labels.
	// p, err := plot.New()
	// if err != nil {
	// 	panic(err)
	// }
	// p.Title.Text = "Points Example"
	// p.X.Label.Text = "X"
	// p.Y.Label.Text = "Y"
	// // Draw a grid behind the data
	// p.Add(plotter.NewGrid())

	// // Make a scatter plotter and set its style.
	// s, err := plotter.NewScatter(scatterData)
	// if err != nil {
	// 	panic(err)
	// }
	// s.GlyphStyle.Color = color.RGBA{R: 255, B: 128, A: 255}

	// // Make a line plotter and set its style.
	// l, err := plotter.NewLine(lineData)
	// if err != nil {
	// 	panic(err)
	// }
	// l.LineStyle.Width = vg.Points(1)
	// l.LineStyle.Dashes = []vg.Length{vg.Points(5), vg.Points(5)}
	// l.LineStyle.Color = color.RGBA{B: 255, A: 255}

	// // Make a line plotter with points and set its style.
	// lpLine, lpPoints, err := plotter.NewLinePoints(linePointsData)
	// if err != nil {
	// 	panic(err)
	// }
	// lpLine.Color = color.RGBA{G: 255, A: 255}
	// lpPoints.Shape = draw.PyramidGlyph{}
	// lpPoints.Color = color.RGBA{R: 255, A: 255}

	// // Add the plotters to the plot, with a legend
	// // entry for each
	// p.Add(s, l, lpLine, lpPoints)
	// p.Legend.Add("scatter", s)
	// p.Legend.Add("line", l)
	// p.Legend.Add("line points", lpLine, lpPoints)

	// // Save the plot to a PNG file.
	// if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
	// 	panic(err)
	// }
}

func makeRequest(ethRPC string,blockHex string,txChan chan<-int,blockChan chan<- uint64) {
	txInCurrentBlock,txErr := load.GetTransactionsInBlock(ethRPC,"eth_getBlockByNumber",[]interface{}{blockHex,true})
	if txErr != nil {
		panic(txErr)
	}
	fmt.Println(toInt(blockHex))
	blockChan <- toInt(blockHex)
	txChan <- len(txInCurrentBlock)
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