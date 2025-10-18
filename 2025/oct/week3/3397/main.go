package main

import (
	"fmt"
	"math"
	"sort"
)

func maxDistinctElements(nums []int, k int) int {
	sort.Ints(nums)
	result := int(0)
	pre := math.MinInt32

	for _, num := range nums {
		tmp := min(max(num - k, pre + 1), num + k)
		if tmp > pre {
			result++
			pre = tmp
		}
	}

	return result
}

func main() {
	// result: 6
	// nums := []int{1,2,2,3,3,4}
	// k := int(2)

	// result: 3
	nums := []int{4,4,4,4}
	k := int(1)

	// result: 
	// nums := []int{}
	// k := int(0)

	result := maxDistinctElements(nums, k)
	fmt.Printf("result = %v\n", result)
}

