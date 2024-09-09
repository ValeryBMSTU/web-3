package main

import "fmt"

func main() {
	var str string
	fmt.Scan(&str)
	for index, digit := range str {
		if index == 0 || index == len(str) {
			fmt.Print(string(digit))
		} else {
			fmt.Print("*")
			fmt.Print(string(digit))
		}
	}
}
