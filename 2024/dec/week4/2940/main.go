package main
import (
	"fmt"
)

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	monoStack := make([][2]int, 0)
	result := make([]int, len(queries))
	for i, _ := range result {
		result[i] = -1
	}

	newQueries := make([][][2]int, len(heights))
	for i, _ := range newQueries {
		newQueries[i] = make([][2]int, 0)
	}

	for i := 0; i < len(queries); i++ {
		a := queries[i][0]
		b := queries[i][1]

		if a > b {
			a, b = b, a
		}

		if heights[b] > heights[a] || a == b {
			result[i] = b
		} else {
			newQueries[b] = append(newQueries[b], [2]int{heights[a], i})
		}
	}

	for i := len(heights) - 1; i >= 0; i-- {
		for _, pair := range newQueries[i] {
			position := search(pair[0], monoStack)
			if position < len(monoStack) && position >= 0 {
				result[pair[1]] = monoStack[position][1]
			}
		}

		for len(monoStack) > 0 && monoStack[len(monoStack) - 1][0] <= heights[i] {
			monoStack = monoStack[:len(monoStack) - 1]
		}

		monoStack = append(monoStack, [2]int{heights[i], i})
	}

	return result
}

func search(height int, monoStack [][2]int) int {
	left := int(0)
	right := len(monoStack) - 1
	result := int(-1)

	for left <= right {
		mid := (left + right) / 2
		if monoStack[mid][0] > height {
			result = max(result, mid)
			left = mid + 1
		} else {
			right = mid - 1
		}
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
	// result: [2,5,-1,5,2]
	// heights := []int{6,4,8,5,2,7}
	// queries := [][]int{{0,1},{0,3},{2,4},{3,4},{2,2}}

	// result: [7,6,-1,4,6]
	heights := []int{5,3,8,2,6,1,4,6}
	queries := [][]int{{0,7},{3,5},{5,2},{3,0},{1,6}}

	// result: []
	// heights := []int{}
	// queries := [][]int{}

	result := leftmostBuildingQueries(heights, queries)
	fmt.Printf("result = %v\n", result)
}

