package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b int
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)

	hypotenuse := math.Sqrt(float64(a*a + b*b))

	fmt.Println(hypotenuse)
}
