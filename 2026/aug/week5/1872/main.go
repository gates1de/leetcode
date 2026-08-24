package main

import (
	"fmt"
)

func stoneGameVIII(stones []int) int {
	prefixSums := make([]int, len(stones))
	prefixSums[0] = stones[0]
	for i := 1; i < len(stones); i++ {
		prefixSums[i] = prefixSums[i - 1] + stones[i]
	}

	best := prefixSums[len(stones) - 1]
	for i := len(stones) - 2; i >= 1; i-- {
		best = max(best, prefixSums[i] - best)
	}

	return best
}

func main() {
	// result: 5
	// stones := []int{-1,2,-3,4,-5}

	// result: 13
	// stones := []int{7,-6,5,10,5,-2,-6}

	// result: -22
	stones := []int{-10, -12}

	// result:
	// stones := []int{}

	result := stoneGameVIII(stones)
	fmt.Printf("result = %v\n", result)
}
