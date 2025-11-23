package main

import (
	"fmt"
	"math"
)

func maxSumDivThree(nums []int) int {
	f := []int{0, math.MinInt32, math.MinInt32}

	for _, num := range nums {
		temp := make([]int, 3)
		for i := range 3 {
			temp[(i + num) % 3] = max(f[(i + num) % 3], f[i] + num)
		}

		f = temp
	}

	return f[0]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 18
	// nums := []int{3,6,5,1,8}

	// result: 0
	// nums := []int{4}

	// result: 12
	nums := []int{1,2,3,4,4}

	// result: 
	// nums := []int{}

	result := maxSumDivThree(nums)
	fmt.Printf("result = %v\n", result)
}

