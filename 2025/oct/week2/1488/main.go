package main

import (
	"fmt"
	"sort"
)

func avoidFlood(rains []int) []int {
	n := len(rains)
	stack := make([]int, 0, n)
	m := make(map[int]int)
	result := make([]int, n)

	for i := range n {
		result[i] = 1
	}

	for i, rain := range rains {
		if rain == 0 {
			stack = append(stack, i)
		} else {
			result[i] = -1

			if day, exists := m[rain]; exists {
				it := sort.SearchInts(stack, day)
				if it == len(stack) {
					return []int{}
				}

				result[stack[it]] = rain
				copy(stack[it:len(stack) - 1], stack[it + 1:])
				stack = stack[:len(stack) - 1]
			}

			m[rain] = i
		}
	}

	return result
}

func main() {
	// result: [-1,-1,-1,-1]
	// rains := []int{1,2,3,4}

	// result: [-1,-1,2,1,-1,-1]
	// rains := []int{1,2,0,0,2,1}

	// result: []
	rains := []int{1,2,0,1,2}

	// result: []
	// rains := []int{}

	result := avoidFlood(rains)
	fmt.Printf("result = %v\n", result)
}

