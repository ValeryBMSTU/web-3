package main

import "fmt"

func main() {
	var dig string
	fmt.Scan(&dig)

	for _, d := range dig {
		num := int(d - '0')
		num *= num
		fmt.Print(num)
	}
}
