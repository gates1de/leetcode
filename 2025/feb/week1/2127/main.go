package main
import (
	"fmt"
)

func maximumInvitations(favorite []int) int {
	n := len(favorite)
	inDegree := make([]int, n)

	for _, person := range favorite {
		inDegree[person]++
	}

	queue := make([]int, 0, n)
	for i, val := range inDegree {
		if val == 0 {
			queue = append(queue, i)
		}
	}

	depth := make([]int, n)
	for i, _ := range depth {
		depth[i] = 1
	}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		next := favorite[current]
		depth[next] = max(depth[next], depth[current] + 1)

		inDegree[next]--
		if inDegree[next] == 0 {
			queue = append(queue, next)
		}
	}

	longestCycle := int(0)
	twoCycleInvitations := int(0)

	for i, _ := range favorite {
		if inDegree[i] == 0 {
			continue
		}

		cycleLength := int(0)
		current := i
		for inDegree[current] != 0 {
			inDegree[current] = 0
			cycleLength++
			current = favorite[current]
		}

		if cycleLength == 2 {
			twoCycleInvitations += depth[i] + depth[favorite[i]]
		} else {
			longestCycle = max(longestCycle, cycleLength)
		}
	}

	return max(longestCycle, twoCycleInvitations)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// favorite := []int{2,2,1,2}

	// result: 3
	// favorite := []int{1,2,0}

	// result: 4
	favorite := []int{3,0,1,4,1}

	// result: 
	// favorite := []int{}

	result := maximumInvitations(favorite)
	fmt.Printf("result = %v\n", result)
}

