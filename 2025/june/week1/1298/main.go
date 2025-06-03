package main
import (
	"fmt"
)

func maxCandies(status []int, candies []int, keys [][]int, containedBoxes [][]int, initialBoxes []int) int {
	n := len(status)
	canOpen := make([]bool, n)
	hasBox := make([]bool, n)
	used := make([]bool, n)

	for i := 0; i < n; i++ {
		canOpen[i] = (status[i] == 1)
	}

	queue := []int{}
	result := int(0)
	for _, box := range initialBoxes {
		hasBox[box] = true

		if canOpen[box] {
      queue = append(queue, box)
			used[box] = true
			result += candies[box]
		}
	}

	for len(queue) > 0 {
		bigBox := queue[0]
		queue = queue[1:]

		for _, key := range keys[bigBox] {
			canOpen[key] = true

			if !used[key] && hasBox[key] {
				queue = append(queue, key)
				used[key] = true
				result += candies[key]
			}
		}

		for _, box := range containedBoxes[bigBox] {
			hasBox[box] = true

			if !used[box] && canOpen[box] {
				queue = append(queue, box)
				used[box] = true
				result += candies[box]
			}
		}
	}

	return result
}

func main() {
	// result: 16
	// status := []int{1,0,1,0}
	// candies := []int{7,5,4,100}
	// keys := [][]int{{},{},{1}, {}}
	// containedBoxes := [][]int{{1,2},{3},{},{}}
	// initialBoxes := []int{0}

	// result: 6
	status := []int{1,0,0,0,0,0}
	candies := []int{1,1,1,1,1,1}
	keys := [][]int{{1,2,3,4,5},{},{},{},{},{}}
	containedBoxes := [][]int{{1,2,3,4,5},{},{},{},{},{}}
	initialBoxes := []int{0}

	// result: 
	// status := []int{}
	// candies := []int{}
	// keys := [][]int{}
	// containedBoxes := [][]int{}
	// initialBoxes := []int{}

	result := maxCandies(status, candies, keys, containedBoxes, initialBoxes)
	fmt.Printf("result = %v\n", result)
}

