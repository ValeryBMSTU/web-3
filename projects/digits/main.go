package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)

	max := 0
	for _, digit := range str {
		d := int(digit) - '0'
		if d > max {
			max = d
		}
	}

	fmt.Println(max)
}
