package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

var mapValue = make([][]float64, 0)
var outputValue = make([]float64, 0)

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
		//fmt.Printf("%v %v %v\n", record[0],record[1],record[5])

		if record[0] == "Output" {
			v := record[1]
			v = strings.Replace(v, "[", "", -1)
			v = strings.Replace(v, "]", "", -1)
			if s, err := strconv.ParseFloat(v, 64); err == nil {
				outputValue = append(outputValue, s)
				break
			}
		}

		v := record[1]
		v = strings.Replace(v, "[", "", -1)
		v = strings.Replace(v, "]", "", -1)

		p := record[2]
		p = strings.Replace(p, "[", "", -1)
		p = strings.Replace(p, "]", "", -1)

		dataSet := []float64{}

		if s, err := strconv.ParseFloat(v, 64); err == nil {
			dataSet = append(dataSet, s)
		}

		if m, err := strconv.ParseFloat(p, 64); err == nil {
			dataSet = append(dataSet, m)
		}
		mapValue = append(mapValue, dataSet)
	}
}

func main() {
	readFile("./result.txt")
	//fmt.Printf("%v\n", mapValue)

	myNumber := outputValue[0]
	distance := math.Abs(mapValue[0][0] - myNumber)
	idx := 0
	for index, data := range mapValue {
		cdistance := math.Abs(data[0] - myNumber)
		if cdistance < distance {
			idx = index
			distance = cdistance
		}
		//fmt.Printf("%g\n", data[0])
	}
	theNumber := mapValue[idx]
	fmt.Printf("%g : %g = %g%%\n", myNumber, theNumber[0], theNumber[1])
}
