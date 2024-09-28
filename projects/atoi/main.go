package main

import (
 "fmt"
 "strconv"
)

func main() {
 var number int
 fmt.Scanln(&number)

 result := ""
 for _, digit := range strconv.Itoa(number) {
  digitInt, _ := strconv.Atoi(string(digit))
  squaredDigit := digitInt * digitInt
  result += strconv.Itoa(squaredDigit)
 }

 fmt.Println(result)
}





