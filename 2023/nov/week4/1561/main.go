package main
import (
	"fmt"
	"sort"
)

func maxCoins(piles []int) int {
	sort.Ints(piles)
	queue := make([]int, 0, len(piles))
	for _, num := range piles {
		queue = append(queue, num)
	}

	result := int(0)
	for len(queue) > 0 {
		queue = queue[:len(queue) - 1]
		result += queue[len(queue) - 1]
		queue = queue[:len(queue) - 1]
		queue = queue[1:]
	}

	return result
}

func main() {
	// result: 9
	// piles := []int{2,4,1,2,7,8}

	// result: 4
	// piles := []int{2,4,5}

	// result: 18
	piles := []int{9,8,7,6,5,1,2,3,4}

	// result: 
	// piles := []int{}

	result := maxCoins(piles)
	fmt.Printf("result = %v\n", result)
}

