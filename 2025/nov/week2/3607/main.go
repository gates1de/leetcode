package main

import (
	"fmt"
)

type DSU struct {
	Parents []int
}

func NewDSU(size int) *DSU {
	parents := make([]int, size)
	for i := range parents {
		parents[i] = i
	}
	return &DSU{Parents: parents}
}

func (this *DSU) Find(x int) int {
	if this.Parents[x] == x {
		return x
	}

	this.Parents[x] = this.Find(this.Parents[x])
	return this.Parents[x]
}

func (this *DSU) Join(u, v int) {
	this.Parents[this.Find(v)] = this.Find(u)
}

func processQueries(c int, connections [][]int, queries [][]int) []int {
	dsu := NewDSU(c + 1)
	for _, p := range connections {
		dsu.Join(p[0], p[1])
	}

	online := make([]bool, c + 1)
	offlineCounts := make([]int, c + 1)
	for i := range online {
		online[i] = true
	}

	minimumOnlineStations := make(map[int]int)
	for _, q := range queries {
		op := q[0]
		x := q[1]
		if op == 2 {
			online[x] = false
			offlineCounts[x]++
		}
	}

	for i := 1; i <= c; i++ {
		root := dsu.Find(i)
		if _, exists := minimumOnlineStations[root]; !exists {
			minimumOnlineStations[root] = -1
		}

		station := minimumOnlineStations[root]
		if online[i] {
			if station == -1 || station > i {
				minimumOnlineStations[root] = i
			}
		}
	}

	result := make([]int, 0, len(queries))
	for i := len(queries) - 1; i >= 0; i-- {
		op, x := queries[i][0], queries[i][1]
		root := dsu.Find(x)
		station := minimumOnlineStations[root]

		if op == 1 {
			if online[x] {
				result = append(result, x)
			} else {
				result = append(result, station)
			}
		}

		if op == 2 {
			if offlineCounts[x] > 1 {
				offlineCounts[x]--
			} else {
				online[x] = true
				if station == -1 || station > x {
					minimumOnlineStations[root] = x
				}
			}
		}
	}

	for i, j := 0, len(result) - 1; i < j; i, j = i + 1, j - 1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func main() {
	// result: [3,2,3]
	// c := int(5)
	// connections := [][]int{{1,2},{2,3},{3,4},{4,5}}
	// queries := [][]int{{1,3},{2,1},{1,1},{2,2},{1,2}}

	// result: [1,-1]
	c := int(3)
	connections := [][]int{}
	queries := [][]int{{1,1},{2,1},{1,1}}

	// result: []
	// c := int(0)
	// connections := [][]int{}
	// queries := [][]int{}

	result := processQueries(c, connections, queries)
	fmt.Printf("result = %v\n", result)
}

