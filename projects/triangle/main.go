package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b)
	c = math.Hypot(a, b)
	fmt.Println(c)
}
