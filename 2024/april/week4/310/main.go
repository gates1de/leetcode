package main
import (
	"fmt"
)

func findMinHeightTrees(n int, edges [][]int) []int {
    graph := make([][]int, n)
    for _,v := range edges {
        graph[v[0]] = append(graph[v[0]], v[1])
        graph[v[1]] = append(graph[v[1]], v[0])
    }

    visited1 := make([]bool, n)
    route1 := dfs(graph, 0, visited1)

    visited2 := make([]bool, n)
    longestRoute := dfs(graph, route1[0], visited2)

    return longestRoute[(len(longestRoute) - 1) / 2 : len(longestRoute) / 2 + 1]
}

func dfs(graph [][]int, n int, visited []bool) []int {
    visited[n] = true

    var res []int
    for _, v := range graph[n] {
        if !visited[v] {
            route := dfs(graph, v, visited)
            if len(route) > len(res) {
                res = route
            }
        }
    }

    return append(res, n)
}

func main() {
	// result: [1]
	// n := int(4)
	// edges := [][]int{{1,0},{1,2},{1,3}}

	// result: [3,4]
	// n := int(6)
	// edges := [][]int{{3,0},{3,1},{3,2},{3,4},{5,4}}

	// result: [0]
	// n := int(1)
	// edges := [][]int{}

	// result: [0,1]
	n := int(2)
	edges := [][]int{{0,1}}

	// result: 
	// n := int(0)
	// edges := [][]int{}

	result := findMinHeightTrees(n, edges)
	fmt.Printf("result = %v\n", result)
}


