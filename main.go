package main

import (
	"fmt"
	"math/rand"

	"github.com/goml/gobrain"
)

func main() {
	// set the random seed to 0
	rand.Seed(0)

	// create the XOR representation patter to train the network
	patterns := [][][]float64{
		{{0, 0, 0, 0, 0}, {0}},
		{{0, 0, 0, 0, 1}, {1}},
		{{0, 0, 0, 1, 0}, {2}},
		{{0, 0, 0, 1, 1}, {3}},
		{{0, 0, 1, 0, 0}, {4}},
		{{0, 0, 1, 0, 1}, {5}},
		{{0, 0, 1, 1, 0}, {6}},
		{{0, 0, 1, 1, 1}, {7}},
		{{0, 1, 0, 0, 0}, {8}},
		{{0, 1, 0, 0, 1}, {9}},
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
	ff.Train(patterns, 1000, 0.6, 0.4, false)

	// testing the network
	ff.Test(patterns)

	// predicting a value
	inputs := []float64{0, 1, 0, 0, 1}
	output := ff.Update(inputs)
	fmt.Printf("%g\n", output)
}
