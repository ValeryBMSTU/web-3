package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scanln(&s)

	maxDigit := '0'

	for _, ch := range s {
		if ch > maxDigit {
			maxDigit = ch
		}
	}

	fmt.Println(string(maxDigit))
}