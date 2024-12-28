package main
import (
	"fmt"
)

func maxScoreSightseeingPair(values []int) int {
	n := len(values)
	if n == 0 {
		return 0
	}

	maxLeftScore := values[0]
	result := int(0)

	for i := 1; i < n; i++ {
		currentRightScore := values[i] - i
		result = max(result, maxLeftScore + currentRightScore)

		currentLeftScore := values[i] + i
		maxLeftScore = max(maxLeftScore, currentLeftScore)
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 11
	// values := []int{8,1,5,2,6}

	// result: 2
	values := []int{1,2}

	// result: 
	// values := []int{}

	result := maxScoreSightseeingPair(values)
	fmt.Printf("result = %v\n", result)
}

