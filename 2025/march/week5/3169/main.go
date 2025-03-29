package main
import (
	"fmt"
	"sort"
)

func countDays(days int, meetings [][]int) int {
	result := int(0)
	latestEnd := int(0)

	sort.Slice(meetings, func(a, b int) bool {
		return meetings[a][0] < meetings[b][0]
	})

	for _, meeting := range meetings {
		start := meeting[0]
		end := meeting[1]

		if start > latestEnd + 1 {
			result += start - latestEnd - 1
		}

		latestEnd = max(latestEnd, end)
	}

	result += days - latestEnd
	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 2
	// days := int(10)
	// meetings := [][]int{{5,7},{1,3},{9,10}}

	// result: 1
	// days := int(5)
	// meetings := [][]int{{2,4},{1,3}}

	// result: 0
	days := int(6)
	meetings := [][]int{{1,6}}

	// result: 
	// days := int(0)
	// meetings := [][]int{}

	result := countDays(days, meetings)
	fmt.Printf("result = %v\n", result)
}

