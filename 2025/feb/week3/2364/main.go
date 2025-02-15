package main
import (
	"fmt"
)

func countBadPairs(nums []int) int64 {
	result := int64(0)
	diffCount := make(map[int]int)

	for i, num := range nums {
		diff := i - num
		goodPairsCount := diffCount[diff]
		result += int64(i - goodPairsCount)
		diffCount[diff] = goodPairsCount + 1
	}

	return result
}

func main() {
	// result: 5
	// nums := []int{4,1,3,3}

	// result: 0
	nums := []int{1,2,3,4,5}

	// result: 
	// nums := []int{}

	result := countBadPairs(nums)
	fmt.Printf("result = %v\n", result)
}

