package main
import (
	"fmt"
	"math"
)

func minDominoRotations(tops []int, bottoms []int) int {
	result := math.MaxInt32
	face := make([]int, 7)

	for i, top := range tops {
		face[top]++
		if top != bottoms[i] {
			face[bottoms[i]]++
		}
	}

	for num := 1; num <= 6; num++ {
		maintainTop := int(0)
		maintainBottom := int(0)
		isPossible := true

		for i, top := range tops {
			if top != num && bottoms[i] != num {
				isPossible = false
				break
			}

			if top != num {
				maintainTop++
			}
			if bottoms[i] != num {
				maintainBottom++
			}
		}

		if isPossible {
			result = min(result, min(maintainTop, maintainBottom))
		}
	}

	if result == math.MaxInt32 {
		return -1
	}

	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 2
	// tops := []int{2, 1, 2, 4, 2, 2}
	// bottoms := []int{5, 2, 6, 2, 3, 2}

	// result: -1
	// tops := []int{3, 5, 1, 2, 3}
	// bottoms := []int{3, 6, 3, 3, 4}

	// result: 1
	// tops := []int{2, 4}
	// bottoms := []int{4, 2}

	// result: 0
	// tops := []int{1, 1, 1, 1, 1}
	// bottoms := []int{2, 3, 4, 5, 6}

	// result: 1
	tops := []int{1,2,1,1,1,2,2,2}
	bottoms := []int{2,1,2,2,2,2,2,2}

	// result: 
	// tops := []int{}
	// bottoms := []int{}

	result := minDominoRotations(tops, bottoms)
	fmt.Printf("result = %v\n", result)
}

