package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
    fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println(math.Sqrt(b*b+a*a))
}
