package main

import (
	"fmt"
	"strconv"
)

func main() {
	var inp string
	fmt.Print("Введите целое число: ")
	fmt.Scan(&inp)
	var max int
	max, _ = strconv.Atoi(string(inp[0]))

	for i := 0; i < len(inp); i++ {
		el, _ := strconv.Atoi(string(inp[i]))
		if el > max {
			max = el
		}
	}

	fmt.Print("Результат: ")
	fmt.Println(max)
}
