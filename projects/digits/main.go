package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	runes := []rune(s)

	var ans rune = '0'

	for _, a := range runes {
		if a > ans {
			ans = a
		}
	}

	fmt.Printf("%c\n", ans)
}
