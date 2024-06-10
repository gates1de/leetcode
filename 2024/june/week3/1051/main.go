package main
import (
	"fmt"
	"sort"
)

func heightChecker(heights []int) int {
	expected := make([]int, len(heights))
	copy(expected, heights)
	sort.Ints(expected)

	result := int(0)
	for i, height := range heights {
		if expected[i] != height {
			result++
		}
	}

	return result
}

func main() {
	// result: 3
	// heights := []int{1,1,4,2,1,3}

	// result: 5
	// heights := []int{5,1,2,3,4}

	// result: 0
	heights := []int{1,2,3,4,5}

	// result: 
	// heights := []int{}

	result := heightChecker(heights)
	fmt.Printf("result = %v\n", result)
}

