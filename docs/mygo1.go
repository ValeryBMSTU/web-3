package main

import (
	"fmt"
	"strconv"
)

func main() {
	maxim := 0
	var input string
	fmt.Scan(&input)
	for _, t := range input {
		k, _ := strconv.Atoi(string(t))
		if k > maxim {
			maxim = k
		}
	}
	fmt.Println(maxim)
}
