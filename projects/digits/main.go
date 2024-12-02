package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	max := 0
	for i := 0; i < len(s); i++{
		if (int(s[i]) - 48 > max){
			max = int(s[i]) - 48
		}
	}
	fmt.Println(max)
}
