package main

import (
	"fmt"
)

func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)
	dist := make([][]int, n)
	q := make([]int, 0, n*n)
	hasThief := false
	for i := 0; i < n; i++ {
		dist[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dist[i][j] = -1
			if grid[i][j] == 1 {
				hasThief = true
				dist[i][j] = 0
				q = append(q, i*n+j)
			}
		}
	}

	if !hasThief {
		return 0
	}

	dirs := []int{0, 1, 0, -1, 0}
	for head := 0; head < len(q); head++ {
		idx := q[head]
		i := idx / n
		j := idx % n
		for k := 0; k < 4; k++ {
			x := i + dirs[k]
			y := j + dirs[k+1]
			if x >= 0 && x < n && y >= 0 && y < n && dist[x][y] == -1 {
				dist[x][y] = dist[i][j] + 1
				q = append(q, x*n+y)
			}
		}
	}

	maxDist := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if dist[i][j] > maxDist {
				maxDist = dist[i][j]
			}
		}
	}

	buckets := make([][]int, maxDist+1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			d := dist[i][j]
			buckets[d] = append(buckets[d], i*n+j)
		}
	}

	parent := make([]int, n*n)
	rank := make([]int, n*n)
	active := make([]bool, n*n)
	for i := range parent {
		parent[i] = i
	}

	var find func(int) int
	find = func(x int) int {
		if parent[x] != x {
			parent[x] = find(parent[x])
		}
		return parent[x]
	}

	union := func(a, b int) {
		ra := find(a)
		rb := find(b)
		if ra == rb {
			return
		}
		if rank[ra] < rank[rb] {
			ra, rb = rb, ra
		}
		parent[rb] = ra
		if rank[ra] == rank[rb] {
			rank[ra]++
		}
	}

	dirs2 := []int{0, 1, 0, -1, 0}
	start := 0
	end := n*n - 1
	for d := maxDist; d >= 0; d-- {
		for _, idx := range buckets[d] {
			active[idx] = true
			i := idx / n
			j := idx % n
			for k := 0; k < 4; k++ {
				x := i + dirs2[k]
				y := j + dirs2[k+1]
				if x >= 0 && x < n && y >= 0 && y < n {
					ni := x*n + y
					if active[ni] {
						union(idx, ni)
					}
				}
			}
			if active[start] && active[end] && find(start) == find(end) {
				return d
			}
		}
	}

	return 0
}

func main() {
	// result: 0
	// grid := [][]int{{1,0,0},{0,0,0},{0,0,1}}

	// result: 2
	// grid := [][]int{{0,0,1},{0,0,0},{0,0,0}}

	// result: 2
	// grid := [][]int{{0,0,0,1},{0,0,0,0},{0,0,0,0},{1,0,0,0}}

	// result: 0
	grid := [][]int{{1,0,0},{0,0,0},{0,0,1}}

	// result: 
	// grid := [][]int{}

	result := maximumSafenessFactor(grid)
	fmt.Printf("result = %v\n", result)
}
