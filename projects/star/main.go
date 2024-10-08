package main

import "fmt"
import "strings"
func main() {
	var a string
	fmt.Scan(&a)
	var b []string = strings.Split(a,"")
	a = strings.Join(b,"*")
	fmt.Println(a)
}
