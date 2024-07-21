package main
import (
	"fmt"
)

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	orderRows := topologicalSort(rowConditions, k)
	orderColumns := topologicalSort(colConditions, k)

	if len(orderRows) == 0 || len(orderColumns) == 0 {
		return make([][]int, 0)
	}

	result := make([][]int, k)
	for i, _ := range result {
		result[i] = make([]int, k)

		for j := 0; j < k; j++ {
			if orderRows[i] == orderColumns[j] {
                result[i][j] = orderRows[i]
            }
		}
	}

	return result
}

func topologicalSort(edges [][]int, n int) []int {
	adj := make([][]int, n + 1)
	for i, _ := range adj {
		adj[i] = make([]int, 0)
	}

	deg := make([]int, n + 1)
	order := make([]int, n)

	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], edge[1])
		deg[edge[1]]++
	}

	queue := make([]int, 0, len(deg))
	for i := 1; i <= n; i++ {
		if deg[i] == 0 {
			queue = append(queue, i)
		}
	}

	index := int(0)
	for len(queue) > 0 {
		f := queue[0]
		queue = queue[1:]
		order[index] = f
		index++
		n--

		for _, v := range adj[f] {
			deg[v]--
			if deg[v] == 0 {
				queue = append(queue, v)
			}
		}
	}

	if n != 0 {
		return make([]int, 0)
	}

	return order
}

func main() {
	// result: [[3,0,0],[0,0,1],[0,2,0]]
	// k := int(3)
	// rowConditions := [][]int{{1,2},{3,2}}
	// colConditions := [][]int{{2,1},{3,2}}

	// result: []
	k := int(3)
	rowConditions := [][]int{{1,2},{2,3},{3,1},{2,3}}
	colConditions := [][]int{{2,1}}

	// result: []
	// k := int(0)
	// rowConditions := [][]int{}
	// colConditions := [][]int{}

	result := buildMatrix(k, rowConditions, colConditions)
	fmt.Printf("result = %v\n", result)
}

