package main
import (
	"fmt"
)

func countMaxOrSubsets(nums []int) int {
	n := len(nums)
	maxOrValue := int(0)

	for _, num := range nums {
		maxOrValue |= num
	}

	totalSubsets := int(1 << n)
	result := int(0)

	for subsetMask := 0; subsetMask < totalSubsets; subsetMask++ {
		currentOrValue := int(0)

		for i := 0; i < n; i++ {
			if ((subsetMask >> i) & 1) == 1 {
				currentOrValue |= nums[i]
			}
		}

		if currentOrValue == maxOrValue {
			result++
		}
	}

	return result
}

func main() {
	// result: 2
	// nums := []int{3,1}

	// result: 7
	// nums := []int{2,2,2}

	// result: 6
	nums := []int{3,2,1,5}

	// result: 
	// nums := []int{}

	result := countMaxOrSubsets(nums)
	fmt.Printf("result = %v\n", result)
}

