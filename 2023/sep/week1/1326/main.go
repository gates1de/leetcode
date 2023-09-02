package main
import (
	"fmt"
)

func minTaps(n int, ranges []int) int {
	maxReach := make([]int, n + 1)

	for i, r := range ranges {
		start := max(0, i - r)
		end := min(n, i + r)

		maxReach[start] = max(maxReach[start], end)
	}

	result := int(0)
	currEnd := int(0)
	nextEnd := int(0)

	for i := 0; i <= n; i++ {
		if i > nextEnd {
			return -1
		}

		if i > currEnd {
			result++
			currEnd = nextEnd
		}

		nextEnd = max(nextEnd, maxReach[i])
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 1
	// n := int(5)
	// ranges := []int{3,4,1,1,0,0}

	// result: -1
	n := int(3)
	ranges := []int{0,0,0,0}

	// result: 
	// n := int(0)
	// ranges := []int{}

	result := minTaps(n, ranges)
	fmt.Printf("result = %v\n", result)
}

