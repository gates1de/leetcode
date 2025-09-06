package main

import (
	"fmt"
)

func minOperations(queries [][]int) int64 {
	result := int64(0)

	for _, q := range queries {
		count1 := get(q[1])
		count2 := get(q[0] - 1)
		result += (count1 - count2 + 1) / 2
	}

	return result
}

func get(num int) int64 {
	result := int64(0)
	i := int(1)
	base := int(1)

	for base <= num {
		end := base * 2 - 1
		if end > num {
			end = num
		}

		result += int64((i + 1) / 2) * int64(end - base + 1)
		i++
		base *= 2
	}

	return result
}

func main() {
	// result: 3
	// queries := [][]int{{1,2},{2,4}}

	// result: 4
	queries := [][]int{{2,6}}

	// result: 
	// queries := [][]int{}

	result := minOperations(queries)
	fmt.Printf("result = %v\n", result)
}

