package main
import (
	"fmt"
)

func maxCount(banned []int, n int, maxSum int) int {
	m := make(map[int]bool)
	for _, num := range banned {
		m[num] = true
	}

	result := int(0)
	for i := 1; i <= n; i++ {
		if m[i] {
			continue
		}

		if maxSum - i < 0 {
			return result
		}

		maxSum -= i
		result++
	}

	return result
}

func main() {
	// result: 2
	// banned := []int{1,6,5}
	// n := int(5)
	// maxSum := int(6)

	// result: 0
	// banned := []int{1,2,3,4,5,6,7}
	// n := int(8)
	// maxSum := int(1)

	// result: 7
	banned := []int{11}
	n := int(7)
	maxSum := int(50)

	// result: 
	// banned := []int{}
	// n := int(0)
	// maxSum := int(0)

	result := maxCount(banned, n, maxSum)
	fmt.Printf("result = %v\n", result)
}

