package main
import (
	"fmt"
)

func magnificentSets(n int, edges [][]int) int {
	adjList := make([][]int, n)
	for i, _ := range adjList {
		adjList[i] = make([]int, 0)
	}

	parent := make([]int, n)
	depth := make([]int, n)
	for i, _ := range parent {
		parent[i] = -1
	}

	for _, edge := range edges {
		adjList[edge[0] - 1] = append(adjList[edge[0] - 1], edge[1] - 1)
		adjList[edge[1] - 1] = append(adjList[edge[1] - 1], edge[0] - 1)
		union(edge[0] - 1, edge[1] - 1, parent, depth)
	}

	groups := make(map[int]int)
	for node := 0; node < n; node++ {
		numOfGroups := getNumOfGroups(adjList, node, n)
		if numOfGroups == -1 {
			return -1
		}

		root := find(node, parent)
		groups[root] = max(groups[root], numOfGroups)
	}

	result := int(0)
	for _, numOfGroups := range groups {
		result += numOfGroups
	}

	return result
}

func getNumOfGroups(adjList [][]int, node int, n int) int {
	queue := make([]int, 0)
	seen := make([]int, n)
	for i, _ := range seen {
		seen[i] = -1
	}
	queue = append(queue, node)
	seen[node] = 0
	result := int(0)

	for len(queue) > 0 {
		numOfNodesInLayer := len(queue)
		for i := 0; i < numOfNodesInLayer; i++ {
			current := queue[0]
			queue = queue[1:]

			for _, neighbor := range adjList[current] {
				if seen[neighbor] == -1 {
					seen[neighbor] = result + 1
					queue = append(queue, neighbor)
				} else {
					if seen[neighbor] == result {
						return -1
					}
				}
			}
		}

		result++
	}

	return result
}

func find(node int, parent []int) int {
	for parent[node] != -1 {
		node = parent[node]
	}
	return node
}

func union(node1 int, node2 int, parent []int, depth []int) {
	node1 = find(node1, parent)
	node2 = find(node2, parent)

	if node1 == node2 {
		return
	}

	if depth[node1] < depth[node2] {
		node1, node2 = node2, node1
	}
	parent[node2] = node1

	if depth[node1] == depth[node2] {
		depth[node1]++
	}
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 4
	// n := int(6)
	// edges := [][]int{{1,2},{1,4},{1,5},{2,6},{2,3},{4,6}}

	// result: -1
	n := int(3)
	edges := [][]int{{1,2},{2,3},{3,1}}

	// result: 
	// n := int(0)
	// edges := [][]int{}

	result := magnificentSets(n, edges)
	fmt.Printf("result = %v\n", result)
}
