package main

import (
	"fmt"
	"sort"
)

const MAX_STABILITY = 200000

type DSU struct {
	Parents []int
}

func NewDSU(parents []int) *DSU {
	p := make([]int, len(parents))
	copy(p, parents)
	return &DSU{Parents: p}
}

func (d *DSU) Find(x int) int {
	if d.Parents[x] != x {
		d.Parents[x] = d.Find(d.Parents[x])
	}
	return d.Parents[x]
}

func (d *DSU) Join(x, y int) {
	px := d.Find(x)
	py := d.Find(y)
	d.Parents[px] = py
}

func maxStability(n int, edges [][]int, k int) int {
	result := int(-1)
	if len(edges) < n-1 {
		return result
	}

	mustEdges := make([][]int, 0)
	optionalEdges := make([][]int, 0)

	for _, e := range edges {
		if e[3] == 1 {
			mustEdges = append(mustEdges, e)
		} else {
			optionalEdges = append(optionalEdges, e)
		}
	}

	if len(mustEdges) > n - 1 {
		return result
	}

	sort.Slice(optionalEdges, func(i, j int) bool {
		return optionalEdges[i][2] > optionalEdges[j][2]
	})

	selectedInit := int(0)
	mustMinStability := MAX_STABILITY

	initParent := make([]int, n)
	for i := range n {
		initParent[i] = i
	}
	dsuInit := NewDSU(initParent)

	for _, e := range mustEdges {
		u, v, s := e[0], e[1], e[2]
		if dsuInit.Find(u) == dsuInit.Find(v) || selectedInit == n-1 {
			return result
		}

		dsuInit.Join(u, v)
		selectedInit++
		if s < mustMinStability {
			mustMinStability = s
		}
	}

	l := int(0)
	r := mustMinStability
	for l < r {
		mid := l + (r - l + 1) / 2

		dsu := NewDSU(dsuInit.Parents)
		selected := selectedInit
		doubledCount := int(0)

		for _, e := range optionalEdges {
			u, v, s := e[0], e[1], e[2]
			if dsu.Find(u) == dsu.Find(v) {
				continue
			}

			if s >= mid {
				dsu.Join(u, v)
				selected++
			} else if doubledCount < k && s * 2 >= mid {
				doubledCount++
				dsu.Join(u, v)
				selected++
			} else {
				break
			}

			if selected == n-1 {
				break
			}
		}

		if selected != n-1 {
			r = mid - 1
		} else {
			result = mid
			l = mid
		}
	}

	return result
}

func main() {
	// result: 2
	// n := int(3)
	// edges := [][]int{{0,1,2,1},{1,2,3,0}}
	// k := int(1)

	// result: 6
	// n := int(4)
	// edges := [][]int{{0,1,4,0},{1,2,3,0},{0,2,1,0}}
	// k := int(2)

	// result: -1
	n := int(3)
	edges := [][]int{{0,1,1,1},{1,2,1,1},{2,0,1,1}}
	k := int(0)

	// result: 
	// n := int(0)
	// edges := [][]int{}
	// k := int(0)

	result := maxStability(n, edges, k)
	fmt.Printf("result = %v\n", result)
}

