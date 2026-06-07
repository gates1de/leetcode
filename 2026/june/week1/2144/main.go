package main

import (
	"fmt"
	"sort"
)

func minimumCost(cost []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(cost)))
	result := int(0)
	for i, price := range cost {
		if (i + 1) % 3 != 0 {
			result += price
		}
	}

	return result
}

func main() {
	// result: 5
	// cost := []int{1,2,3}

	// result: 23
	// cost := []int{6,5,7,9,2,2}

	// result: 10
	cost := []int{5,5}

	// result: 
	// cost := []int{}

	result := minimumCost(cost)
	fmt.Printf("result = %v\n", result)
}
