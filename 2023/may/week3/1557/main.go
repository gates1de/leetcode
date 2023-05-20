package main
import (
	"fmt"
)

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	isIncomingEdgeExists := make([]bool, n)
	for _, edge := range edges {
      isIncomingEdgeExists[edge[1]] = true;
	}

	requiredNodes := make([]int, 0, n)
	for i := 0; i < n; i++ {
      if !isIncomingEdgeExists[i] {
        requiredNodes = append(requiredNodes, i)
      }
    }

    return requiredNodes
}

func main() {
	// result: [0,3]
	// n := int(6)
	// edges := [][]int{{0,1},{0,2},{2,5},{3,4},{4,2}}

	// result: [0,2,3]
	n := int(5)
	edges := [][]int{{0,1},{2,1},{3,1},{1,4},{2,4}}

	// result: 
	// n := int(0)
	// edges := [][]int{}

	result := findSmallestSetOfVertices(n, edges)
	fmt.Printf("result = %v\n", result)
}

