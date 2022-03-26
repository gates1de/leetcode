package main
import (
	"fmt"
)

const MAX = int(20001)

func minDominoRotations(tops []int, bottoms []int) int {
	n := len(tops)
	result := MAX

	for num := 1; num <= 6; num++ {
		topCount := int(0)
		bottomCount := int(0)
		minCount := int(0)

		isShortage := false
		for i := 0; i < n; i++ {
			top := tops[i]
			bottom := bottoms[i]

			if top != num && bottom != num {
				isShortage = true
				break
			}

			if top != num && bottom == num {
				topCount++
			}
			if top == num && bottom != num {
				bottomCount++
			}
		}

		if isShortage {
			continue
		}

		if topCount <= bottomCount {
			minCount = topCount
		} else {
			minCount = bottomCount
		}

		swap := MAX
		if minCount <= n - minCount {
			swap = minCount
		} else {
			swap = n - minCount
		}

		if swap < result {
			result = swap
		}
	}

	if result == MAX {
		return -1
	}

	return result
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

