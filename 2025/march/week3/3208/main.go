package main
import (
	"fmt"
)

func numberOfAlternatingGroups(colors []int, k int) int {
	n := len(colors)
	result := int(0)
	alternatingElementsCount := int(1)
	lastColor := colors[0]

	for i := 1; i < n + k - 1; i++ {
		index := i % n

		if colors[index] == lastColor {
			alternatingElementsCount = 1
			lastColor = colors[index]
			continue
		}

		alternatingElementsCount += 1

		if alternatingElementsCount >= k {
			result++
		}

		lastColor = colors[index]
	}

	return result
}

func main() {
	// result: 3
	// colors := []int{0,1,0,1,0}
	// k := int(3)

	// result: 2
	// colors := []int{0,1,0,0,1,0,1}
	// k := int(6)

	// result: 0
	colors := []int{1,1,0,1}
	k := int(4)

	// result: 
	// colors := []int{}
	// k := int(0)

	result := numberOfAlternatingGroups(colors, k)
	fmt.Printf("result = %v\n", result)
}

