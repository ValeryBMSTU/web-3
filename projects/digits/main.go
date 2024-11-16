package main
import "fmt"
func main() {
var myString string
fmt.Scan(&myString)
var mx rune
for _, v := range myString {
if v > mx {
mx = v
}
}
fmt.Printf("%c", mx)
}
