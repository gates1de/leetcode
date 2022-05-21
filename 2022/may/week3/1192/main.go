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

    visited := map[int]int{}
    result := [][]int{}

    dfs(0, 0, 0, graph, visited, &result)

    return result
}

func dfs(node int, pre int, count int, graph [][]int, visited map[int]int, result *[][]int) {
    visited[node] = count

    for i := 0; i < len(graph[node]); i++ {
        to := graph[node][i]
        if to == pre {
            continue
        }

        if _, ok := visited[to]; !ok {
            dfs(to, node, count + 1, graph, visited, result)
        }

        visited[node] = min(visited[to], visited[node])

        if count < visited[to] {
            *result = append(*result, []int{node, to})
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
	// result: [[1,3]]
	// n := int(4)
	// connections := [][]int{{0,1},{1,2},{2,0},{1,3}}

	// result: [[0,1]]
	// n := int(2)
	// connections := [][]int{{0,1}}

	// result: [[0,4],[0,3]]
	// n := int(5)
	// connections := [][]int{{0,4},{0,3},{2,1},{0,1},{0,2}}

	// result: [[1,3]]
	n := int(6)
	connections := [][]int{{0,1},{1,2},{2,0},{1,3},{3,4},{4,5},{5,3}}

	// result: 
	// n := int(0)
	// connections := [][]int{}

	result := criticalConnections(n, connections)
	fmt.Printf("result = %v\n", result)
}

