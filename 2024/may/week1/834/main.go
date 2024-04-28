package main
import (
	"fmt"
)

func sumOfDistancesInTree(n int, edges [][]int) []int {
    adjList := make(map[int][]int)

    for _, edge := range edges {
        adjList[edge[0]] = append(adjList[edge[0]], edge[1])
        adjList[edge[1]] = append(adjList[edge[1]], edge[0])
    }

    count := make([]int, n)
    for i := 0; i < len(count); i++ {
        count[i] = 1
    }

    result := make([]int, n)

    dfs(0, -1, adjList, count, result)
    dfs2(0, -1, adjList, count, result)

    return result
}

func dfs(node int, parent int, adjList map[int][]int, count []int, result []int) {
    for _, child := range adjList[node] {
        if child == parent {
            continue
        }

        dfs(child, node, adjList, count, result)
        count[node] += count[child]
        result[node] += result[child] + count[child]
    }
}

func dfs2(node int, parent int, adjList map[int][]int, count []int, result []int) {
    for _, child := range adjList[node] {
        if child == parent {
            continue
        }

        result[child] = result[node] - count[child] + len(count) - count[child]
        dfs2(child, node, adjList, count, result)
    }
}

func main() {
	// result: [8,12,6,10,10,10]
	// n := int(6)
	// edges := [][]int{{0,1},{0,2},{2,3},{2,4},{2,5}}

	// result: [0]
	// n := int(1)
	// edges := [][]int{}

	// result: [1,1]
	n := int(2)
	edges := [][]int{{1,0}}

	// result: 
	// n := int(0)
	// edges := [][]int{}

	result := sumOfDistancesInTree(n, edges)
	fmt.Printf("result = %v\n", result)
}

