package main

import (
	"fmt"
)

func countValidSelections(nums []int) int {
	sum := int(0)
	for _, num := range nums {
		sum += num
	}

	result := int(0)
	leftSum := int(0)
	rightSum := sum

	for _, num := range nums {
		if num == 0 {
			if leftSum - rightSum >= 0 && leftSum - rightSum <= 1 {
				result++
			}
			if rightSum - leftSum >= 0 && rightSum - leftSum <= 1 {
				result++
			}
		} else {
			leftSum += num
			rightSum -= num
		}
	}

	return result
}

func main() {
	// result: 2
	// nums := []int{1,0,2,0,3}

	// result: 0
	nums := []int{2,3,4,0,4,1,0}

	// result: 
	// nums := []int{}

	result := countValidSelections(nums)
	fmt.Printf("result = %v\n", result)
}

