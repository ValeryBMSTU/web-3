package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string
	fmt.Scan(&input)
	a := strings.Join(strings.Split(input, ""), "*")
	fmt.Print(a)
}
