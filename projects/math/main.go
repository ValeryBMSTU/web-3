package main

import "fmt"
import "math"
func T(w float64)(float64){
	return (6/w)
}
func M(p,v float64)(float64){
	return (p*v)
}
func W(k,m float64)(float64){
	return math.Sqrt(k/m)
}
func main() {
	var k,p,v float64
	fmt.Scan(&k)
	fmt.Scan(&p)
	fmt.Scan(&v)
	fmt.Println(T(W(k,M(p,v))))
}
