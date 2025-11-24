package main

import (
	"fmt"
)

func prefixesDivBy5(nums []int) []bool {
	result := make([]bool, len(nums))
	num := int(0)

	for i, n := range nums {
		num = (num * 2 + n) % 5
		result[i] = num == 0
	}

	return result
}

func main() {
	// result: [true,false,false]
	// nums := []int{0,1,1}

	// result: [false,false,false]
	nums := []int{1,1,1}

	// result: []
	// nums := []int{}

	result := prefixesDivBy5(nums)
	fmt.Printf("result = %v\n", result)
}

