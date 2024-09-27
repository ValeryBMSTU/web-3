package main

import (
	"fmt"
	"math"
)

func M(p, v float64) float64 {
	return p * v
}

func W(k, m float64) float64 {
	return math.Sqrt(k / m)
}

func T(w float64) float64 {
	return 6 / w
}

func main() {
	var P, V, K float64
	fmt.Print("/* Эта программа вычисляет период колебаний математического маятника */\nВведите значения для P, V и K (в СИ) через пробел и без указания единиц измерения: ")
	fmt.Scan(&P, &V, &K)

	fmt.Printf("Результат: t = %.3f c", T(W(K, M(P, V))))
	fmt.Println()
}
