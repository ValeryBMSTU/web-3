package main

import (
	"fmt"
)

func main() {
	var st string
	fmt.Scan(&st)
	for i := 0; i < len(st); i++ {
		tmp := int(st[i] - 48)
		fmt.Print(tmp * tmp)
	}
}
