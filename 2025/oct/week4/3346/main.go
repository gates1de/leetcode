package main

import (
	"fmt"
)

func maxFrequency(nums []int, k int, numOperations int) int {
	maxNum := nums[0]
	for _, num := range nums {
		maxNum = max(maxNum, num)
	}

	n := maxNum + k + 2
	counter := make([]int, n)
	for _, num := range nums {
		counter[num]++
	}

	pre := make([]int, n)
	pre[0] = counter[0]
	for i := 1; i < n; i++ {
		pre[i] = pre[i - 1] + counter[i]
	}

	result := int(0)
	for i := 0; i < n; i++ {
		if counter[i] == 0 && numOperations == 0 {
			continue
		}

		left := i - k
		if left < 0 {
			left = 0
		}

		right := i + k
		if right > n-1 {
			right = n - 1
		}

		tot := pre[right]
		if left > 0 {
			tot -= pre[left - 1]
		}

		adj := tot - counter[i]
		val := counter[i] + min(numOperations, adj)
		result = max(result, val)
	}

	return result
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

