package main
import (
	"fmt"
	"sort"
)

const modulo = 1e9 + 7

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	if h <= 0 || w <= 0 {
		return 0
	}

	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)

	if horizontalCuts[len(horizontalCuts) - 1] != h {
		horizontalCuts = append(horizontalCuts, h)
	}
	if verticalCuts[len(verticalCuts) - 1] != w {
		verticalCuts = append(verticalCuts, w)
	}

	result := int(0)
	pre := int(0)
	maxWidth := int(0)
	for _, val := range horizontalCuts {
		diff := val - pre
		if maxWidth < diff {
			maxWidth = diff
		}
		pre = val
	}

	result = maxWidth
	pre = 0
	maxHeight := int(0)
	for _, val := range verticalCuts {
		diff := val - pre
		if maxHeight < diff {
			maxHeight = diff
		}
		pre = val
	}

	maxWidth %= modulo
	maxHeight %= modulo

	if result < maxWidth * maxHeight {
		return maxWidth * maxHeight % modulo
	}

	return result % modulo
}

func main() {
	// result: 4
	// h := int(5)
	// w := int(4)
	// horizontalCuts := []int{1, 2, 4}
	// verticalCuts := []int{1, 3}

	// result: 6
	h := int(5)
	w := int(4)
	horizontalCuts := []int{3, 1}
	verticalCuts := []int{1}

	// result: 
	// h := int(0)
	// w := int(0)
	// horizontalCuts := []int{}
	// verticalCuts := []int{}

	result := maxArea(h, w, horizontalCuts, verticalCuts)
	fmt.Printf("result = %v\n", result)
}

