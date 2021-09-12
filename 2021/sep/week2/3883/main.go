package main
import (
	"fmt"
	"sort"
	heap "container/heap"
)

// REF: https://github.com/carrot369/LeetCode-in-Go/blob/master/Algorithms/0882.reachable-nodes-in-subdivided-graph/reachable-nodes-in-subdivided-graph.go
func reachableNodes(edges [][]int, maxMoves int, n int) int {
	nodes := make(map[int]int, len(edges))
	nextTo := make([][]int, n)
	for i := range nextTo {
		nextTo[i] = make([]int, 0, 16)
	}
	for _, e := range edges {
		i, j, cnt := e[0], e[1], e[2]
		nodes[encode(i, j)] = cnt
		nextTo[i] = append(nextTo[i], j)
		nextTo[j] = append(nextTo[j], i)
	}

	pq := make(PQ, 1, n * 2)
	pq[0] = []int{maxMoves, 0}

	seen := make([]bool, n)
	maxRemainMoves := make([]int, n)

	res := 0
	for len(pq) > 0 {
		m := pq[0][0]
		i := pq[0][1]
		heap.Pop(&pq)

		if seen[i] {
			continue
		}

		seen[i] = true
		maxRemainMoves[i] = m
		res++

		for _, j := range nextTo[i] {
			if seen[j] {
				continue
			}
			cnt := nodes[encode(i, j)]
			jRemainMoves := m - cnt - 1
			if jRemainMoves >= 0 {
				heap.Push(&pq, []int{jRemainMoves, j})
			}
		}
	}

	for _, e := range edges {
		i, j, n := e[0], e[1], e[2]
		mi := maxRemainMoves[i]
		mj := maxRemainMoves[j]
		res += min(mi + mj, n)
	}

	return res
}

func encode(i, j int) int {
	if i > j {
		i, j = j, i
	}
	return i<<16 | j
}

type PQ [][]int

func (pq PQ) Len() int { return len(pq) }

func (pq PQ) Less(i, j int) bool {
	return pq[i][0] > pq[j][0]
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PQ) Push(x interface{}) {
	temp := x.([]int)
	*pq = append(*pq, temp)
}

func (pq *PQ) Pop() interface{} {
	temp := (*pq)[len(*pq)-1]
	*pq = (*pq)[0 : len(*pq)-1]
	return temp
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Wrong Answer
func ngSolution(edges [][]int, maxMoves int, n int) int {
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

