package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int
	fmt.Scan(&a)
	fmt.Scan(&b)

	c = int(math.Sqrt(float64(a*a + b*b)))

	fmt.Println(c)
}
