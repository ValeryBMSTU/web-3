package main

import "fmt"

func main() {
	var strin string
	fmt.Println("Введите строку из цифр:")
	fmt.Scan(&strin)
	for i := 0; i < len(strin); i++ {
		tmp := int(strin[i] - '0')
		fmt.Print(tmp * tmp)
	}
}
