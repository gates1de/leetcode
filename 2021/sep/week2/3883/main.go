package main
import (
	"fmt"
	"sort"
)

func reachableNodes(edges [][]int, maxMoves int, n int) int {
	if maxMoves == 0 {
		return 0
	}

	sort.Slice(edges, func (a, b int) bool { return edges[a][0] < edges[b][0] })
	sum := int(0)
	for _, edge := range edges {
		sum++
		sum += edge[2]
	}

	// fmt.Printf("sum = %v\n", sum)
	visited := map[[2]int]bool{}
	helper(0, 0, maxMoves, visited, edges)
	afterSum := int(0)
	// fmt.Printf("edges = %v\n", edges)
	for _, edge := range edges {
		if !visited[[2]int{edge[0], -1}] || !visited[[2]int{edge[1], -1}] {
			afterSum++
		}
		afterSum += edge[2]
	}

	if afterSum == sum {
		return 1
	}

	return sum - afterSum
}

func helper(head int, current int, maxMoves int, visited map[[2]int]bool, edges [][]int) {
	// fmt.Printf("head = %v, current = %v, edges = %v\n", head, current, edges)
	visited[[2]int{head, -1}] = true
	if current >= maxMoves {
		return
	}

	for _, edge := range edges {
		if head == edge[0] {
			if visited[[2]int{edge[1], edge[0]}] == true || edge[2] == 0 {
				continue
			}

			visited[[2]int{edge[1], edge[0]}] = true
			cnt := edge[2]
			if current + cnt + 1 > maxMoves {
				edge[2] -= maxMoves - current
				continue
			}

			edge[2] = 0
			helper(edge[1], current + cnt + 1, maxMoves, visited, edges)
		} else if head == edge[1] {
			if visited[[2]int{edge[0], edge[1]}] == true || edge[2] == 0 {
				continue
			}

			visited[[2]int{edge[0], edge[1]}] = true
			cnt := edge[2]
			if current + cnt + 1 > maxMoves {
				edge[2] -= maxMoves - current
				continue
			}

			edge[2] = 0
			helper(edge[0], current + cnt + 1, maxMoves, visited, edges)
		}
	}
}

func main() {
	// result: 13
	// edges := [][]int{{0,1,10},{0,2,1},{1,2,2}}
	// maxMoves := int(6)
	// n := int(3)

	// result: 23
	// edges := [][]int{{0,1,4},{1,2,6},{0,2,8},{1,3,1}}
	// maxMoves := int(10)
	// n := int(4)

	// result: 1
	// edges := [][]int{{1,2,4},{1,4,5},{1,3,1},{2,3,4},{3,4,5}}
	// maxMoves := int(17)
	// n := int(5)

	// result: 14
	edges := [][]int{{0,1,10},{0,2,1},{1,2,2},{1,3,3}}
	maxMoves := int(6)
	n := int(4)

	// result: 
	// edges := [][]int{}
	// maxMoves := int(0)
	// n := int(0)

	result := reachableNodes(edges, maxMoves, n)
	fmt.Printf("result = %v\n", result)
}

