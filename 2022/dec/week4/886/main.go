package main
import (
	"fmt"
)

type UnionFind struct {
	N    int
	Root []int
}

func newUnionFind(n int) UnionFind {
	root := make([]int, n)
	for i := 0; i < n; i++ {
		root[i] = -1
	}
	uf := UnionFind{N: n, Root: root}
	return uf
}

func (this *UnionFind) Find(x int) int {
	if this.Root[x] < 0 {
		return x
	}
	this.Root[x] = this.Find(this.Root[x])
	return this.Root[x]
}

func (this *UnionFind) Union(x, y int) {
	x = this.Find(x)
	y = this.Find(y)
	if x == y {
		return
	}
	if -this.Root[x] < -this.Root[y] {
		x, y = y, x
	}
	this.Root[x] += this.Root[y]
	this.Root[y] = x
}

func (this *UnionFind) IsSame(x, y int) bool {
	return this.Find(x) == this.Find(y)
}

func (this *UnionFind) Size(x int) int {
	return -this.Root[this.Find(x)]
}

func possibleBipartition(n int, dislikes [][]int) bool {
	if len(dislikes) <= 1 {
		return true
	}

	unionFind := newUnionFind(n * n)
	for _, dislike := range dislikes {
      unionFind.Union(dislike[0], dislike[1] + n)
      unionFind.Union(dislike[1], dislike[0] + n)
    }
	for i := 1; i <= n; i++ {
      if unionFind.IsSame(i, i + n) {
        return false
      }
    }
	return true
}

func main() {
	// result: true
	// n := int(4)
	// dislikes := [][]int{{1,2},{1,3},{2,4}}

	// result: false
	// n := int(3)
	// dislikes := [][]int{{1,2},{1,3},{2,3}}

	// result: false
	n := int(5)
	dislikes := [][]int{{1,2},{2,3},{3,4},{4,5},{1,5}}

	// result: 
	// n := int(0)
	// dislikes := [][]int{}

	result := possibleBipartition(n, dislikes)
	fmt.Printf("result = %v\n", result)
}

