package main

import (
	"fmt"
	"math"
)

const mark = math.MinInt / 2

func maxSumTrionic(nums []int) int64 {
	pBeg := nums[0]
	pInc1 := mark
	pDec := mark
	pInc2 := mark
	result := math.MinInt

	for i := 1; i < len(nums); i++ {
		n := nums[i]
		if pBeg < n {
			pBeg, pInc1, pDec, pInc2 = n, max(pInc1 + n, pBeg + n), mark, max(pInc2 + n, pDec + n)
			result = max(result, pInc2)
		} else if n < pBeg {
			pBeg, pInc1, pDec, pInc2 = n, mark, max(pDec + n, pInc1 + n), mark
		} else {
			pBeg, pInc1, pDec, pInc2 = n, mark, mark, mark
		}
	}

	return int64(result)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: -4
	// nums := []int{0,-2,-1,-3,0,2,-1}

	// result: 14
	nums := []int{1,4,2,7}

	// result: 
	// nums := []int{}

	result := maxSumTrionic(nums)
	fmt.Printf("result = %v\n", result)
}

