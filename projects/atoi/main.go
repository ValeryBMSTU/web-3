package main
import "fmt"
func main() {
var n string
fmt.Scan(&n)
for i := 0; i < len(n); i++ {
fmt.Print((n[i] - 48) * (n[i] - 48))
}
}
