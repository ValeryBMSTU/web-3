package main

import (
	"fmt"
	"math"
)

func T(k float64, p float64, v float64) float64 {
	return 6 / W(k, p, v)
}

func W(k float64, p float64, v float64) float64 {
	return math.Sqrt(k / M(p, v))
}

func M(p float64, v float64) float64 {
	return p * v
}

func main() {
	var k, p, v float64
	fmt.Scan(&k, &p, &v)

	fmt.Println(T(k, p, v))
}
