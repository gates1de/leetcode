package main
import (
	"fmt"
	"sort"
)

func sumOfDistancesInTree(n int, edges [][]int) []int {
	sort.Slice(edges, func (a, b int) bool {
		if edges[a][0] <= edges[b][0] {
			return true
		}
		return edges[a][1] < edges[b][1]
	})
	parent := map[int]int{}
	childCount := map[int]int{}
	for i := 0; i < n; i++ {
		parent[i] = -1
	}
	for _, e := range edges {
		child := e[0]
		p := e[1]
		if e[0] < e[1] || parent[e[0]] >= 0 {
			child = e[1]
			p = e[0]
		}
		parent[child] = p

		loopP := p
		for loopP >= 0 {
			childCount[loopP]++
			loopP = parent[loopP]
		}
	}

	// fmt.Printf("childCount = %v\n", childCount)
	result := make([]int, n)
	levelMap := map[int]int{}
	queue := []int{0}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		level := levelMap[current]
		for key, p := range parent {
			if p == current {
				queue = append(queue, key)
				levelMap[key] = level + 1
			}
		}
		result[0] += level
	}

	queue = []int{0}
	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		for key, p := range parent {
			if p == i {
				queue = append(queue, key)
			}
		}

		count := childCount[i]
		p := parent[i]
		if p == -1 {
			continue
		}
		if count > 0 {
			result[i] = result[p] + n - 2 - count * 2
		} else {
			result[i] = result[p] + n - 2
		}
	}

	return result
}

func main() {
	// result: [8,12,6,10,10,10]
	// n := int(6)
	// edges := [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}

	// result: [0]
	// n := int(1)
	// edges := [][]int{}

	// result: [1,1]
	// n := int(2)
	// edges := [][]int{{1, 0}}

	// result: [12,14,20,20,12,18,18,18]
	// n := int(8)
	// edges := [][]int{{0, 1}, {1, 2}, {1, 3}, {0, 4}, {4, 5}, {4, 6}, {4, 7}}

	// result: 
	n := int(3)
	edges := [][]int{{2, 1}, {0, 2}}

	// result: 
	// n := int(0)
	// edges := [][]int{}

	result := sumOfDistancesInTree(n, edges)
	fmt.Printf("result = %v\n", result)
}

