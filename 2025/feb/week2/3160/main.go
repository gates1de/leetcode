package main
import (
	"fmt"
)

func queryResults(limit int, queries [][]int) []int {
	n := len(queries)
	result := make([]int, n)
	colorMap := make(map[int]int)
	ballMap := make(map[int]int)

	for i := 0; i < n; i++ {
		ball := queries[i][0]
		color := queries[i][1]

		if _, exists := ballMap[ball]; exists {
			prevColor := ballMap[ball]
			colorMap[prevColor]--

			if colorMap[prevColor] == 0 {
				delete(colorMap, prevColor)
			}
		}

		ballMap[ball] = color
		colorMap[color]++
		result[i] = len(colorMap)
	}

	return result
}

func main() {
	// result: [1,2,2,3]
	// limit := int(4)
	// queries := [][]int{{1,4},{2,5},{1,3},{3,4}}

	// result: [1,2,2,3,4]
	limit := int(4)
	queries := [][]int{{0,1},{1,2},{2,2},{3,4},{4,5}}

	// result: []
	// limit := int(0)
	// queries := [][]int{}

	result := queryResults(limit, queries)
	fmt.Printf("result = %v\n", result)
}

