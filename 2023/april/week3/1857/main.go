package main
import (
	"fmt"
)

func largestPathValue(colors string, edges [][]int) int {
	n := len(colors)
	adjacents := make(map[int][]int)
	indegree := make([]int, n)

	for _, edge := range edges {
		if adjacents[edge[0]] == nil {
			adjacents[edge[0]] = make([]int, 0)
		}
		adjacents[edge[0]] = append(adjacents[edge[0]], edge[1])
		indegree[edge[1]]++
	}

	counts := make([][26]int, n)
	queue := make([]int, 0)
	for i := 0; i < n; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	result := int(1)
	nodesSeen := int(0)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		counts[node][colors[node] - 'a']++
		result = max(result, counts[node][colors[node] - 'a'])
		nodesSeen++

		if adjacents[node] == nil {
			continue
		}

		for _, neighbor := range adjacents[node] {
			for i := 0; i < 26; i++ {
				counts[neighbor][i] = max(counts[neighbor][i], counts[node][i])
			}

			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	if nodesSeen < n {
		return -1
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// colors := "abaca"
	// edges := [][]int{{0,1},{0,2},{2,3},{3,4}}

	// result: -1
	colors := "a"
	edges := [][]int{{0,0}}

	// result: 
	// colors := ""
	// edges := [][]int{}

	result := largestPathValue(colors, edges)
	fmt.Printf("result = %v\n", result)
}

