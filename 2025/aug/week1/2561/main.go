package main

import (
	"fmt"
	"math"
	"sort"
)

func minCost(basket1 []int, basket2 []int) int64 {
	freq := make(map[int]int)
	m := math.MaxInt32

	for _, b := range basket1 {
		freq[b]++

		if b < m {
			m = b
		}
	}

	for _, b := range basket2 {
		freq[b]--

		if b < m {
			m = b
		}
	}

	merge := make([]int, 0)
	for k, c := range freq {
		if c % 2 != 0 {
			return -1
		}

		for i := 0; i < abs(c) / 2; i++ {
			merge = append(merge, k)
		}
	}

	sort.Ints(merge)
	result := int64(0)
	for i := 0; i < len(merge) / 2; i++ {
		if 2 * m < merge[i] {
			result += int64(2 * m)
		} else {
			result += int64(merge[i])
		}
	}

	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	// result: 1
	// basket1 := []int{4,2,2,2}
	// basket2 := []int{1,4,1,2}

	// result: -1
	basket1 := []int{2,3,4,1}
	basket2 := []int{3,2,5,1}

	// result: 
	// basket1 := []int{}
	// basket2 := []int{}

	result := minCost(basket1, basket2)
	fmt.Printf("result = %v\n", result)
}

