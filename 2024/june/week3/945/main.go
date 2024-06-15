package main
import (
	"fmt"
)

func minIncrementForUnique(nums []int) int {
	n := len(nums)
	maxNum := int(0)
	result := int(0)

	for _, num := range nums {
		maxNum = max(maxNum, num)
	}

	freq := make([]int, n + maxNum)
	for _, num := range nums {
		freq[num]++
	}

	for i, count := range freq {
		if count <= 1 {
			continue
		}

		duplicates := count - 1
		freq[i + 1] += duplicates
		freq[i] = 1
		result += duplicates
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
	// result: 1
	// nums := []int{1,2,2}

	// result: 6
	nums := []int{3,2,1,2,1,7}

	// result: 
	// nums := []int{}

	result := minIncrementForUnique(nums)
	fmt.Printf("result = %v\n", result)
}

