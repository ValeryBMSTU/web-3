package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Scan(&k, &p, &v)
	M()
	W()
	fmt.Println(T())
}

var k, p, v float64

func M() float64 {
	return p * v
}
func W() float64 {
	return math.Sqrt(k / M())
}
func T() float64 {
	return 6 / W()
}
