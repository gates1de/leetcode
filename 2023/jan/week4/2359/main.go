package main
import (
	"fmt"
	"math"
)

func closestMeetingNode(edges []int, node1 int, node2 int) int {
	n := len(edges)
	distances1 := make([]int, n)
	distances2 := make([]int, n)
	for i := 0; i < n; i++ {
		distances1[i] = math.MaxInt32
		distances2[i] = math.MaxInt32
	}
	distances1[node1] = 0
	distances2[node2] = 0
	visited1 := make([]bool, n)
	visited2 := make([]bool, n)

    dfs(node1, edges, distances1, visited1)
    dfs(node2, edges, distances2, visited2)
	result := int(-1)
	minDistanceTillNow := math.MaxInt32
	for currentNode := 0; currentNode < n; currentNode++ {
		if minDistanceTillNow > max(distances1[currentNode], distances2[currentNode]) {
			result = currentNode
			minDistanceTillNow = max(distances1[currentNode], distances2[currentNode])
		}
	}

	return result
}

func dfs(node int, edges []int, distances []int, visited []bool) {
	visited[node] = true
	neighbor := edges[node]
	if neighbor != -1 && !visited[neighbor] {
		distances[neighbor] = distances[node] + 1
		dfs(neighbor, edges, distances, visited)
	}
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// Wrong Answer
func ngSolution(edges []int, node1 int, node2 int) int {
	n := len(edges)
	nodes := make(map[int][]int, n)
	for i, edge := range edges {
		if edge == -1 {
			continue
		}

		if nodes[i] == nil {
			nodes[i] = make([]int, 0, n)
		}
		if nodes[edge] == nil {
			nodes[edge] = make([]int, 0, n)
		}
		nodes[i] = append(nodes[i], edge)
		nodes[edge] = append(nodes[edge], i)
	}

	visited := make(map[int]bool, n)
	depths := make(map[int]int, n)
	depths[node1] = 0
	visited[node1] = true
	queue := make([]int, 0, n)
	for _, node := range nodes[node1] {
		queue = append(queue, node)
		depths[node] = 1
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}
		visited[current] = true

		if nodes[current] != nil {
			for _, node := range nodes[current] {
				if node == node1 {
					continue
				}

				queue = append(queue, node)
				if depths[node] == 0 {
					depths[node] = depths[current] + 1
				} else {
					depths[node] = min(depths[node], depths[current] + 1)
				}
			}
		}
	}

	return depths[node2]
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 2
	// edges := []int{2,2,3,-1}
	// node1 := int(0)
	// node2 := int(1)

	// result: 2
	edges := []int{1,2,-1}
	node1 := int(0)
	node2 := int(2)

	// result: 
	// edges := []int{}
	// node1 := int(0)
	// node2 := int(0)

	result := closestMeetingNode(edges, node1, node2)
	fmt.Printf("result = %v\n", result)
}

