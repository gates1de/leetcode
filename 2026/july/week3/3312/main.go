package main

import (
	"fmt"
	"sort"
)

func gcdValues(nums []int, queries []int64) []int {
	maxValue := int(0)
	for _, num := range nums {
		if num > maxValue {
			maxValue = num
		}
	}

	freq := make([]int64, maxValue+1)
	for _, num := range nums {
		freq[num]++
	}

	exact := make([]int64, maxValue+1)
	for g := maxValue; g >= 1; g-- {
		count := int64(0)
		for multiple := g; multiple <= maxValue; multiple += g {
			count += freq[multiple]
		}

		pairs := count * (count - 1) / 2
		for multiple := g * 2; multiple <= maxValue; multiple += g {
			pairs -= exact[multiple]
		}
		exact[g] = pairs
	}

	prefix := make([]int64, maxValue+1)
	running := int64(0)
	for g := 1; g <= maxValue; g++ {
		running += exact[g]
		prefix[g] = running
	}

	result := make([]int, len(queries))
	for i, query := range queries {
		idx := sort.Search(len(prefix), func(j int) bool {
			return prefix[j] > query
		})
		result[i] = idx
	}

	return result
}

func main() {
	// result: [1,2,2]
	// nums := []int{2,3,4}
	// queries := []int64{0,2,2}

	// result: [4,2,1,1]
	// nums := []int{4,4,2,1}
	// queries := []int64{5,3,1,0}

	// result: [2,2]
	nums := []int{2,2}
	queries := []int64{0,0}

	// result: 
	// nums := []int{}
	// queries := []int64{}

	result := gcdValues(nums, queries)
	fmt.Printf("result = %v\n", result)
}
