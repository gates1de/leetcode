package main
import (
	"fmt"
	"math"
	"sort"
)

func longestSquareStreak(nums []int) int {
	streakLengths := make(map[int]int)
	sort.Ints(nums)

	for _, num := range nums {
		root := int(math.Sqrt(float64(num)))

		if root * root == num && streakLengths[root] > 0 {
			streakLengths[num] = streakLengths[root] + 1
		} else {
			streakLengths[num] = 1
		}
	}

	result := int(0)
	for _, streak := range streakLengths {
		result = max(result, streak)
	}

	if result == 1 {
		return -1
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
	// result: 3
	// nums := []int{4,3,6,16,8,2}

	// result: -1
	nums := []int{2,3,5,6,7}

	// result: 
	// nums := []int{}

	result := longestSquareStreak(nums)
	fmt.Printf("result = %v\n", result)
}

