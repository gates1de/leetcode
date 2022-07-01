package main
import (
	"fmt"
	"sort"
)

func maximumUnits(boxTypes [][]int, truckSize int) int {
	sort.Slice(boxTypes, func (a, b int) bool {
		return boxTypes[a][1] > boxTypes[b][1]
	})

	result := int(0)
	i := int(0)
	for truckSize > 0 {
		if i >= len(boxTypes) {
			break
		}
		putCount := min(truckSize, boxTypes[i][0])
		result += putCount * boxTypes[i][1]
		truckSize -= putCount
		i++
	}

	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 8
	// boxTypes := [][]int{{1, 3}, {2, 2}, {3, 1}}
	// truckSize := int(4)

	// result: 91
	// boxTypes := [][]int{{5,10},{2,5},{4,7},{3,9}}
	// truckSize := int(10)

	// result: 15
	boxTypes := [][]int{{3, 5}}
	truckSize := int(4)

	// result: 
	// boxTypes := [][]int{}
	// truckSize := int(0)

	result := maximumUnits(boxTypes, truckSize)
	fmt.Printf("result = %v\n", result)
}

