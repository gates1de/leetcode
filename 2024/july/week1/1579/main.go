package main
import (
	"fmt"
)

type UnionFind struct {
	Representative []int
	ComponentSize []int
	Components int
}

func constructor(size int) UnionFind {
	representative := make([]int, size + 1)
	componentSize := make([]int, size + 1)
	for i := 0; i <= size; i++ {
		representative[i] = i
		componentSize[i] = 1
	}
	return UnionFind{
		Representative: representative,
		ComponentSize: componentSize,
		Components: size,
	}
}

func (this *UnionFind) Find(node int) int {
	if this.Representative[node] != node {
		this.Representative[node] = this.Find(this.Representative[node])
	}
	return this.Representative[node]
}


func (this *UnionFind) Join(node1, node2 int) int {
	node1 = this.Find(node1)
	node2 = this.Find(node2)

	if node1 == node2 {
		return 0
	}

	if this.ComponentSize[node1] > this.ComponentSize[node2] {
		this.ComponentSize[node1] += this.ComponentSize[node2]
		this.Representative[node2] = node1
	} else {
		this.ComponentSize[node2] += this.ComponentSize[node1]
		this.Representative[node1] = node2
	}

	this.Components--
	return 1
}

func (this *UnionFind) IsConnected() bool {
	return this.Components == 1
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	aliceUf := constructor(n)
	bobUf := constructor(n)
	edgesRequired := int(0)
	for _, edge := range edges {
		if edge[0] == 3 {
			edgesRequired += aliceUf.Join(edge[1], edge[2]) | bobUf.Join(edge[1], edge[2])
		}
	}

	for _, edge := range edges {
		if edge[0] == 1 {
			edgesRequired += aliceUf.Join(edge[1], edge[2])
		} else if edge[0] == 2 {
			edgesRequired += bobUf.Join(edge[1], edge[2])
		}
	}

	if aliceUf.IsConnected() && bobUf.IsConnected() {
		return len(edges) - edgesRequired
	}
	return -1
}

// Wrong answer
func ngSolution(n int, edges [][]int) int {
	aliceResult := bfs(true, n, edges)
	if aliceResult == -1 {
		return -1
	}

	bobResult := bfs(false, n, edges)
	if bobResult == -1 {
		return -1
	}

	return aliceResult + bobResult
}

func bfs(isAlice bool, n int, edges [][]int) int {
	edgeCount := int(0)
	queue := make([]int, 0, n)
	visited := make([]bool, n + 1)
	adjacents := make([][]int, n + 1)

	for _, edge := range edges {
		t := edge[0]
		node1 := edge[1]
		node2 := edge[2]

		if (isAlice && t == 2) || (!isAlice && t == 1) {
			continue
		}

		if adjacents[node1] == nil {
			adjacents[node1] = make([]int, 0, n)
		}
		if adjacents[node2] == nil {
			adjacents[node2] = make([]int, 0, n)
		}

		adjacents[node1] = append(adjacents[node1], node2)
		adjacents[node2] = append(adjacents[node2], node1)
		edgeCount++

		if len(queue) == 0 {
			queue = append(queue, node1)
			visited[node1] = true
		}
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if adjacents[node] == nil {
			continue
		}

		for _, neighbor := range adjacents[node] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
			visited[neighbor] = true
		}
	}

	if edgeCount + 1 < n {
		return -1
	}

	return edgeCount + 1 - n
}

func main() {
	// result: 2
	// n := int(4)
	// edges := [][]int{{3,1,2},{3,2,3},{1,1,3},{1,2,4},{1,1,2},{2,3,4}}

	// result: 0
	// n := int(4)
	// edges := [][]int{{3,1,2},{3,2,3},{1,1,4},{2,1,4}}

	// result: -1
	n := int(4)
	edges := [][]int{{3,2,3},{1,1,2},{2,3,4}}

	// result: 
	// n := int(0)
	// edges := [][]int{}

	result := maxNumEdgesToRemove(n, edges)
	fmt.Printf("result = %v\n", result)
}


