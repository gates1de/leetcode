package main
import (
	"fmt"
	"sort"
)

func lastStoneWeight(stones []int) int {
	n := len(stones)
	if n == 0 {
		return 0
	} else if n == 1 {
		return stones[0]
	}

	sort.Slice(stones, func(a, b int) bool { return stones[a] > stones[b] })
	stones[0] -= stones[1]
	if stones[0] == 0 {
		return lastStoneWeight(stones[2:])
	}

	newStones := make([]int, n - 1)
	newStones[0] = stones[0]
	copy(newStones[1:], stones[2:])
	return lastStoneWeight(newStones)
}

func main() {
	// result: 1
	// stones := []int{2, 7, 4, 1, 8, 1}

	// result: 1
	// stones := []int{1}

	// result: 0
	// stones := []int{1, 1, 1, 1, 1, 1}

	// result: 0
	stones := []int{5,4,8,9,3,1,7,6,5,8,9,3,4,7,5,9,3,4}

	// result: 
	// stones := []int{}

	result := lastStoneWeight(stones)
	fmt.Printf("result = %v\n", result)
}

