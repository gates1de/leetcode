package main

import (
	"fmt"
)

func isTrionic(nums []int) bool {
	n := len(nums)
	i := int(1)

	for i < n && nums[i - 1] < nums[i] {
		i++
	}
	p := i - 1

	for i < n && nums[i - 1] > nums[i] {
		i++
	}
	q := i - 1

	for i < n && nums[i - 1] < nums[i] {
		i++
	}
	flag := i - 1

	return p != 0 && q != p && flag == n - 1 && flag != q
}

func main() {
	// result: true
	// nums := []int{1,3,5,4,2,6}

	// result: false
	nums := []int{2,1,3}

	// result: 
	// nums := []int{}

	result := isTrionic(nums)
	fmt.Printf("result = %v\n", result)
}

