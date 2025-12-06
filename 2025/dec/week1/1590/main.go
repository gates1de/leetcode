package main

import (
	"fmt"
)

func minSubarray(nums []int, p int) int {
	n := len(nums)
	totalSum := int(0)

	for _, num := range nums {
		totalSum = (totalSum + num) % p
	}

	target := totalSum % p
	if target == 0 {
		return 0
	}

	m := make(map[int]int)
	m[0] = -1
	currentSum := int(0)
	minLen := n

	for i := 0; i < n; i++ {
		currentSum = (currentSum + nums[i]) % p

		needed := (currentSum - target + p) % p
		if _, exists := m[needed]; exists {
			minLen = min(minLen, i - m[needed])
		}

		m[currentSum] = i
	}

	if minLen == n {
		return -1
	}

	return minLen
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 1
	// nums := []int{3,1,4,2}
	// p := int(6)

	// result: 2
	// nums := []int{6,3,5,2}
	// p := int(9)

	// result: 0
	nums := []int{1,2,3}
	p := int(3)

	// result: 
	// nums := []int{}
	// p := int()

	result := minSubarray(nums, p)
	fmt.Printf("result = %v\n", result)
}

