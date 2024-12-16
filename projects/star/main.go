package main

import (
	"fmt"
	"strings"
)

func main() {
	var input, re string

	fmt.Scan(&input)

	re = strings.Join(strings.Split(input, ""), "*")

	fmt.Println(re)
}
