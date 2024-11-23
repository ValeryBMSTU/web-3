package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var s string
	fmt.Print("Введите число: ")
	fmt.Fscan(os.Stdin, &s)

	a, err := strconv.Atoi(s)
	if err == nil {
		c := 0
		s := ""
		for i := math.Floor(math.Log10(float64(a))); i >= 0; i-- {
			s += strconv.Itoa((a/int(math.Pow(10, (i))) - c) * (a/int(math.Pow(10, (i))) - c))
			c = (a / int(math.Pow(10, (i)))) * 10
		}
		fmt.Println(s)
	}
}
