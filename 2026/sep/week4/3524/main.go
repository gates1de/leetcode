package main

import (
	"fmt"
)

func resultArray(nums []int, k int) []int64 {
	if k <= 0 {
		return []int64{}
	}

	result := make([]int64, k)
	ending := make([]int64, k)
	next := make([]int64, k)

	for _, num := range nums {
		for i := range next {
			next[i] = 0
		}

		value := num % k
		next[value]++
		for previous, count := range ending {
			next[previous*value%k] += count
		}

		for remainder, count := range next {
			result[remainder] += count
		}
		ending, next = next, ending
	}

	return result
}

func main() {
	// result: [9,2,4]
	// nums := []int{1,2,3,4,5}
	// k := int(3)

	// result: [18,1,2,0]
	// nums := []int{1,2,4,8,16,32}
	// k := int(4)

	// result: [9,6]
	nums := []int{1, 1, 2, 1, 1}
	k := int(2)

	// result: []
	// nums := []int{}
	// k := int(0)

	result := resultArray(nums, k)
	fmt.Printf("result = %v\n", result)
}
