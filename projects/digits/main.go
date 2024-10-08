package main

import "fmt"
import "strings"
import "strconv"
func main() {
	var a string
	fmt.Scan(&a)
	
	for i:= 9; i >= 0; i--{
		if strings.Contains(a,strconv.Itoa(i)){
			fmt.Println(i)
			break
		}
		
	}
}
