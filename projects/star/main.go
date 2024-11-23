package main

import (
	"fmt"
	"os"
	
)

func main() {
	var str string;
	fmt.Fscan(os.Stdin, &str)
	for i:=0;i<len(str);i++{
		fmt.Print(string(str[i]))
		if i<len(str)-1{
			fmt.Print("*")
		} else{
			fmt.Print(string('\n'))
		}
	}
}
