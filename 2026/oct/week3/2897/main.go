package main

import (
	"fmt"
)

const modulo = int64(1e9 + 7)

func maxSum(nums []int, k int) int {
	bits := [31]int{}
	for _, value := range nums {
		for bit := range bits {
			if value&(1<<bit) != 0 {
				bits[bit]++
			}
		}
	}

	result := int64(0)
	for range k {
		value := int64(0)
		for bit := range bits {
			if bits[bit] > 0 {
				value |= int64(1) << bit
				bits[bit]--
			}
		}
		result = (result + value*value) % modulo
	}

	return int(result)
}

func main() {
	// result: 2
	// nums := []int{2,6,5,8}
	// k := int(2)

	// result: 90
	nums := []int{4, 5, 4, 7}
	k := int(3)

	// result: 0
	// nums := []int{}
	// k := int(0)

	result := maxSum(nums, k)
	fmt.Printf("result = %v\n", result)
}
