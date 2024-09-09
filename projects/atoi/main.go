package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Scan(&str)

	for _, digit := range str {
		d := int(digit) - '0'
		fmt.Printf("%d", d*d)
	}
}
