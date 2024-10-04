package main

import (
	"fmt"
	"math"
)

func main() {
	var a float64
	var b float64

	fmt.Scan(&a)
	fmt.Scan(&b)

	fmt.Println(math.Sqrt((a*a + b*b)))
}
