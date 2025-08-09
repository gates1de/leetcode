package main

import (
	"fmt"
)

func maxTotalFruits(fruits [][]int, startPos int, k int) int {
	n := len(fruits)
	left := int(0)
	right := int(0)
	sum := int(0)
	result := int(0)

	for right < n {
		sum += fruits[right][1]
		for left <= right && step(left, right, fruits, startPos) > k {
			sum -= fruits[left][1]
			left++
		}

		result = max(result, sum)
		right++
	}

	return result
}

func step(left int, right int, fruits [][]int, startPos int) int {
	if fruits[right][0] <= startPos {
		return startPos - fruits[left][0]
	} else if fruits[left][0] >= startPos {
		return fruits[right][0] - startPos
	} else {
		return min(abs(startPos-fruits[right][0]), abs(startPos-fruits[left][0])) + fruits[right][0] - fruits[left][0]
	}
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	// result: 9
	// fruits := [][]int{{2,8},{6,3},{8,6}}
	// startPos := int(5)
	// k := int(4)

	// result: 14
	// fruits := [][]int{{0,9},{4,1},{5,7},{6,2},{7,4},{10,9}}
	// startPos := int(5)
	// k := int(4)

	// result: 0
	fruits := [][]int{{0,3},{6,4},{8,5}}
	startPos := int(3)
	k := int(2)

	// result: 
	// fruits := [][]int{}
	// startPos := int(0)
	// k := int(0)

	result := maxTotalFruits(fruits, startPos, k)
	fmt.Printf("result = %v\n", result)
}

