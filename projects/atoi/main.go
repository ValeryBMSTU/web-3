package main

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		var input string = scanner.Text()
		var output string = ""
		for _, val := range input {
			var num = val - '0'
			num = num * num
			output = output + strconv.FormatInt(int64(num), 10)
		}
		fmt.Println(output)
	}
}
