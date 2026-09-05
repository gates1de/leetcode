package main

import (
	"fmt"
)

func uniformArray(nums1 []int) bool {
	if len(nums1) == 0 {
		return false
	}

	odd, even := 0, 0
	for _, num := range nums1 {
		if num % 2 == 0 {
			even++
		} else {
			odd++
		}
	}

	return odd > 0 || even > 0
}

func main() {
	// result: true
	// nums := []int{2,3}

	// result: true
	nums := []int{4, 6}

	// result:
	// nums := []int{}

	result := uniformArray(nums)
	fmt.Printf("result = %v\n", result)
}
