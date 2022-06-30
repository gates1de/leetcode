package main
import (
	"fmt"
	"strconv"
)

func minPartitions(n string) int {
	resultChar := rune('0')
	for _, char := range n {
		if resultChar < char {
			resultChar = char
		}
		if resultChar == rune('9') {
			break
		}
	}

	result, _ := strconv.Atoi(string(resultChar))
	return result
}

func main() {
	// result: 3
	// n := "32"

	// result: 8
	// n := "82734"

	// result: 9
	// n := "27346209830709182346"

	// result: 9
	n := "109"

	// result: 
	// n := ""

	result := minPartitions(n)
	fmt.Printf("result = %v\n", result)
}

