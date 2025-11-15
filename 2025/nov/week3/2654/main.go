package main

import (
	"fmt"
)

func minOperations(nums []int) int {
	n := len(nums)
	num1 := int(0)

	g := int(0)
	for _, x := range nums {
		if x == 1 {
			num1++
		}
		g = gcd(g, x)
	}

	if num1 > 0 {
		return n - num1
	}
	if g > 1 {
		return -1
	}

	minLen := n
	for i := range n {
		currentGcd := 0

		for j := i; j < n; j++ {
			currentGcd = gcd(currentGcd, nums[j])
			if currentGcd == 1 {
				if j - i + 1 < minLen {
					minLen = j - i + 1
				}

				break
			}
		}
	}

	return minLen + n - 2
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a % b
	}
	return a
}

func main() {
	// result: 4
	// nums := []int{2,6,3,4}

	// result: -1
	nums := []int{2,10,6,14}

	// result: 
	// nums := []int{}

	result := minOperations(nums)
	fmt.Printf("result = %v\n", result)
}

