package main
import (
	"fmt"
)

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	adjacents := make(map[int][][]int)
	for _, e := range redEdges {
		if adjacents[e[0]] == nil {
			adjacents[e[0]] = make([][]int, 0, n)
		}
		adjacents[e[0]] = append(adjacents[e[0]], []int{e[1], 0})
	}
	for _, e := range blueEdges {
		if adjacents[e[0]] == nil {
			adjacents[e[0]] = make([][]int, 0, n)
		}
		adjacents[e[0]] = append(adjacents[e[0]], []int{e[1], 1})
	}

	result := make([]int, n)
	for i, _ := range result {
		result[i] = -1
	}

	visited := make([][]bool, n)
	for i, _ := range visited {
		visited[i] = make([]bool, 2)
	}

	queue := make([][]int, 0, n)
	queue = append(queue, []int{0, 0, -1})
	result[0] = 0
	visited[0][1] = true
	visited[0][0] = true

	for len(queue) > 0 {
		element := queue[0]
		queue = queue[1:]
		node := element[0]
		steps := element[1]
		prevColor := element[2]

		if adjacents[node] == nil {
			continue
		}

		for _, adjacent := range adjacents[node] {
			neighbor := adjacent[0]
			color := adjacent[1]

			if !visited[neighbor][color] && color != prevColor {
				if result[neighbor] == -1 {
					result[neighbor] = 1 + steps
				}
				visited[neighbor][color] = true
				queue = append(queue, []int{neighbor, 1 + steps, color})
			}
		}
	}

	return result
}

func main() {
	// result: [0,1,-1]
	// n := int(3)
	// redEdges := [][]int{{0,1},{1,2}}
	// blueEdges := [][]int{}

	// result: [0,1,-1]
	n := int(3)
	redEdges := [][]int{{0,1}}
	blueEdges := [][]int{{2,1}}

	// result: 
	// n := int(0)
	// redEdges := [][]int{}
	// blueEdges := [][]int{}

	result := shortestAlternatingPaths(n, redEdges, blueEdges)
	fmt.Printf("result = %v\n", result)
}

