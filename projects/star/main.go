package main

import "fmt"

func main() {
	var st, str string
	fmt.Scan(&st)
	for i := 0; i < len(st); i++ {
		if i == len(st)-1 {
			str = str + string(st[i])
			break
		}
		str = str + string(st[i]) + "*"
	}
	fmt.Println(str)
}
