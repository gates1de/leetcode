package main

import (
	"fmt"
)

func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
	first, second, third := int(0), int(0), int(0)

	for offset := range stoneValue {
		i := n - 1 - offset
		current := stoneValue[i] - first
		if i+1 < n {
			candidate := stoneValue[i] + stoneValue[i+1] - second
			if candidate > current {
				current = candidate
			}
		}
		if i+2 < n {
			candidate := stoneValue[i] + stoneValue[i+1] + stoneValue[i+2] - third
			if candidate > current {
				current = candidate
			}
		}

		third = second
		second = first
		first = current
	}

	if first > 0 {
		return "Alice"
	}
	if first < 0 {
		return "Bob"
	}

	return "Tie"
}

func main() {
	// result: "Bob"
	stoneValue := []int{1, 2, 3, 7}

	// result: "Alice"
	// stoneValue := []int{1,2,3,-9}

	// result: "Tie"
	// stoneValue := []int{1,2,3,6}

	// result:
	// stoneValue := []int{}

	result := stoneGameIII(stoneValue)
	fmt.Printf("result = %v\n", result)
}
