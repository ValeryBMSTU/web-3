package main

import (
 "fmt"
 "math"
 "strconv"
)

// Определение функций вне main
func M(p, v float64) float64 {
 return p * v
}

func W(k, m float64) float64 {
 return math.Sqrt(k / m)
}

func T(w float64) float64 {
 return 6 / w
}

func main() {
 var k, p, v float64

 // Ввод значений с консоли
 fmt.Print("Введите значение k: ")
 var inputK string
 fmt.Scanln(&inputK)
 k, _ = strconv.ParseFloat(inputK, 64)

 fmt.Print("Введите значение p: ")
 var inputP string
 fmt.Scanln(&inputP)
 p, _ = strconv.ParseFloat(inputP, 64)

 fmt.Print("Введите значение v: ")
 var inputV string
 fmt.Scanln(&inputV)
 v, _ = strconv.ParseFloat(inputV, 64)

 // Вычисления и вывод результата
 m := M(p, v)
 w := W(k, m)
 t := T(w)
 fmt.Println("Период колебаний (t):", t)
}
