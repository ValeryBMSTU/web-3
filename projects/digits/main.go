package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)

	max := '0'

	for _, sym := range a {
		if sym > max {
			max = sym
		}
	}

	fmt.Println(string(max))
}
