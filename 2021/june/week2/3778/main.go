package main
import (
	"fmt"
	"sort"
)

func maximumUnits(boxTypes [][]int, truckSize int) int {
	result := int(0)
	sort.Slice(boxTypes, func (a, b int) bool { return boxTypes[a][1] > boxTypes[b][1] })

	boxIndex := int(0)
	i := int(0)
	for i < truckSize && boxIndex < len(boxTypes) {
		box := boxTypes[boxIndex]
		for j := 0; j < box[0]; j++ {
			result += box[1]
			i++
			if i >= truckSize {
				break
			}
		}
		boxIndex++
	}
	return result
}

func main() {
	// result: 8
	// boxTypes := [][]int{{1, 3}, {2, 2}, {3, 1}}
	// truckSize := int(4)

	// result: 91
	// boxTypes := [][]int{{5, 10}, {2, 5}, {4, 7}, {3 ,9}}
	// truckSize := int(10)

	// result: 107
	boxTypes := [][]int{{5, 10}, {3, 11}, {2, 5}, {4, 7}, {1, 15}, {3 ,9}}
	truckSize := int(10)

	// result: 
	// boxTypes := [][]int{}
	// truckSize := int()

	result := maximumUnits(boxTypes, truckSize)
	fmt.Printf("result = %v\n", result)
}

