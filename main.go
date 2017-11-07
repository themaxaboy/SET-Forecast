package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"

	//"github.com/goml/gobrain"
	"./gobrain"
)

var setIndex = make([]float64, 0)
var setIndexPercent = make([]float64, 0)
var patterns = make([][][]float64, 0)

func readDir(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		//fmt.Println(file.Name())
		readFile(path + file.Name())
	}
}

func readFile(filePath string) {
	// Load a TXT file.
	f, _ := os.Open(filePath)

	// Create a new reader.
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		}

		//fmt.Println(record)
		if record[0] == "SET" {
			//fmt.Printf("%v %v %v\n", record[0],record[1],record[5])
			v := record[5]
			if s, err := strconv.ParseFloat(v, 64); err == nil {
				//fmt.Printf("%v\n", s)
				setIndex = append(setIndex, s)
			}
			break
		}
	}
}

func convertToPatterns() {
	for i := 10; i < len(setIndexPercent); i++ {
		dataSet := [][]float64{[]float64{setIndexPercent[i-10], setIndexPercent[i-9], setIndexPercent[i-8], setIndexPercent[i-7], setIndexPercent[i-6], setIndexPercent[i-5], setIndexPercent[i-4], setIndexPercent[i-3], setIndexPercent[i-2], setIndexPercent[i-1]}, {setIndexPercent[i]}}
		patterns = append(patterns, dataSet)
	}
}

func convertToPercent() {
	for i := 0; i < len(setIndex); i++ {
		if i == 0 {
			setIndexPercent = append(setIndexPercent, 0)
		} else {
			cal := ((setIndex[i] - setIndex[i-1]) / setIndex[i-1]) * 100
			cal = toFixed(cal, 2)
			setIndexPercent = append(setIndexPercent, cal)
		}
	}
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func main() {
	readDir("./set-history_EOD/1975-1999/")
	readDir("./set-history_EOD/2000-2016/")
	readDir("./set-history_EOD/2017-2018/")
	//fmt.Printf("%v\n", setIndex)

	convertToPercent()
	//fmt.Printf("%v\n", setIndexPercent)

	convertToPatterns()
	//fmt.Printf("%v\n", patterns)

	// set the random seed to 0
	rand.Seed(0)

	// instantiate the Feed Forward
	ff := &gobrain.FeedForward{}

	// initialize the Neural Network;
	// the networks structure will contain:
	// 2 inputs, 2 hidden nodes and 1 output.
	ff.Init(10, 10, 1)

	// train the network using the XOR patterns
	// the training will run for 1000 epochs
	// the learning rate is set to 0.6 and the momentum factor to 0.4
	// use true in the last parameter to receive reports about the learning error
	ff.Train(patterns, 100000, 0.6, 0.4, false)

	// testing the network
	/*store = */
	ff.Test(patterns)
	/*for _, p := range store {
		fmt.Printf("%v\n", p)
	}*/

	// predicting a value
	//inputs := []float64{0.42075325, 0.15326072, 0.15768098, -0.39619605, -0.73605319}
	inputs := []float64{-1.41, 0.54, 0.55, 0.41, 0.42, 0.15, 0.16, -0.40, -0.74, -0.03}
	output := ff.Update(inputs)

	fmt.Printf("Output,%g\n", output)
}
