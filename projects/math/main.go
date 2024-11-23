package main

import (
	"fmt"
	"math"
)

func T() float64 {
	return 6 / W()
}

func W() float64 {
	return math.Sqrt(float64(k) / M())
}

func M() float64 {
	return float64(p * v)
}

var k = 1
var p = 2
var v = 5

func main() {
	fmt.Print(T())
}
