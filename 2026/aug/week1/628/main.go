package main

import (
	"fmt"
)

const (
	minimumValue = int(-1001)
	maximumValue = int(1001)
)

func maximumProduct(nums []int) int {
	max1, max2, max3 := minimumValue, minimumValue, minimumValue
	min1, min2 := maximumValue, maximumValue

	for _, num := range nums {
		if num >= max1 {
			max3 = max2
			max2 = max1
			max1 = num
		} else if num >= max2 {
			max3 = max2
			max2 = num
		} else if num > max3 {
			max3 = num
		}

		if num <= min1 {
			min2 = min1
			min1 = num
		} else if num < min2 {
			min2 = num
		}
	}

	result := max1 * max2 * max3
	if candidate := max1 * min1 * min2; candidate > result {
		result = candidate
	}

	return result
}

func main() {
	// result: 6
	// nums := []int{1,2,3}

	// result: 24
	// nums := []int{1,2,3,4}

	// result: -6
	nums := []int{-1,-2,-3}

	// result: 
	// nums := []int{}

	result := maximumProduct(nums)
	fmt.Printf("result = %v\n", result)
}
