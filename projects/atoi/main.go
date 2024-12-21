package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	t := []rune(text)
	for i := 0; i < len(t)-1; i++ {
		r := t[i] - '0'
		fmt.Print(r * r)
	}
}
