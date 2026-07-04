package main

import (
	"fmt"
	"sort"
)

type Edge struct {
	To   int
	Cost int
}

func findMaxPathScore(edges [][]int, online []bool, k int64) int {
	n := len(online)

	adj := make([][]Edge, n)
	indegree := make([]int, n)
	costs := make([]int, 0, len(edges))
	for _, e := range edges {
		u := e[0]
		v := e[1]
		c := e[2]
		adj[u] = append(adj[u], Edge{To: v, Cost: c})
		indegree[v]++
		costs = append(costs, c)
	}

	queue := make([]int, 0, n)
	for i := range n {
		if indegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	topo := make([]int, 0, n)
	for head := 0; head < len(queue); head++ {
		u := queue[head]
		topo = append(topo, u)
		for _, e := range adj[u] {
			indegree[e.To]--
			if indegree[e.To] == 0 {
				queue = append(queue, e.To)
			}
		}
	}

	sort.Ints(costs)
	uniq := make([]int, 0, len(costs))
	for _, c := range costs {
		if len(uniq) == 0 || uniq[len(uniq) - 1] != c {
			uniq = append(uniq, c)
		}
	}

	check := func(limit int) bool {
		const inf = int64(1<<62)
		dist := make([]int64, n)
		for i := range n {
			dist[i] = inf
		}
		dist[0] = 0

		for _, u := range topo {
			if !online[u] || dist[u] == inf {
				continue
			}

			for _, e := range adj[u] {
				if e.Cost < limit {
					continue
				}

				if !online[e.To] {
					continue
				}

				nextCost := dist[u] + int64(e.Cost)
				if nextCost < dist[e.To] {
					dist[e.To] = nextCost
				}
			}
		}

		return dist[n-1] <= k
	}

	result := -1
	left := 0
	right := len(uniq) - 1
	for left <= right {
		mid := (left + right) / 2
		if check(uniq[mid]) {
			result = uniq[mid]
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result
}

func main() {
	// result: 3
	// edges := [][]int{{0,1,5},{1,3,10},{0,2,3},{2,3,4}}
	// online := []bool{true,true,true,true}
	// k := int64(0)

	// result: 6
	edges := [][]int{{0,1,7},{1,4,5},{0,2,6},{2,3,6},{3,4,2},{2,4,6}}
	online := []bool{true,true,true,false,true}
	k := int64(12)

	// result: 
	// edges := [][]int{}
	// online := []bool{}
	// k := int64(0)

	result := findMaxPathScore(edges, online, k)
	fmt.Printf("result = %v\n", result)
}
