package main
import (
	"fmt"
	"strconv"
)

func minOperations(boxes string) []int {
	n := len(boxes)
	result := make([]int, n)

	ballsToLeft := int(0)
	movesToLeft := int(0)
	ballsToRight := int(0)
	movesToRight := int(0)

	for i := 0; i < n; i++ {
		result[i] += movesToLeft
		ballLeft, _ := strconv.Atoi(string(boxes[i]))
		ballsToLeft += ballLeft
		movesToLeft += ballsToLeft

		j := n - 1 - i
		result[j] += movesToRight
		ballRight, _ := strconv.Atoi(string(boxes[j]))
		ballsToRight += ballRight
		movesToRight += ballsToRight
	}

	return result
}

func main() {
	// result: [1,1,3]
	// boxes := "110"

	// result: [11,8,5,4,3,4]
	boxes := "001011"

	// result: []
	// boxes := ""

	result := minOperations(boxes)
	fmt.Printf("result = %v\n", result)
}

