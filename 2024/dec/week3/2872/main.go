package main
import (
	"fmt"
)

func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	if n < 2 {
		return 1
	}

	result := int(0)
	graph := make([][]int, 0, n)
	inDegree := make([]int, n)

	for i := 0; i < n; i++ {
		graph = append(graph, make([]int, 0))
	}

	for _, edge := range edges {
		node1 := edge[0]
		node2 := edge[1]
		graph[node1] = append(graph[node1], node2)
		graph[node2] = append(graph[node2], node1)
		inDegree[node1]++
		inDegree[node2]++
	}

	longValues := make([]int64, n)
	for i, value := range values {
		longValues[i] = int64(value)
	}

	queue := make([]int, 0)
	for node := 0; node < n; node++ {
		if inDegree[node] == 1 {
			queue = append(queue, node)
		}
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		inDegree[current]--

		addValue := int64(0)

		if longValues[current] % int64(k) == 0 {
			result++
		} else {
			addValue = longValues[current]
		}

		for _, neighbor := range graph[current] {
			if inDegree[neighbor] == 0 {
				continue
			}

			inDegree[neighbor]--
			longValues[neighbor] += addValue

			if inDegree[neighbor] == 1 {
				queue = append(queue, neighbor)
			}
		}
	}

	return result
}

func main() {
	// result: 2
	// n := int(5)
	// edges := [][]int{{0,2},{1,2},{1,3},{2,4}}
	// values := []int{1,8,1,4,4}
	// k := int(6)

	// result: 3
	n := int(7)
	edges := [][]int{{0,1},{0,2},{1,3},{1,4},{2,5},{2,6}}
	values := []int{3,0,6,1,5,2,1}
	k := int(3)

	// result: 
	// n := int(0)
	// edges := [][]int{}
	// values := []int{}
	// k := int(0)

	result := maxKDivisibleComponents(n, edges, values, k)
	fmt.Printf("result = %v\n", result)
}

