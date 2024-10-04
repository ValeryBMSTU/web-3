package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string

	fmt.Scan(&a)

	res := strings.Join(strings.Split(a, ""), "*")

	fmt.Println(res)
}
