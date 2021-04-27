package main
import (
	"fmt"
)

func furthestBuilding(heights []int, bricks int, ladders int) int {
	diffs := make([]int, len(heights))
	maxDiff := int(0)
	for i, height := range heights {
		if i == len(heights) - 1 {
			break
		}
		diffs[i] = heights[i + 1] - height
		if maxDiff < diffs[i] {
			maxDiff = diffs[i]
		}
	}
	result := int(0)
	for _, diff := range diffs[:len(diffs) - 1] {
		if diff <= 0 {
			result++
			continue
		}
		if diff == maxDiff && ladders > 0 {
			// fmt.Printf("use ladder\n")
			result++
			ladders--
			continue
		}
		if bricks >= diff {
			// fmt.Printf("use %v bricks (remain: %v)\n", diff, bricks - diff)
			result++
			bricks -= diff
			continue
		}
		if ladders > 0 {
			// fmt.Printf("use ladder\n")
			result++
			ladders--
			continue
		}
		break
	}
	return result
}

func main() {
	// result: 4
	// heights := []int{4, 2, 7, 6, 9, 14, 12}
	// bricks := int(5)
	// ladders := int(1)

	// result: 7
	// heights := []int{4, 12, 2, 7, 3, 18, 20, 3, 19}
	// bricks := int(10)
	// ladders := int(2)

	// result: 3
	// heights := []int{14, 3, 19, 3}
	// bricks := int(17)
	// ladders := int(0)

	// result: 5
	heights := []int{1, 5, 1, 2, 3, 4, 10000}
	bricks := int(4)
	ladders := int(1)

	// result: 
	// heights := []int{}
	// bricks := int(0)
	// ladders := int(0)

	result := furthestBuilding(heights, bricks, ladders)
	fmt.Printf("result = %v\n", result)
}

