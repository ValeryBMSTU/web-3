package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Scanln(&input)

	maxDigit := 0
	for _, char := range input {
		digit, _ := strconv.Atoi(string(char))
		if digit > maxDigit {
			maxDigit = digit
		}
	}

	fmt.Println(maxDigit)
}
