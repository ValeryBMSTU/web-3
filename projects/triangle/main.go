package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Введите катеты:")
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(math.Sqrt(float64(a*a + b*b)))
}
