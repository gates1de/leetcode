package main
import (
	"fmt"
	"math"
)

type Graph struct {
	Matrix [][]int
}

func Constructor(n int, edges [][]int) Graph {
	matrix := make([][]int, n)
	for i, _ := range matrix {
		matrix[i] = make([]int, n)
		for j, _ := range matrix[i] {
			matrix[i][j] = math.MaxInt32
		}
	}

	for _, edge := range edges {
		matrix[edge[0]][edge[1]] = edge[2]
	}
	for i, _ := range matrix {
		matrix[i][i] = 0
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				matrix[j][k] = min(matrix[j][k], matrix[j][i] + matrix[i][k])
			}
		}
	}

	return Graph{Matrix: matrix}
}

func (this *Graph) AddEdge(edge []int)  {
	n := len(this.Matrix)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			cost := this.Matrix[i][edge[0]] + this.Matrix[edge[1]][j] + edge[2]
			this.Matrix[i][j] = min(this.Matrix[i][j], cost)
		}
	}
}

func (this *Graph) ShortestPath(node1 int, node2 int) int {
	if this.Matrix[node1][node2] == math.MaxInt32 {
		return -1
	}
	return this.Matrix[node1][node2]
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: [null, 6, -1, null, 6]
	n := int(4)
	edges := [][]int{{0,2,5}, {0,1,2}, {1,2,1}, {3,0,3}}
	operations := [][]int{{1,3,2},{1,0,3},{0,1,3,4},{1,0,3}}

	// result: 
	// n := int(0)
	// edges := [][]int{}
	// operations := [][]int{{0,1},{1,0,0}}

	obj := Constructor(n, edges)
	for _, ope := range operations {
		if ope[0] == 0 {
			obj.AddEdge(ope[1:])
		} else if ope[0] == 1 {
			fmt.Printf("obj.ShortestPath(%v, %v) = %v\n", ope[1], ope[2], obj.ShortestPath(ope[1], ope[2]))
		}
	}
}

