package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Printf("Введите длины катетов треугольника через пробел: ")
	fmt.Scan(&a, &b)

	fmt.Printf("Гипотенуза равна: %.3f\n", math.Hypot(a, b))
}
