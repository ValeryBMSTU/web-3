package main

import "fmt"
import "math"

var k, p, v float64 = 3, 4, 5

func main(){
	fmt.Println(T())
}


func T() float64{
    return 6 / W()
}

func W() float64{
    return math.Sqrt(k / M())
}

func M() float64{
    return p * v
}


