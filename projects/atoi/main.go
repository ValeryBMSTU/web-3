package main 

import "fmt"

func  main() {
    var str string
    _, _ = fmt.Scan(&str)
    allToSquare(str)
    fmt.Println()
}

func allToSquare(s string) {
    for _, elem := range s {
        num := int(elem-'0')
        res := num*num
        fmt.Print(res)
    }
}





