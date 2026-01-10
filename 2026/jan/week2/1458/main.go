package main

import (
	"fmt"
	"math"
)

func maxDotProduct(nums1 []int, nums2 []int) int {
	firstMax := math.MinInt32
	secondMax := math.MinInt32
	firstMin := math.MaxInt32
	secondMin := math.MaxInt32

	for _, num := range nums1 {
		firstMax = max(firstMax, num)
		firstMin = min(firstMin, num)
	}

	for _, num := range nums2 {
		secondMax = max(secondMax, num)
		secondMin = min(secondMin, num)
	}

	if firstMax < 0 && secondMin > 0 {
		return firstMax * secondMin
	}
	if firstMin > 0 && secondMax < 0 {
		return firstMin * secondMax
	}


	m := len(nums2) + 1
	dp := make([]int, m)
	prevDp := make([]int, m)

	for i := len(nums1) - 1; i >= 0 ; i-- {
		dp = make([]int, m)
		for j := len(nums2) - 1; j >= 0; j-- {
			use := nums1[i] * nums2[j] + prevDp[j + 1]
			dp[j] = max(use, max(prevDp[j], dp[j + 1]))
		}

		prevDp = dp
	}

	return dp[0]
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
	// result: 18
	// nums1 := []int{2,1,-2,5}
	// nums2 := []int{3,0,-6}

	// result: 21
	// nums1 := []int{3,-2}
	// nums2 := []int{2,-6,7}

	// result: -1
	nums1 := []int{-1,-1}
	nums2 := []int{1,1}

	// result: 
	// nums1 := []int{}
	// nums2 := []int{}

	result := maxDotProduct(nums1, nums2)
	fmt.Printf("result = %v\n", result)
}

