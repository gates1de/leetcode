package main

import (
	"fmt"
)

func stoneGameIX(stones []int) bool {
	counts := [3]int{}
	for _, stone := range stones {
		counts[stone % 3]++
	}

	if counts[0]%2 == 0 {
		return counts[1] > 0 && counts[2] > 0
	}

	difference := counts[1] - counts[2]
	if difference < 0 {
		difference = -difference
	}

	return difference > 2
}

func main() {
	// result: true
	// stones := []int{2,1}

	// result: false
	// stones := []int{2}

	// result: false
	stones := []int{5, 1, 2, 4, 3}

	// result:
	// stones := []int{}

	result := stoneGameIX(stones)
	fmt.Printf("result = %v\n", result)
}
