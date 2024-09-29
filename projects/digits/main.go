package main

import "fmt"

func getMaxDigit(str string) int{
    maxD := 0
    for _, elem := range str{
        d := int(elem - '0')
        if d > maxD{
            maxD = d 
        }
    }
    return maxD
}

func main() {
    var s string
    _, _ = fmt.Scan(&s)
    fmt.Println(getMaxDigit(s))
}





