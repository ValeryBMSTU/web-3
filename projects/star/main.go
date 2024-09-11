package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	r := make([]rune, 0, len(s)*2)
	for _, c := range s {
		r = append(r, c, '*')
	}
	r = r[:len(r)-1]

	fmt.Println(string(r))
}
