package main

import "fmt"
		
import "strconv"


func main() {
	var a int;
	fmt.Scan(&a)
	var rez string = ""
	for  a > 0{
		temp := (a%10)*(a%10)
		rez += strconv.Itoa(temp)
		a/=10
	}
	fmt.Println(rez)
}
