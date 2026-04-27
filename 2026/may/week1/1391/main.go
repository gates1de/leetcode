package main

import (
	"fmt"
)

type DisjointSet struct {
	F []int
}

func (this *DisjointSet) Init(m, n int) {
	this.F = make([]int, m * n)
	for i := range this.F {
		this.F[i] = i
	}
}

func (this *DisjointSet) Find(x int) int {
	if this.F[x] == x {
		return x
	}
	return this.Find(this.F[x])
}

func (this *DisjointSet) Merge(x, y int) {
	this.F[this.Find(x)] = this.Find(y)
}

func (this *DisjointSet) GetId(x, y, n int) int {
	return x * n + y
}

func (this *DisjointSet) DetectL(grid [][]int, x, y, m, n int) {
	if y - 1 >= 0 && (grid[x][y - 1] == 4 || grid[x][y - 1] == 6 || grid[x][y - 1] == 1) {
		this.Merge(this.GetId(x, y, n), this.GetId(x, y - 1, n))
	}
}

func (this *DisjointSet) DetectR(grid [][]int, x, y, m, n int) {
	if y + 1 < n && (grid[x][y + 1] == 3 || grid[x][y + 1] == 5 || grid[x][y + 1] == 1) {
		this.Merge(this.GetId(x, y, n), this.GetId(x, y + 1, n));
	}
}

func (this *DisjointSet) DetectU(grid [][]int, x, y, m, n int) {
	if x - 1 >= 0 && (grid[x - 1][y] == 3 || grid[x - 1][y] == 4 || grid[x - 1][y] == 2) {
		this.Merge(this.GetId(x, y, n), this.GetId(x - 1, y, n));
	}
}

func (this *DisjointSet) DetectD(grid [][]int, x, y, m, n int) {
	if x + 1 < m && (grid[x + 1][y] == 5 || grid[x + 1][y] == 6 || grid[x + 1][y] == 2) {
		this.Merge(this.GetId(x, y, n), this.GetId(x + 1, y, n))
	}
}

func (this *DisjointSet) Handler(grid [][]int, x, y, m, n int) {
	switch grid[x][y] {
	case 1:
		this.DetectL(grid, x, y, m, n)
		this.DetectR(grid, x, y, m, n)
	case 2:
		this.DetectU(grid, x, y, m, n)
		this.DetectD(grid, x, y, m, n)
	case 3:
		this.DetectL(grid, x, y, m, n)
		this.DetectD(grid, x, y, m, n)
	case 4:
		this.DetectR(grid, x, y, m, n)
		this.DetectD(grid, x, y, m, n)
	case 5:
		this.DetectL(grid, x, y, m, n)
		this.DetectU(grid, x, y, m, n)
	case 6:
		this.DetectR(grid, x, y, m, n)
		this.DetectU(grid, x, y, m, n)
	}
}

func hasValidPath(grid [][]int) bool {
	m := len(grid)
	n := len(grid[0])
	ds := &DisjointSet{}
	ds.Init(len(grid), len(grid[0]))
	for i := range m {
		for j := range n {
			ds.Handler(grid, i, j, m, n)
		}
	}

	return ds.Find(ds.GetId(0, 0, n)) == ds.Find(ds.GetId(m - 1, n - 1, n))
}

func main() {
	// result: true
	// grid := [][]int{{2,4,3},{6,5,2}}

	// result: false
	// grid := [][]int{{1,2,1},{1,2,1}}

	// result: false
	grid := [][]int{{1,1,2}}

	// result: 
	// grid := [][]int{}

	result := hasValidPath(grid)
	fmt.Printf("result = %v\n", result)
}

