package main

import "fmt"

func main() {
	var st string
	var m int = 0
	fmt.Scan(&st)
	for i := 0; i < len(st); i++ {
		tmp := int(st[i] - 48)
		if tmp > m {
			m = tmp
		}
	}
	fmt.Println("The biggest cipher:", m)
}
