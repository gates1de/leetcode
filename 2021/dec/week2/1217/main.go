package main
import (
	"fmt"
)

func minCostToMoveChips(position []int) int {
	odd := int(0)
	even := int(0)
	for _, num := range position {
		if num % 2 == 0 {
			even++
		} else {
			odd++
		}
	}

	if odd < even {
		return odd
	}
	return even
}

func main() {
	// result: 1
	// position := []int{1, 2, 3}

	// result: 2
	// position := []int{2, 2, 2, 3, 3}

	// result: 1
	// position := []int{1, 1000000000}

	// result: 8
	position := []int{8,3,9,1,4,6,2,3,9,4,6,8,7,1,9,6,4}

	// result: 
	// position := []int{}

	result := minCostToMoveChips(position)
	fmt.Printf("result = %v\n", result)
}

