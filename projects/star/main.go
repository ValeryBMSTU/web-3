package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		var input string = scanner.Text()
		var vals []rune

		for idx, val := range input {
			if idx < len(input) - 1 {
				vals = append(vals, val, '*')
			} else {
				vals = append(vals, val)
			}
		}

		fmt.Println(string(vals))
	}	
}
