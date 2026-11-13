package main

import (
	"fmt"
)

func countNonDecreasingSubarrays(nums []int, k int) int64 {
	n := len(nums)
	if n == 0 {
		return 0
	}

	prefix := make([]int64, n+1)
	for i, value := range nums {
		prefix[i+1] = prefix[i] + int64(value)
	}

	nextGreater := make([]int, n)
	stack := make([]int, 0, n)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] <= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			nextGreater[i] = n
		} else {
			nextGreater[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	levels := 1
	for (1 << levels) <= n {
		levels++
	}
	next := make([][]int, levels)
	cost := make([][]int64, levels)
	next[0] = nextGreater
	cost[0] = make([]int64, n)
	for i, end := range nextGreater {
		cost[0][i] = int64(nums[i])*int64(end-i) - (prefix[end] - prefix[i])
	}

	for level := 1; level < levels; level++ {
		next[level] = make([]int, n)
		cost[level] = make([]int64, n)
		previousNext := next[level-1]
		previousCost := cost[level-1]
		for i := range nums {
			middle := previousNext[i]
			if middle == n {
				next[level][i] = n
				cost[level][i] = previousCost[i]
				continue
			}
			next[level][i] = previousNext[middle]
			cost[level][i] = previousCost[i] + previousCost[middle]
		}
	}

	windowCost := func(left int, right int) int64 {
		result := int64(0)
		current := left
		for level := levels - 1; level >= 0; level-- {
			end := next[level][current]
			if end <= right {
				result += cost[level][current]
				current = end
			}
		}
		return result + int64(nums[current])*int64(right-current+1) - (prefix[right+1] - prefix[current])
	}

	result := int64(0)
	left := 0
	for right := range nums {
		for left <= right && windowCost(left, right) > int64(k) {
			left++
		}
		result += int64(right - left + 1)
	}

	return result
}

func main() {
	// result: 17
	// nums := []int{6,3,1,2,4,4}
	// k := 7

	// result: 12
	nums := []int{6, 3, 1, 3, 6}
	k := 4

	// result: 0
	// nums := []int{}
	// k := 0

	result := countNonDecreasingSubarrays(nums, k)
	fmt.Printf("result = %v\n", result)
}
