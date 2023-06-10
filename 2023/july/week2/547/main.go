package main
import (
	"fmt"
)

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	result := int(0)
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		if !visited[i] {
			result++
			dfs(i, isConnected, visited)
		}
	}

	return result
}

func dfs(node int, isConnected [][]int, visited []bool) {
	visited[node] = true
	for i := 0; i < len(isConnected); i++ {
		if isConnected[node][i] == 1 && !visited[i] {
			dfs(i, isConnected, visited)
		}
	}
}

func main() {
	// result: 2
	// isConnected := [][]int{{1,1,0},{1,1,0},{0,0,1}}

	// result: 3
	isConnected := [][]int{{1,0,0},{0,1,0},{0,0,1}}

	// result: 
	// isConnected := [][]int{}

	result := findCircleNum(isConnected)
	fmt.Printf("result = %v\n", result)
}

