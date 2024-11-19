package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Print("Введите целое число: ")
	var input int
	fmt.Scan(&input)

	var strlast string
	strnew := strconv.Itoa(input)
	for _, char := range strnew {
		a, _ := strconv.Atoi(string(char))
		b := a * a
		strlast += strconv.Itoa(b)

	}
	fmt.Println(strlast)
} 
