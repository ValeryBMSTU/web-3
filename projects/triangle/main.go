package main

import (
    "fmt"
    "math"
)

func Gepot(a float64, b float64) float64 {
    var c float64 = math.Sqrt(a*a+b*b)
    return c
}

func main() {
    var a, b float64
    fmt.Scanf("%f %f", &a, &b)
    fmt.Println(Gepot(a,b))
}






