package main

import "fmt"
import "math"
func main() {
	var a,b float64;
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Println(math.Sqrt(a*a + b*b))
}
