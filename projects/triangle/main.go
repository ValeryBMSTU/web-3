package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		var input string = scanner.Text()
		a, _ := strconv.Atoi(strings.Split(input, " ")[0])
		b, _ := strconv.Atoi(strings.Split(input, " ")[1])
		fmt.Println(math.Sqrt(float64(a * b)))
	}
}
