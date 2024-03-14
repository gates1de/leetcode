package main
import (
	"fmt"
)

func numSubarraysWithSum(nums []int, goal int) int {
	start := int(0)
	prefixZeros := int(0)
	currentSum := int(0)
	totalCount := int(0)

	for end := 0; end < len(nums); end++ {
		currentSum += nums[end]

		for start < end && (nums[start] == 0 || currentSum > goal) {
			if nums[start] == 1 {
				prefixZeros = 0
			} else {
				prefixZeros++
			}

			currentSum -= nums[start]
			start++
		}

		if currentSum == goal {
			totalCount += 1 + prefixZeros
		}
	}

	return totalCount
}

func main() {
	// result: 4
	// nums := []int{1,0,1,0,1}
	// goal := int(2)

	// result: 15
	nums := []int{0,0,0,0,0}
	goal := int(0)

	// result: 
	// nums := []int{}
	// goal := int(0)

	result := numSubarraysWithSum(nums, goal)
	fmt.Printf("result = %v\n", result)
}

