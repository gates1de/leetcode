package main
import (
	"fmt"
)

func maxAscendingSum(nums []int) int {
	maxSum := int(0)
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i - 1] {
			maxSum = max(maxSum, currentSum)
			currentSum = 0
		}
		currentSum += nums[i]
	}

	return max(maxSum, currentSum);
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 65
	// nums := []int{10,20,30,5,10,50}

	// result: 150
	// nums := []int{10,20,30,40,50}

	// result: 33
	nums := []int{12,17,15,13,10,11,12}

	// result: 
	// nums := []int{}

	result := maxAscendingSum(nums)
	fmt.Printf("result = %v\n", result)
}

