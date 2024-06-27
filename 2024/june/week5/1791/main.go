package main
import (
	"fmt"
)

func findCenter(edges [][]int) int {
	first := edges[0]
	second := edges[1]

	if first[0] == second[0] || first[0] == second[1] {
		return first[0]
	}

	return first[1]
}

func main() {
	// result: 2
	// edges := [][]int{{1,2},{2,3},{4,2}}

	// result: 1
	edges := [][]int{{1,2},{5,1},{1,3},{1,4}}

	// result: 
	// edges := [][]int{}

	result := findCenter(edges)
	fmt.Printf("result = %v\n", result)
}

