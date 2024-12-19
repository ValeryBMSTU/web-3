package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	t := []rune(text)
	max := t[0]
	for i := range t {
		if t[i] > max {
			max = t[i]
		}
	}
	fmt.Println(string(max))
}
