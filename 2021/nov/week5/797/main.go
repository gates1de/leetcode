package main
import (
	"fmt"
)

func allPathsSourceTarget(graph [][]int) [][]int {
	result := [][]int{}
	helper([]int{0}, graph, &result)
	return result
}

func helper(path []int, graph [][]int, result *[][]int) {
	if len(path) > 0 && (path[len(path) - 1] == len(graph) - 1) {
		// fmt.Printf("path = %v, len(graph) - 1 = %v\n", path, len(graph) - 1)
		*result = append(*result, path)
		return
	}

	current := int(0)
	if len(path) > 0 {
		current = path[len(path) - 1]
	}

	for _, to := range graph[current] {
		newPath := make([]int, len(path) + 1)
		copy(newPath, path)
		newPath[len(newPath) - 1] = to
		helper(newPath, graph, result)
	}
}

func main() {
	// result: [[0,1,3],[0,2,3]]
	// graph := [][]int{{1,2},{3},{3},{}}

	// result: [[0,4],[0,3,4],[0,1,3,4],[0,1,2,3,4],[0,1,4]]
	// graph := [][]int{{4,3,1},{3,2,4},{3},{4},{}}

	// result: [[0,1]]
	// graph := [][]int{{1},{}}

	// result: [[0,1,2,3],[0,2,3],[0,3]]
	// graph := [][]int{{1,2,3},{2},{3},{}}

	// result: [[0,1,2,3],[0,3]]
	// graph := [][]int{{1,3},{2},{3},{}}

	// result: [[0,3,6,7],[0,3,4,7],[0,3,4,6,7],[0,3,4,5,6,7],[0,1,4,7],[0,1,4,6,7],[0,1,4,5,6,7],[0,1,6,7],[0,1,7],[0,1,2,4,7],[0,1,2,4,6,7],[0,1,2,4,5,6,7],[0,1,2,6,7],[0,1,2,3,6,7],[0,1,2,3,4,7],[0,1,2,3,4,6,7],[0,1,2,3,4,5,6,7],[0,1,5,6,7]]
	graph := [][]int{{3,1},{4,6,7,2,5},{4,6,3},{6,4},{7,6,5},{6},{7},{}}

	// result: 
	// graph := [][]int{}

	result := allPathsSourceTarget(graph)
	fmt.Printf("result = %v\n", result)
}

