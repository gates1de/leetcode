package main
import (
	"fmt"
	"sort"
)

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Slice(horizontalCuts, func(a, b int) bool { return horizontalCuts[a] < horizontalCuts[b] })
	sort.Slice(verticalCuts, func(a, b int) bool { return verticalCuts[a] < verticalCuts[b] })

	maxHeight := int(0)
	prev := int(0)
	for _, num := range horizontalCuts {
		if maxHeight < num - prev {
			maxHeight = num - prev
		}
		prev = num
	}
	if maxHeight < h - prev {
		maxHeight = h - prev
	}

	prev = 0
	maxWidth := int(0)
	for _, num := range verticalCuts {
		if maxWidth < num - prev {
			maxWidth = num - prev
		}
		prev = num
	}
	if maxWidth < w - prev {
		maxWidth = w - prev
	}

	return maxHeight * maxWidth % (1e9 + 7)
}

func main() {
	// result: 4
	// h := int(5)
	// w := int(4)
	// horizontalCuts := []int{1, 2, 4}
	// verticalCuts := []int{1, 3}

	// result: 6
	// h := int(5)
	// w := int(4)
	// horizontalCuts := []int{3, 1}
	// verticalCuts := []int{1}

	// result: 9
	// h := int(5)
	// w := int(4)
	// horizontalCuts := []int{3}
	// verticalCuts := []int{3}

	// result: 992399247
	// h := int(500000)
	// w := int(400000)
	// horizontalCuts := []int{1, 2, 4, 300000}
	// verticalCuts := []int{1, 3, 20}

	// result: 3
	h := int(2)
	w := int(7)
	horizontalCuts := []int{1}
	verticalCuts := []int{2, 1, 5}

	// result: 
	// h := int(0)
	// w := int(0)
	// horizontalCuts := []int{}
	// verticalCuts := []int{}

	result := maxArea(h, w, horizontalCuts, verticalCuts)
	fmt.Printf("result = %v\n", result)
}

