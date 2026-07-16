package main

import (
	"fmt"
	"sort"
)

func gcdSum(nums []int) int64 {
	n := len(nums)
	prefixGcd := make([]int, n)

	currentMax := 0
	for i := range n {
		if nums[i] > currentMax {
			currentMax = nums[i]
		}
		prefixGcd[i] = gcd(nums[i], currentMax)
	}

	sort.Ints(prefixGcd)

	result := int64(0)
	for i := range n / 2 {
		result += int64(gcd(prefixGcd[i], prefixGcd[n - 1 - i]))
	}

	return result
}

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	// result: 2
	// nums := []int{2, 6, 4}

	// result: 5
	nums := []int{3, 6, 2, 8}

	// result: 
	// nums := []int{}

	result := gcdSum(nums)
	fmt.Printf("result = %v\n", result)
}
