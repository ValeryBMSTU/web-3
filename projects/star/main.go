package main

import "fmt"

func main() {
	var input string
	fmt.Scanln(&input)

	var result string
	for i, ch := range input {
		if i > 0 {
			result += "*"
		}
		result += string(ch)
	}

	fmt.Println(result)
}
