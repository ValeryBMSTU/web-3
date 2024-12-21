package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	t := []rune(text)
	for i := 0; i < len(t); i++ {
		if i == len(t)-2 {
			fmt.Print(string(t[i]))
		} else {
			fmt.Print(string(t[i]) + "*")
		}
	}
}
