package main

import (
	"fmt"
	"math"
)

func minimumPairRemoval(nums []int) int {
	n := len(nums)
	result := int(0)

	for n > 1 {
		isNonDecreasing := true
		maxSum := math.MaxInt32
		maxIndex := int(-1)

		for i := range n - 1 {
			if nums[i] > nums[i + 1] {
				isNonDecreasing = false
			}
			if nums[i] + nums[i + 1] < maxSum {
				maxSum = nums[i] + nums[i + 1]
				maxIndex = i
			}
		}

		if isNonDecreasing {
			break
		}

		nums[maxIndex] = maxSum
		for i := maxIndex + 1; i < n - 1; i++ {
			nums[i] = nums[i + 1]
		}

		n--
		result++
	}

	return result
}

func main() {
	// result: 2
	// nums := []int{5,2,3,1}

	// result: 0
	nums := []int{1,2,2}

	// result: 
	// nums := []int{}

	result := minimumPairRemoval(nums)
	fmt.Printf("result = %v\n", result)
}

