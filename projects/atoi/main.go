package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	fmt.Scan(&input)

	for i := 0; i < len(input); i++ {
		digit, _ := strconv.Atoi(string(input[i]))
		square := digit * digit
		fmt.Print(square)
	}

}
