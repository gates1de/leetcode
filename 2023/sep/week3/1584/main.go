package main
import (
	"fmt"
	"sort"
)

type UnionFind struct {
    Parents []int
    Sizes []int
}

func Constructor(n int) UnionFind {
    uf := UnionFind{Parents: make([]int, n), Sizes: make([]int, n)}
    for i := 0; i < n; i++ {
        uf.Parents[i] = i
    }
    return uf
}

func (this *UnionFind) Find(x int) int {
    if x == this.Parents[x] {
        return x
    }

    this.Parents[x] = this.Find(this.Parents[x])
    return this.Parents[x]
}

func (this *UnionFind) Unite(x int, y int) {
    rx := this.Find(x)
    ry := this.Find(y)

    if rx == ry {
        return
    }

    if this.Sizes[rx] < this.Sizes[ry] {
        rx, ry = ry, rx
    }

    this.Parents[ry] = rx
    this.Sizes[rx] += this.Sizes[ry]
}

func (this *UnionFind) Same(x int, y int) bool {
    return this.Find(x) == this.Find(y)
}

func (this *UnionFind) Size(x int) int {
    return this.Sizes[this.Find(x)]
}

type Edge struct{
    A int
    B int
    Cost int
}

func minCostConnectPoints(points [][]int) int {
    result := int(0)
    edges := []Edge{}
    for i, p1 := range points {
        for j, p2 := range points {
            if i == j {
                continue
            }
            cost := abs(p1[0], p2[0]) + abs(p1[1], p2[1])
            edges = append(edges, Edge{A: i, B: j, Cost: cost})
        }
    }

    sort.Slice(edges, func(a, b int) bool {
        return edges[a].Cost < edges[b].Cost
    })

    uf := Constructor(len(points))
    for i := 0; i < len(edges); i++ {
        edge := edges[i]
        if !uf.Same(edge.A, edge.B) {
            result += edge.Cost
            uf.Unite(edge.A, edge.B)
        }
    }

    return result
}

func abs(a int, b int) int {
    if a < b {
        return b - a
    }
    return a - b
}

func main() {
	// result: 20
	// points := [][]int{{0,0},{2,2},{3,10},{5,2},{7,0}}

	// result: 18
	points := [][]int{{3,12},{-2,5},{-4,1}}

	// result: 
	// points := [][]int{}

	result := minCostConnectPoints(points)
	fmt.Printf("result = %v\n", result)
}

