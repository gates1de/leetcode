package main
import (
	"fmt"
	"sort"
)

func maximumBeauty(items [][]int, queries []int) []int {
	n := len(queries)
	result := make([]int, n)
	sort.Slice(items, func(a, b int) bool { return items[a][0] < items[b][0] })

	queriesWithIndices := make([][]int, n);
	for i, _ := range queriesWithIndices {
		queriesWithIndices[i] = []int{queries[i], i}
	}

	sort.Slice(queriesWithIndices, func(a, b int) bool {
		return queriesWithIndices[a][0] < queriesWithIndices[b][0]
	})

	itemIndex := int(0)
	maxBeauty := int(0)

	for i := 0; i < n; i++ {
		query := queriesWithIndices[i][0]
		originalIndex := queriesWithIndices[i][1]

		for itemIndex < len(items) && items[itemIndex][0] <= query {
			maxBeauty = max(maxBeauty, items[itemIndex][1])
			itemIndex++
		}

		result[originalIndex] = maxBeauty
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
	// result: [2,4,5,5,6,6]
	// items := [][]int{{1,2},{3,2},{2,4},{5,6},{3,5}}
	// queries := []int{1,2,3,4,5,6}

	// result: [4]
	// items := [][]int{{1,2},{1,2},{1,3},{1,4}}
	// queries := []int{1}

	// result: [0]
	items := [][]int{{10,1000}}
	queries := []int{5}

	// result: []
	// items := [][]int{}
	// queries := []int{}

	result := maximumBeauty(items, queries)
	fmt.Printf("result = %v\n", result)
}

