package main
import (
	"fmt"
)

func validArrangement(pairs [][]int) [][]int {
	count := make(map[int]int)
	graph := make(map[int][]int)

	for _, p := range pairs {
		graph[p[0]] = append(graph[p[0]], p[1])
		count[p[0]]++
		count[p[1]]--
	}

	start := pairs[0][0]
	for vertex, inOutDegree := range count {
		if inOutDegree > 0 {
			start = vertex
			break
		}
	}

	stack := []int{start}
	path := make([]int, 0, len(graph))

	for len(stack) > 0 {
		cur := stack[len(stack) - 1]

		if len(graph[cur]) > 0 {
			last := len(graph[cur]) - 1

			stack = append(stack, graph[cur][last])
			graph[cur] = graph[cur][:last]
		} else {
			last := len(stack) - 1

			path = append(path, stack[last])
			stack = stack[:last]
		}
	}

	result := make([][]int, 0, len(graph))
	for i := len(path) - 1; i > 0; i-- {
		result = append(result, []int{path[i], path[i - 1]})
	}

	return result
}

func main() {
	// result: [[11,9],[9,4],[4,5],[5,1]]
	// pairs := [][]int{{5,1},{4,5},{11,9},{9,4}}

	// result: [[1,3],[3,2],[2,1]]
	// pairs := [][]int{{1,3},{3,2},{2,1}}

	// result: [[1,2],[2,1],[1,3]]
	pairs := [][]int{{1,2},{1,3},{2,1}}

	// result: []
	// pairs := [][]int{}

	result := validArrangement(pairs)
	fmt.Printf("result = %v\n", result)
}

