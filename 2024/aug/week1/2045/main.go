package main
import (
	"fmt"
)

func secondMinimum(n int, edges [][]int, time int, change int) int {
	adj := make(map[int][]int)
	for _, edge := range edges {
		a := edge[0]
		b := edge[1]
		if adj[a] == nil {
			adj[a] = make([]int, 0)
		}
		if adj[b] == nil {
			adj[b] = make([]int, 0)
		}

		adj[a] = append(adj[a], b)
		adj[b] = append(adj[b], a)
	}

	dist1 := make([]int, n + 1)
	dist2 := make([]int, n + 1)

	for i := 1; i <= n; i++ {
		dist1[i] = -1
		dist2[i] = -1
	}

	queue := make([][]int, 0)
	queue = append(queue, []int{1, 1})
	dist1[1] = 0

	for len(queue) > 0 {
		temp := queue[0]
		queue = queue[1:]
		node := temp[0]
		freq := temp[1]

		timeTaken := dist1[node]
		if freq != 1 {
			timeTaken = dist2[node]
		}

		if (timeTaken / change) % 2 == 1 {
			timeTaken = change * (timeTaken / change + 1) + time
		} else {
			timeTaken += time
		}

		if adj[node] == nil {
			continue
		}

		for _, neighbor := range adj[node] {
			if dist1[neighbor] == -1 {
				dist1[neighbor] = timeTaken
				queue = append(queue, []int{neighbor, 1})
			} else if dist2[neighbor] == -1 && dist1[neighbor] != timeTaken {
				if neighbor == n {
					return timeTaken
				}

				dist2[neighbor] = timeTaken
				queue = append(queue, []int{neighbor, 2})
			}
		}
	}

	return 0
}

func main() {
	// result: 13
	// n := int(5)
	// edges := [][]int{{1,2},{1,3},{1,4},{3,4},{4,5}}
	// time := int(3)
	// change := int(5)

	// result: 11
	n := int(2)
	edges := [][]int{{1,2}}
	time := int(3)
	change := int(2)

	// result: 
	// n := int(0)
	// edges := [][]int{}
	// time := int(0)
	// change := int(0)

	result := secondMinimum(n, edges, time, change)
	fmt.Printf("result = %v\n", result)
}

