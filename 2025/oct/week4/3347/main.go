package main

import (
	"fmt"
	"sort"
)

func maxFrequency(nums []int, k int, numOperations int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}

	sort.Ints(nums)
	right := min(numOperations, getMaxNums(nums, k))
	index := int(0)
	left := int(0)
	freq := int(0)

	for i, num := range nums {
		if i > 0 && num == nums[i - 1] {
			freq = freq + 1
		} else {
			freq = 1
		}

		minNum := num - k
		maxNum := num + k

		for true {
			if index < n && nums[index] < minNum {
				index++
			} else if left < n && nums[left] <= maxNum {
				left++
			} else {
				break
			}
		}

		right = max(right, min(freq + numOperations, left - index))
	}

	return right
}

func getMaxNums(nums []int, k int) int {
	left := int(0)
	right := int(0)

	for i, num := range nums {
		target := num + 2 * k

		for left < len(nums) && nums[left] <= target {
			left++
		}

		right = max(right, left - i)
	}

	return right
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 2
	// nums := []int{1,4,5}
	// k := int(1)
	// numOperations := int(2)

	// result: 2
	nums := []int{5,11,20,20}
	k := int(5)
	numOperations := int(1)

	// result: 
	// nums := []int{}
	// k := int(0)
	// numOperations := int(0)

	result := maxFrequency(nums, k, numOperations)
	fmt.Printf("result = %v\n", result)
}

