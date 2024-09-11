package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	var b int
	fmt.Scanf("%d %d", &a, &b)
	fmt.Print(math.Sqrt(float64(a*a + b*b)))
}
