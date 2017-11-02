package main

import (
	"fmt"
	//"math/rand"
	"encoding/csv"
	"bufio"
    "os"
	"io"
	"io/ioutil"
	"log"
	"strconv"

	//"github.com/goml/gobrain"
)

var setIndex = make([]float64,0)
var setIndexPercent = make([]float64,0)
var patterns = make([][][]float64,0)

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
		if (record[0] == "SET") {
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

func convertToPatterns () {
	for i := 5; i < len(setIndexPercent); i++ {
		patterns = append(patterns, ({0, 1, 2, 3, 4}, {5}) )
	}
}

func convertToPercent () {
	for i := 0; i < len(setIndex); i++ {
		if (i == 0) {
			setIndexPercent = append(setIndexPercent, 0)
		} else {
		cal := (1 - (setIndex[i-1] / setIndex[i])) * 100
		setIndexPercent = append(setIndexPercent, cal)
		}
	}
}

func readDir(path string){
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		//fmt.Println(file.Name())
		readFile(path + file.Name())
	}
}

func main() {
	//readDir("./set-history_EOD/1975-2016/")
	readDir("./set-history_EOD/2016-2017/")
	//fmt.Printf("%v\n", setIndex)

	convertToPercent()
	fmt.Printf("%v\n", setIndexPercent)
	
	convertToPatterns()
	fmt.Printfln("%v\n", patterns)
	
	/*// set the random seed to 0
	rand.Seed(0)

	// create the XOR representation patter to train the network
	patterns := [][][]float64{
		{{0, 1, 2, 3, 4}, {5}},
		{{0, 1, -1, 1, -1}, {0}},
		{{0, 5, 4, 3, 2}, {-2}},
		{{0, -2, -3, -4, -5}, {-8}},
		{{0, -8, -6, -2, -3}, {2}},
	}

	// instantiate the Feed Forward
	ff := &gobrain.FeedForward{}

	// initialize the Neural Network;
	// the networks structure will contain:
	// 2 inputs, 2 hidden nodes and 1 output.
	ff.Init(5, 5, 1)

	// train the network using the XOR patterns
	// the training will run for 1000 epochs
	// the learning rate is set to 0.6 and the momentum factor to 0.4
	// use true in the last parameter to receive reports about the learning error
	ff.Train(patterns, 10000, 0.6, 0.4, false)

	// testing the network
	ff.Test(patterns)

	// predicting a value
	inputs := []float64{0, -1, 2, -2, 1}
	output := ff.Update(inputs)
	fmt.Printf("%g\n", output)*/
}
