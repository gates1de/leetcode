package main

import (
	"fmt"
	"sort"
)

func maxBuilding(n int, restrictions [][]int) int {
	restrictions = append(restrictions, []int{1, 0}, []int{n, n - 1})
	sort.Slice(restrictions, func(i, j int) bool {
		if restrictions[i][0] == restrictions[j][0] {
			return restrictions[i][1] < restrictions[j][1]
		}
		return restrictions[i][0] < restrictions[j][0]
	})

	for i := 1; i < len(restrictions); i++ {
		d := restrictions[i][0] - restrictions[i-1][0]
		limit := restrictions[i-1][1] + d
		if restrictions[i][1] > limit {
			restrictions[i][1] = limit
		}
	}

	for i := len(restrictions) - 2; i >= 0; i-- {
		d := restrictions[i+1][0] - restrictions[i][0]
		limit := restrictions[i+1][1] + d
		if restrictions[i][1] > limit {
			restrictions[i][1] = limit
		}
	}

	result := int(0)
	for i := 1; i < len(restrictions); i++ {
		leftID, leftH := restrictions[i-1][0], restrictions[i-1][1]
		rightID, rightH := restrictions[i][0], restrictions[i][1]
		d := rightID - leftID
		peak := (leftH + rightH + d) / 2
		result = max(result, peak)
	}

	return result
}

func main() {
	// result: 2
	// n := int(5)
	// restrictions := [][]int{{2,1},{4,1}}

	// result: 5
	// n := int(6)
	// restrictions := [][]int{}

	// result: 5
	n := int(10)
	restrictions := [][]int{{5,3},{2,5},{7,4},{10,3}}

	// result: 
	// n := int(0)
	// restrictions := [][]int{}

	result := maxBuilding(n, restrictions)
	fmt.Printf("result = %v\n", result)
}

