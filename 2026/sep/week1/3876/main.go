package main

import (
	"fmt"
)

func uniformArray(nums1 []int) bool {
	minOdd, minEven := 0, 0
	for _, num := range nums1 {
		if num % 2 == 0 {
			if minEven == 0 || num < minEven {
				minEven = num
			}
		} else if minOdd == 0 || num < minOdd {
			minOdd = num
		}
	}

	if minOdd == 0 || minEven == 0 {
		return true
	}

	return minOdd < minEven
}

func main() {
	// result: true
	// nums1 := []int{1,4,7}

	// result: false
	// nums1 := []int{2,3}

	// result: true
	nums1 := []int{4, 6}

	// result:
	// nums1 := []int{}

	result := uniformArray(nums1)
	fmt.Printf("result = %v\n", result)
}
