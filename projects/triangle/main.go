package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Println("Add the length of katets:")
	fmt.Scan(&a, &b)
	fmt.Print(math.Sqrt(a*a + b*b))
}
