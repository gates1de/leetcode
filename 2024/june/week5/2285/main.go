package main
import (
	"fmt"
	"sort"
)

func maximumImportance(n int, roads [][]int) int64 {
	degree := make([]int64, n)
	for _, edge := range roads {
		degree[edge[0]]++
		degree[edge[1]]++
	}

	sort.Slice(degree, func (a, b int) bool { return degree[a] < degree[b] })
	val := int64(1)
	result := int64(0)

	for _, d := range degree {
		result += val * d
		val++
	}

	return result
}

func main() {
	// result: 43
	// n := int(5)
	// roads := [][]int{{0,1},{1,2},{2,3},{0,2},{1,3},{2,4}}

	// result: 20
	n := int(5)
	roads := [][]int{{0,3},{2,4},{1,3}}

	// result: 
	// n := int(0)
	// roads := [][]int{}

	result := maximumImportance(n, roads)
	fmt.Printf("result = %v\n", result)
}

