package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n string
	fmt.Scanln(&n)

	var result string

	for _, digit := range n {
		num, _ := strconv.Atoi(string(digit))
		result += strconv.Itoa(num * num)

	}

	fmt.Println(result)
}


