package main
import (
	"fmt"
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	indegree := make([]int, numCourses)
	adjacents := make([][]int, numCourses)
	for i, _ := range adjacents {
		adjacents[i] = make([]int, 0)
	}

	for _, p := range prerequisites {
		a := p[0]
		b := p[1]
		adjacents[b] = append(adjacents[b], a)
		indegree[a]++
	}

	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	nodesVisited := int(0)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		nodesVisited++

		for _, neighbor := range adjacents[node] {
			indegree[neighbor]--
			if indegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return nodesVisited == numCourses
}

func main() {
	// result: true
	// numCourses := int(2)
	// prerequisites := [][]int{{1,0}}

	// result: false
	numCourses := int(2)
	prerequisites := [][]int{{1,0},{0,1}}

	// result: 
	// numCourses := int(0)
	// prerequisites := [][]int{}

	result := canFinish(numCourses, prerequisites)
	fmt.Printf("result = %v\n", result)
}

