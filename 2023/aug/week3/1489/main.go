package main
import (
	"fmt"
	"sort"
)

type UnionFind struct {
    Parents []int
	Sizes []int
	MaxSize int
}

func Constructor(size int) *UnionFind {
	parents := make([]int, size)
	sizes := make([]int, size)
	maxSize := int(1)

	for i := 0; i < size; i++ {
		parents[i] = i
		sizes[i] = 1
	}

	return &UnionFind{
		Parents: parents,
		Sizes: sizes,
		MaxSize: maxSize,
	}
}

func (this *UnionFind) Find(node int) int {
	if node != this.Parents[node] {
		this.Parents[node] = this.Find(this.Parents[node])
	}
	return this.Parents[node]
}

func (this *UnionFind) Union(node1 int, node2 int) bool {
	rootNode1 := this.Find(node1)
	rootNode2 := this.Find(node2)

	if rootNode1 != rootNode2 {
		if this.Sizes[rootNode1] < this.Sizes[rootNode2] {
			temp := rootNode1
			rootNode1 = rootNode2
			rootNode2 = temp
		}

		this.Parents[rootNode2] = rootNode1
		this.Sizes[rootNode1] += this.Sizes[rootNode2]
		this.MaxSize = max(this.MaxSize, this.Sizes[rootNode1])
		return true
	}

	return false
}

func findCriticalAndPseudoCriticalEdges(n int, edges [][]int) [][]int {
	m := len(edges)
	newEdges := make([][]int, m)
	for i, _ := range newEdges {
		newEdges[i] = make([]int, 4)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < 3; j++ {
			newEdges[i][j] = edges[i][j]
		}
		newEdges[i][3] = i
	}

	sort.Slice(newEdges, func(a, b int) bool {
		return newEdges[a][2] < newEdges[b][2]
	})

	uf := Constructor(n)
	stdWeight := int(0)

	for _, edge := range newEdges {
		if uf.Union(edge[0], edge[1]) {
			stdWeight += edge[2]
		}
	}

	result := make([][]int, 0)
	for i := 0; i < 2; i++ {
		result = append(result, make([]int, 0))
	}

	for i := 0; i < m; i++ {
		ufIgnore := Constructor(n)
		ignoreWeight := int(0)

		for j := 0; j < m; j++ {
			if i != j && ufIgnore.Union(newEdges[j][0], newEdges[j][1]) {
				ignoreWeight += newEdges[j][2]
			}
		}

		if ufIgnore.MaxSize < n || ignoreWeight > stdWeight {
			result[0] = append(result[0], newEdges[i][3])
		} else {
			ufForce := Constructor(n)
			ufForce.Union(newEdges[i][0], newEdges[i][1])
			forceWeight := newEdges[i][2]

			for j := 0; j < m; j++ {
				if i != j && ufForce.Union(newEdges[j][0], newEdges[j][1]) {
					forceWeight += newEdges[j][2]
				}
			}

			if forceWeight == stdWeight {
				result[1] = append(result[1], newEdges[i][3])
			}
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
	// result: [[0,1],[2,3,4,5]]
	// n := int(5)
	// edges := [][]int{{0,1,1},{1,2,1},{2,3,2},{0,3,2},{0,4,3},{3,4,3},{1,4,6}}

	// result: [[],[0,1,2,3]]
	n := int(4)
	edges := [][]int{{0,1,1},{1,2,1},{2,3,1},{0,3,1}}

	// result: 
	// n := int(0)
	// edges := [][]int{}

	result := findCriticalAndPseudoCriticalEdges(n, edges)
	fmt.Printf("result = %v\n", result)
}

