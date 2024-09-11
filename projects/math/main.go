package main

import (
	"fmt"
	"math"
)

func M(p float64, v float64) float64{return p*v}

func W(k float64, p float64, v float64) float64{return math.Sqrt(k/M(p, v))}

func T(k float64, p float64, v float64) float64{return 6/W(k, p, v)}

func main() {
	var k float64
	var p float64
	var v float64
	fmt.Scanf("%f %f %f", &k, &p, &v)

	fmt.Print(T(k, p, v))
}
