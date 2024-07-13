package main
import (
	"fmt"
	"sort"
)

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	n := len(positions)
	indices := make([]int, n)
	result := make([]int, 0, n)
	stack := make([]int, 0, n)

	for i := 0; i < n; i++ {
		indices[i] = i
	}

	sort.Slice(indices, func(a, b int) bool { return positions[indices[a]] < positions[indices[b]] })

	for _, i := range indices {
		if directions[i] == 'R' {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && healths[i] > 0 {
				topIndex := stack[len(stack) - 1]
				stack = stack[:len(stack) - 1]

				if healths[topIndex] > healths[i] {
					healths[topIndex]--
					healths[i] = 0
					stack = append(stack, topIndex)
				} else if healths[topIndex] < healths[i] {
					healths[i]--
					healths[topIndex] = 0
				} else {
					healths[i] = 0
					healths[topIndex] = 0
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		if healths[i] > 0 {
			result = append(result, healths[i])
		}
	}

	return result
}

func main() {
	// result: [2,17,9,15,10]
	// positions := []int{5,4,3,2,1}
	// healths := []int{2,17,9,15,10}
	// directions := "RRRRR"

	// result: [14]
	positions := []int{3,5,2,6}
	healths := []int{10,10,15,12}
	directions := "RLRL"

	// result: []
	// positions := []int{1,2,5,6}
	// healths := []int{10,10,11,11}
	// directions := "RLRL"

	// result: []
	// positions := []int{}
	// healths := []int{}
	// directions := ""

	result := survivedRobotsHealths(positions, healths, directions)
	fmt.Printf("result = %v\n", result)
}

