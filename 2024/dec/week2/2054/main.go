package main
import (
	"fmt"
	"sort"
)

func maxTwoEvents(events [][]int) int {
	times := make([][3]int, 0, len(events))

	for _, event := range events {
		times = append(times, [3]int{event[0], 1, event[2]})
		times = append(times, [3]int{event[1] + 1, 0, event[2]})
	}

	sort.Slice(times, func(a, b int) bool {
		if times[a][0] == times[b][0] {
			return times[a][1] < times[b][1]
		}
		return times[a][0] < times[b][0]
	})

	result := int(0)
	maxVal := int(0)

	for _, time := range times {
		if time[1] == 1 {
			result = max(result, time[2] + maxVal)
		} else {
			maxVal = max(maxVal, time[2])
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
	// result: 4
	// events := [][]int{{1,3,2},{4,5,2},{2,4,3}}

	// result: 5
	// events := [][]int{{1,3,2},{4,5,2},{1,5,5}}

	// result: 8
	events := [][]int{{1,5,3},{1,5,1},{6,6,5}}

	// result: 
	// events := [][]int{}

	result := maxTwoEvents(events)
	fmt.Printf("result = %v\n", result)
}

