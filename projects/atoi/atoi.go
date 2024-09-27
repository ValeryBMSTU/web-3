package main

import (
	"fmt"
	"strconv"
)

func main() {
	var inp string
	fmt.Print("Введите целое число: ")
	fmt.Scan(&inp)

	M := make([]int, len(inp))
	for i := 0; i < len(inp); i++ {
		M[i], _ = strconv.Atoi(string(inp[i]))
		M[i] = M[i] * M[i]
	}

	fmt.Print("Результат: ")
	for i := 0; i < len(M); i++ {
		fmt.Print(M[i])
	}
	fmt.Println()
}
