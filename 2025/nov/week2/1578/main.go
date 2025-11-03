package main

import (
	"fmt"
	"sort"
)

func minCost(colors string, neededTime []int) int {
	n := len(colors)
	if n <= 1 {
		return 0
	}

	result := int(0)
	times := []int{neededTime[0]}
	pre := colors[0]
	for i := 1; i < n; i++ {
		color := colors[i]
		if pre == color {
			times = append(times, neededTime[i])
			continue
		}

		if len(times) > 0 {
			sort.Ints(times)
			for _, time := range times[:len(times) - 1] {
				result += time
			}
		}

		pre = color
		times = []int{neededTime[i]}
	}

	if len(times) > 0 {
		sort.Ints(times)
		for _, time := range times[:len(times) - 1] {
			result += time
		}
	}

	return result
}

func main() {
	// result: 3
	// colors := "abaac"
	// neededTime := []int{1,2,3,4,5}

	// result: 0
	// colors := "abc"
	// neededTime := []int{1,2,3}

	// result: 2
	colors := "aabaa"
	neededTime := []int{1, 2, 3, 4, 1}

	// result:
	// colors := ""
	// neededTime := []int{}

	result := minCost(colors, neededTime)
	fmt.Printf("result = %v\n", result)
}
