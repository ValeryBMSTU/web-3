package main

import (
	"fmt"
	"math"
)

func T( k float64, p float64, v float64) float64 {
	var t float64 = 6/W(k,p,v)
	return t
}
func M(p float64, v float64) float64 {
	var m float64 = p*v
	return m
}
func W (k float64, p float64, v float64) float64 {
	var w float64 = math.Sqrt(k / M(p,v))
	return w
}

func main() {
	var res, n1, n2, n3 float64
	fmt.Scanf("%f %f %f", &n1, &n2, &n3)
	res = T(n1,n2,n3)
	fmt.Println(res)
}
