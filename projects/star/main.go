package main

import "fmt"

func main() {
	var strin, strout string
	fmt.Println("Введите строку:")
	fmt.Scan(&strin)
	for i := 0; i < len(strin); i++ {
		if i == len(strin)-1 {
			strout = strout + string(strin[i])
		} else {
			strout = strout + string(strin[i]) + "*"
		}
	}
	fmt.Println(strout)
}
