package main
import "fmt"
import "math"
func main() {
var a, b int
fmt.Scan(&a, &b)
fmt.Print(math.Sqrt(float64(a * a + b * b)))
}