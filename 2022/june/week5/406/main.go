package main
import (
	"fmt"
	"sort"
)

func reconstructQueue(people [][]int) [][]int {
	n := len(people)
	queue := make([][]int, n)
	queue[0] = people[0]

	sort.Slice(people, func (a, b int) bool {
		if people[a][1] == 0 {
			if people[b][1] == 0 {
				return people[a][0] < people[b][0]
			}
			return true
		}

		if people[a][1] == people[b][1] {
			return people[a][0] > people[b][0]
		}
		return people[a][1] < people[b][1]
	})

	TOP:
	for i, person := range people {
		if person[1] == 0 {
			queue[i] = person
			continue
		}

		tallerCount := person[1]
		for j := 0; j < i; j++ {
			if tallerCount == 0 {
				copy(queue[j + 1:], queue[j:])
				queue[j] = person
				continue TOP
			}

			queuedPerson := queue[j]

			if person[0] <= queuedPerson[0] {
				tallerCount--
			}
		}
		queue[i] = person
	}

	return queue
}

func main() {
	// result: [[5,0],[7,0],[5,2],[6,1],[4,4],[7,1]]
	// people := [][]int{{7,0},{4,4},{7,1},{5,0},{6,1},{5,2}}

	// result: [[4,0],[5,0],[2,2],[3,2],[1,4],[6,0]]
	// people := [][]int{{6,0},{5,0},{4,0},{3,2},{2,2},{1,4}}

	// result: [[1,0],[2,0],[9,0],[3,1],[1,4],[9,1],[4,2],[7,2],[8,2],[4,5]]
	people := [][]int{{8,2},{4,2},{4,5},{2,0},{7,2},{1,4},{9,1},{3,1},{9,0},{1,0}}

	// result: 
	// people := [][]int{}

	result := reconstructQueue(people)
	fmt.Printf("result = %v\n", result)
}

