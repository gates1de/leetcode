package main
import (
	"fmt"
	"sort"
)

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) <= 1 {
		return false
	}

	sort.Slice(coordinates, func(a, b int) bool {
		if coordinates[a][0] == coordinates[b][0] {
			return coordinates[a][1] < coordinates[b][1]
		}
		return coordinates[a][0] < coordinates[b][0]
	})

	startDiffX := coordinates[1][0] - coordinates[0][0]
	startDiffY := coordinates[1][1] - coordinates[0][1]

	for i := 2; i < len(coordinates); i++ {
		diffX := coordinates[i][0] - coordinates[0][0]
		diffY := coordinates[i][1] - coordinates[0][1]

		if startDiffY * diffX != startDiffX * diffY {
			return false
		}
	}

	return true
}

func main() {
	// result: true
	// coordinates := [][]int{{1,2},{2,3},{3,4},{4,5},{5,6},{6,7}}

	// result: false
	// coordinates := [][]int{{1,1},{2,2},{3,4},{4,5},{5,6},{7,7}}

	// result: true
	coordinates := [][]int{{2,4},{4,8},{6,12}}

	// result: 
	// coordinates := [][]int{}

	result := checkStraightLine(coordinates)
	fmt.Printf("result = %v\n", result)
}

