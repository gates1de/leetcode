package main

import (
	"fmt"
	"sort"
)

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Slice(happiness, func(a, b int) bool {
		return happiness[a] > happiness[b]
	})

	result := int64(0)
	turns := int(0)

	for i := range k {
		result += int64(max(happiness[i] - turns, 0))
		turns++
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
	// result: 4
	// happiness := []int{1,2,3}
	// k := int(2)

	// result: 1
	// happiness := []int{1,1,1,1}
	// k := int(2)

	// result: 5
	happiness := []int{2,3,4,5}
	k := int(1)

	// result: 
	// happiness := []int{}
	// k := int(0)

	result := maximumHappinessSum(happiness, k)
	fmt.Printf("result = %v\n", result)
}

