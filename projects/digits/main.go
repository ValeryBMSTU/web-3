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
		var max uint = 0
		for _, val := range input {
			a := uint(val - '0')
			if a > max {
				max = a
			}
		}
		fmt.Println(max)
	}
}
