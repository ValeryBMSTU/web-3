package main

import (
	"fmt"
)

func remove(input string) string {
	seen := make(map[rune]bool)
	var result []rune

	for _, char := range input {
		if !seen[char] {
			seen[char] = true
			result = append(result, char)
		}
	}
	return string(result)
}

func biggest(s string) rune {

	max := rune(s[0])
	for i := 1; i < len(s); i++ {
		currentChar := rune(s[i])
		if currentChar > max {
			max = currentChar
		}
	}
	return max
}
func main() {
	var input string
	fmt.Scan(&input)

	re := biggest(remove(input))

	fmt.Println(string(re))
}
