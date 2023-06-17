package main
import (
	"fmt"
)

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	adjacents := make([][]int, n)
	for i := 0; i < n; i++ {
		adjacents = append(adjacents, make([]int, 0))
	}
	for i := 0; i < n; i++ {
		if manager[i] != -1 {
			adjacents[manager[i]] = append(adjacents[manager[i]], i)
		}
	}

	queue := make([][2]int, 0)
	queue = append(queue, [2]int{headID, 0})
	result := int(0)
	for len(queue) > 0 {
		pair := queue[0]
		queue = queue[1:]
		parent := pair[0]
		time := pair[1]

		for _, adjacent := range adjacents[parent] {
			queue = append(queue, [2]int{adjacent, time + informTime[parent]})
			result = max(result, time + informTime[parent])
		}
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
	// result: 0
	// n := int(1)
	// headID := int(0)
	// manager := []int{-1}
	// informTime := []int{0}

	// result: 1
	n := int(6)
	headID := int(2)
	manager := []int{2,2,-1,2,2,2}
	informTime := []int{0,0,1,0,0,0}

	// result: 
	// n := int(0)
	// headID := int(0)
	// manager := []int{}
	// informTime := []int{}

	result := numOfMinutes(n, headID, manager, informTime)
	fmt.Printf("result = %v\n", result)
}

