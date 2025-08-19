package main

import (
	"fmt"
)

func zeroFilledSubarray(nums []int) int64 {
	result := int64(0)
	streak := int64(0)

	for _, num := range nums {
		if num == 0 {
			streak++
			result += streak
		} else {
			streak = 0
		}
	}

	return result
}

func main() {
	// result: 6
	// nums := []int{1,3,0,0,2,0,0,4}

	// result: 9
	// nums := []int{0,0,0,2,0,0}

	// result: 0
	nums := []int{2,10,2019}

	// result: 
	// nums := []int{}

	result := zeroFilledSubarray(nums)
	fmt.Printf("result = %v\n", result)
}

