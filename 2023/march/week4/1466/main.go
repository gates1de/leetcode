package main
import (
	"fmt"
)

type Neighbor struct {
	Child int
	Sign int
}

func minReorder(n int, connections [][]int) int {
	adjacents := make(map[int][]Neighbor)
	for _, connection := range connections {
		if adjacents[connection[0]] == nil {
			adjacents[connection[0]] = make([]Neighbor, 0)
		}
		if adjacents[connection[1]] == nil {
			adjacents[connection[1]] = make([]Neighbor, 0)
		}

		adjacents[connection[0]] = append(adjacents[connection[0]], Neighbor{Child: connection[1], Sign: 1})
		adjacents[connection[1]] = append(adjacents[connection[1]], Neighbor{Child: connection[0], Sign: 0})
	}

	result := int(0)
	dfs(0, -1, adjacents, &result)
	return result
}

func dfs(node int, parent int, adjacents map[int][]Neighbor, result *int) {
	if adjacents[node] == nil {
		return
	}

	for _, neighbor := range adjacents[node] {
		if neighbor.Child != parent {
			*result += neighbor.Sign
			dfs(neighbor.Child, node, adjacents, result)
		}
	}
}

func main() {
	// result: 3
	// n := int(6)
	// connections := [][]int{{0,1},{1,3},{2,3},{4,0},{4,5}}

	// result: 2
	// n := int(5)
	// connections := [][]int{{1,0},{1,2},{3,2},{3,4}}

	// result: 0
	n := int(3)
	connections := [][]int{{1,0},{2,0}}

	// result: 
	// n := int(0)
	// connections := [][]int{}

	result := minReorder(n, connections)
	fmt.Printf("result = %v\n", result)
}

