package main

import (
	"fmt"
	"sort"
)

func minimumBoxes(apple []int, capacity []int) int {
	sum := int(0)
	for _, a := range apple {
		sum += a
	}

	sort.Sort(sort.Reverse(sort.IntSlice(capacity)))
	result := int(0)
	for sum > 0 {
		sum -= capacity[result]
		result++
	}

	return result
}

func main() {
	// result: 2
	// apple := []int{1,3,2}
	// capacity := []int{4,3,1,5,2}

	// result: 4
	apple := []int{5,5,5}
	capacity := []int{2,4,2,7}

	// result: 
	// apple := []int{}
	// capacity := []int{}

	result := minimumBoxes(apple, capacity)
	fmt.Printf("result = %v\n", result)
}

