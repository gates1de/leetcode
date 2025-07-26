package main

import (
	"fmt"
	"math"
)

func maxSubarrays(n int, conflictingPairs [][]int) int64 {
	bMin1 := make([]int, n + 1)
	bMin2 := make([]int, n + 1)
	for i := 0; i <= n; i++ {
		bMin1[i] = math.MaxInt32
		bMin2[i] = math.MaxInt32
	}

	for _, pair := range conflictingPairs {
		a := min(pair[0], pair[1])
		b := max(pair[0], pair[1])

		if bMin1[a] > b {
			bMin2[a] = bMin1[a]
			bMin1[a] = b
		} else if bMin2[a] > b {
			bMin2[a] = b
		}
	}

	result := int64(0)
	ib1 := n
	b2 := math.MaxInt32
	deleteCount := make([]int64, n + 1)
	for i := n; i >= 1; i-- {
		if bMin1[ib1] > bMin1[i] {
			b2 = min(b2, bMin1[ib1])
			ib1 = i
		} else {
			b2 = min(b2, bMin1[i])
		}

		result += int64(min(bMin1[ib1], n + 1) - i)
		deleteCount[ib1] += int64(min(min(b2, bMin2[ib1]), n + 1) - min(bMin1[ib1], n + 1))
	}

	maxVal := int64(0)
	for _, val := range deleteCount {
			maxVal = max64(maxVal, val)
	}

	return result + maxVal
}

func max64(a, b int64) int64 {
	if b > a {
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

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 9
	// n := int(4)
	// conflictingPairs := [][]int{{2,3},{1,4}}

	// result: 12
	n := int(5)
	conflictingPairs := [][]int{{1,2},{2,5},{3,5}}

	// result: 
	// n := int(0)
	// conflictingPairs := [][]int{}

	result := maxSubarrays(n, conflictingPairs)
	fmt.Printf("result = %v\n", result)
}

