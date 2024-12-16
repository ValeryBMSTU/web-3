package main

import (
	"fmt"
	"math"
)

// 全局变量
var k, p, v, m, w, t float64

func M() float64 {
	m = p * v
	return m
}

func W() float64 {
	w = math.Sqrt(k / m)
	return w
}

func T() float64 {
	t = 6 / w
	return t
}

func main() {
	// 输入变量
	fmt.Scan(&k, &p, &v)

	// 计算 m, w, t
	M()
	W()
	T()

	// 输出 t
	fmt.Println(T())
}
