package main

import (
	"fmt"
)

func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	n := len(startTime)
	result := int(0)
	sums := make([]int, n + 1)

	for i := 0; i < n; i++ {
		sums[i + 1] = sums[i] + endTime[i] - startTime[i]
	}

	for i := k - 1; i < n; i++ {
		right := int(0)

		if i == n - 1 {
			right = eventTime
		} else {
			right = startTime[i + 1]
		}

		left := int(0)
		if i == k - 1 {
			left = 0
		} else {
			left = endTime[i - k]
		}

		result = max(result, right - left - (sums[i + 1] - sums[i - k + 1]))
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 2
	// eventTime := int(5)
	// k := int(1)
	// startTime := []int{1,3}
	// endTime := []int{2,5}

	// result: 6
	// eventTime := int(10)
	// k := int(1)
	// startTime := []int{0,2,9}
	// endTime := []int{1,4,10}

	// result: 0
	eventTime := int(5)
	k := int(2)
	startTime := []int{0,1,2,3,4}
	endTime := []int{1,2,3,4,5}

	// result: 
	// eventTime := int(0)
	// k := int(0)
	// startTime := []int{}
	// endTime := []int{}

	result := maxFreeTime(eventTime, k, startTime, endTime)
	fmt.Printf("result = %v\n", result)
}

