package main

import (
	"fmt"
)

func findSmallestInteger(nums []int, value int) int {
	m := make([]int, value)
	for _, num := range nums {
		val := ((num % value) + value) % value
		m[val]++
	}

	result := int(0)
	for m[result % value] > 0 {
		m[result % value]--
		result++
	}

	return result
}

func main() {
	// result: 4
	// nums := []int{}
	// value := int(5)

	// result: 2
	nums := []int{1,-10,7,13,6,8}
	value := int(5)

	// result: 
	// nums := []int{}
	// value := int(0)

	result := findSmallestInteger(nums, value)
	fmt.Printf("result = %v\n", result)
}

