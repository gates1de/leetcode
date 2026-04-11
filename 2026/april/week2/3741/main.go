package main

import (
	"fmt"
)

func minimumDistance(nums []int) int {
	n := len(nums)
	next := make([]int, n)
	for i := range next {
		next[i] = -1
	}

	occur := make(map[int]int)
	result := n + 1
	for i := n - 1; i >= 0; i-- {
		if val, ok := occur[nums[i]]; ok {
			next[i] = val
		}
		occur[nums[i]] = i
	}

	for i := range n {
		secondPos := next[i]
		if secondPos != -1 {
			thirdPos := next[secondPos]
			if thirdPos != -1 {
				if dist := thirdPos - i; dist < result {
					result = dist
				}
			}
		}
	}

	if result == n + 1 {
		return -1
	}

	return result * 2
}

func main() {
	// result: 6
	// nums := []int{1,2,1,1,3}

	// result: 8
	// nums := []int{1,1,2,3,2,1,2}

	// result: -1
	nums := []int{1}

	// result: 
	// nums := []int{}

	result := minimumDistance(nums)
	fmt.Printf("result = %v\n", result)
}

