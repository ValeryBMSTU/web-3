package main

import "fmt"

func main() {

	var a string
	fmt.Scan(&a)

	for i := 0; i < len(a); i++ {
		fmt.Print((a[i] - '0') * (a[i] - '0'))
	}

}
