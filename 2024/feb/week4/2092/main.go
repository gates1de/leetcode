package main
import (
	"fmt"
	"math"
)

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	graph := make(map[int][][]int)
	for _, meeting := range meetings {
		x := meeting[0]
		y := meeting[1]
		t := meeting[2]

		if graph[x] == nil {
			graph[x] = make([][]int, 0)
		}
		if graph[y] == nil {
			graph[y] = make([][]int, 0)
		}

		graph[x] = append(graph[x], []int{t, y})
		graph[y] = append(graph[y], []int{t, x})
	}

	earliest := make([]int, n)
	for i, _ := range earliest {
		earliest[i] = math.MaxInt32
	}
	earliest[0] = 0
	earliest[firstPerson] = 0

	dfs(0, 0, graph, earliest)
	dfs(firstPerson, 0, graph, earliest)

	result := make([]int, 0, n)
	for i, val := range earliest {
		if val != math.MaxInt32 {
			result = append(result, i)
		}
	}

	return result
}

func dfs(person int, time int, graph map[int][][]int, earliest []int) {
	if graph[person] == nil {
		return
	}

	for _, nextPersonTime := range graph[person] {
		t := nextPersonTime[0]
		nextPerson := nextPersonTime[1]
		if t >= time && earliest[nextPerson] > t {
			earliest[nextPerson] = t
			dfs(nextPerson, t, graph, earliest)
		}
	}
}

func main() {
	// result: [0,1,2,3,5]
	// n := int(6)
	// meetings := [][]int{{1,2,5},{2,3,8},{1,5,10}}
	// firstPerson := int(1)

	// result: [0,1,3]
	// n := int(4)
	// meetings := [][]int{{3,1,3},{1,2,2},{0,3,3}}
	// firstPerson := int(3)

	// result: [0,1,2,3,4]
	n := int(5)
	meetings := [][]int{{3,4,2},{1,2,1},{2,3,1}}
	firstPerson := int(1)

	// result: 
	// n := int(0)
	// meetings := [][]int{}
	// firstPerson := int(0)

	result := findAllPeople(n, meetings, firstPerson)
	fmt.Printf("result = %v\n", result)
}

