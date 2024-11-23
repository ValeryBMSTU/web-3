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

	_, err := strconv.Atoi(s)
	if err == nil {
		c := 0
		for _, i := range s {
			in := int(i - '0')
			c = int((math.Max(float64(c), float64(in))))
		}
		fmt.Println(c)
	}
}
