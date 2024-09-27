package main

import (
	"fmt"
)

func main() {
	var imp string
	fmt.Print("Введите строку: ")
	fmt.Scan(&imp)

	var O []byte = make([]byte, len(imp)*(1+1/2), len(imp)*(1+1/2))

	for i := 0; i < len(imp); i++ {
		O = append(O, imp[i])
		if i != (len(imp) - 1) {
			O = append(O, '*')
		}
	}

	fmt.Print("Результат: ")
	for i := 0; i < len(O); i++ {
		fmt.Print(string(O[i]))
	}
	fmt.Println()
}
