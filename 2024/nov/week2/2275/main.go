package main
import (
	"fmt"
)

func largestCombination(candidates []int) int {
	result := int(0)

	for i := 0; i < 24; i++ {
		count := int(0)
		for _, num := range candidates {
			if (num & (1 << i)) != 0 {
				count++
			}
		}

		result = max(result, count)
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
	// candidates := []int{16,17,71,62,12,24,14}

	// result: 2
	candidates := []int{8,8}

	// result: 
	// candidates := []int{}

	result := largestCombination(candidates)
	fmt.Printf("result = %v\n", result)
}

