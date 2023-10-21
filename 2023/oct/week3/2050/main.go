package main
import (
	"fmt"
)

func minimumTime(n int, relations [][]int, time []int) int {
	graph := make(map[int][]int)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, 0)
	}
	for _, edge := range relations {
		x := edge[0] - 1
		y := edge[1] - 1
		graph[x] = append(graph[x], y)
	}

	memo := make(map[int]int)
	result := int(0)
	for node := 0; node < n; node++ {
		result = max(result, dfs(node, time, memo, graph))
	}

	return result
}

func dfs(node int, time []int, memo map[int]int, graph map[int][]int) int {
	if memo[node] > 0 {
		return memo[node]
	}

	if len(graph[node]) == 0 {
		return time[node]
	}

	result := int(0)
	for _, neighbor := range graph[node] {
		result = max(result, dfs(neighbor, time, memo, graph))
	}

	memo[node] = time[node] + result
	return memo[node]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 8
	// n := int(3)
	// relations := [][]int{{1,3},{2,3}}
	// time := []int{3,2,5}

	// result: 12
	n := int(5)
	relations := [][]int{{1,5},{2,5},{3,5},{3,4},{4,5}}
	time := []int{1,2,3,4,5}

	// result: 
	// n := int(0)
	// relations := [][]int{}
	// time := []int{}

	result := minimumTime(n, relations, time)
	fmt.Printf("result = %v\n", result)
}

