package main

import (
	"fmt"
)

func maxDistance(nums1 []int, nums2 []int) int {
	i, j := int(0), int(0)
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			i++
		}

		j++
	}

	return max(0, j - i - 1)
}

func main() {
	// result: 2
	// nums1 := []int{55,30,5,4,2}
	// nums2 := []int{100,20,10,10,5}

	// result: 1
	// nums1 := []int{2,2,2}
	// nums2 := []int{10,10,1}

	// result: 2
	nums1 := []int{30,29,19,5}
	nums2 := []int{25,25,25,25,25}

	// result: 
	// nums1 := []int{}
	// nums2 := []int{}

	result := maxDistance(nums1, nums2)
	fmt.Printf("result = %v\n", result)
}

