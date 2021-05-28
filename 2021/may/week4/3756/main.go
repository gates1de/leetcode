package main
import (
	"fmt"
	"strconv"
)

func minPartitions(n string) int {
	max := rune('0')
	for _, r := range n {
		if max < r {
			max = r
		}
	}
	result, _ := strconv.Atoi(string(max))
	return result
}

func main() {
	// result: 3
	// n := "32"

	// result: 8
	// n := "82734"

	// result: 9
	n := "27346209830709182346"

	// result: 
	// n := ""

	result := minPartitions(n)
	fmt.Printf("result = %v\n", result)
}

