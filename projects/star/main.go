package main

import "fmt"

func main() {
	var input string
	fmt.Scanln(&input)

	output := ""
	for i := 0; i < len(input); i++ {
		output += string(input[i])
		if i < len(input)-1 {
			output += "*"
		}
	}

	fmt.Println(output)
}
