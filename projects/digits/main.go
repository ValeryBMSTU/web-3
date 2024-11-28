package main

import "fmt"

func main() {
	var st string
	var m int = 0
	fmt.Println("Введите строку из цифр:")
	fmt.Scan(&st)
	for i := 0; i < len(st); i++ {
		tmp := int(st[i] - '0')
		if tmp > m {
			m = tmp
		}
	}
	fmt.Println("Наибольшая цифра:", m)
}
