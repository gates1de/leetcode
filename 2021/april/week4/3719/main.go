package main
import (
	"fmt"
)

func criticalConnections(n int, connections [][]int) [][]int {
	graph := make([][]int, n)
	for _, conn := range connections {
		graph[conn[0]] = append(graph[conn[0]], conn[1])
		graph[conn[1]] = append(graph[conn[1]], conn[0])
	}

	visited := make(map[int]int)
	timer := int(0)
	critical := make([][]int, 0)

	dfs(0, 0, graph, visited, timer, &critical)

	return critical
}

func dfs(v int, p int, graph [][]int, visited map[int]int, timer int, critical *[][]int) {
	visited[v] = timer
	currentTime := timer

	for i := 0; i < len(graph[v]); i++ {
		to := graph[v][i]
		if to == p {
			continue
		}

		if _, ok := visited[to]; !ok {
			dfs(to, v, graph, visited, timer + 1, critical)
		}
		visited[v] = min(visited[to], visited[v])
		if currentTime < visited[to] {
			*critical = append(*critical, []int{v, to})
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// result: [[1, 3]]
	n := int(4)
	connections := [][]int{{0, 1}, {1, 2}, {2, 0}, {1, 3}}

	// result: 
	// n := int()
	// connections := [][]int{}

	result := criticalConnections(n, connections)
	fmt.Printf("result = %v\n", result)
}

