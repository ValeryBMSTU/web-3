package main
import "fmt"
func main() {
var myString string
fmt.Scan(&myString)
for i, v := range myString {
if i != 0 {
fmt.Print("*")
}
fmt.Printf("%c", v)
}
}
