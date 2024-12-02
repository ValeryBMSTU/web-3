package main

import "fmt"

func main() {
    var s string
    res := 0
    fmt.Scan(&s)
    for i := 0; i < len(s) - 1; i++{
        var cur, next int = int(s[i]) - 48, int(s[i + 1]) - 48
        res += cur * cur
        if (next > 3){
            res *= 100
        }else{
            res *= 10
        }
    }
    res += (int(s[len(s) - 1]) - 48) * (int(s[len(s) - 1]) - 48)
    fmt.Println(res)
}