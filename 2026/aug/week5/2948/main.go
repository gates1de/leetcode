package main

import (
	"fmt"
	"sort"
)

func lexicographicallySmallestArray(nums []int, limit int) []int {
	if len(nums) == 0 {
		return nums
	}

	sortedNums := append([]int(nil), nums...)
	sort.Ints(sortedNums)

	valueGroup := make(map[int]int, len(nums))
	groupNext := []int{0}
	group := 0
	for i, num := range sortedNums {
		if i > 0 && num-sortedNums[i-1] > limit {
			group++
			groupNext = append(groupNext, i)
		}
		valueGroup[num] = group
	}

	for i, num := range nums {
		group := valueGroup[num]
		nums[i] = sortedNums[groupNext[group]]
		groupNext[group]++
	}

	return nums
}

func main() {
	// result: [1,3,5,8,9]
	// nums := []int{1,5,3,9,8}
	// limit := int(2)

	// result: [1,6,7,18,1,2]
	// nums := []int{1,7,6,18,2,1}
	// limit := int(3)

	// result: [1,7,28,19,10]
	nums := []int{1, 7, 28, 19, 10}
	limit := int(3)

	// result: []
	// nums := []int{}
	// limit := int(0)

	result := lexicographicallySmallestArray(nums, limit)
	fmt.Printf("result = %v\n", result)
}
