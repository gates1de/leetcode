package main
import (
	"fmt"
	"math"
	"sort"
)

func minStoneSum(piles []int, k int) int {
	counts := make([]int, 100000)
	result := int(0)
	for _, pile := range piles {
		counts[pile]++
		result += pile
	}

	for val := 10000; val > 0 && k > 0; val-- {
		for counts[val] > 0 && k > 0 {
			counts[val]--
			k--
			nextVal := (val + 1) / 2
			counts[nextVal]++
			result -= val / 2
		}
	}
	return result
}

// Time Limit Exceeded
func ngSolution(piles []int, k int) int {
	sort.Ints(piles)

	result := int(0)
	for _, pile := range piles {
		result += pile
	}

	for i := 0; i < k; i++ {
		maxPile := piles[len(piles) - 1]
		newPile := int(math.Ceil(float64(maxPile) / 2))
		result -= maxPile - newPile
		piles[len(piles) - 1] = newPile
		sort.Ints(piles)
	}

	return result
}

func main() {
	// result: 12
	// piles := []int{5,4,9}
	// k := int(2)

	// result: 12
	piles := []int{4,3,6,7}
	k := int(3)

	// result: 
	// piles := []int{}
	// k := int(0)

	result := minStoneSum(piles, k)
	fmt.Printf("result = %v\n", result)
}

