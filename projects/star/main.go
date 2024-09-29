package main

import (
    "fmt"
    "strings"
)

func main() {
    var str string
    fmt.Scan(&str)
    
    res:= strings.Join(strings.Split(str,""),"*")
    
    fmt.Println(res)
}





