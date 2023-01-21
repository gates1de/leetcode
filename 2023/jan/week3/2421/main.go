package main
import (
	"fmt"
	"sort"
)


type UnionFind struct {
	Parents []int
}

func Constructor(n int) *UnionFind {
	par := make([]int, n)
	for i := 0; i < n; i++ {
		par[i] = -1
	}
	return &UnionFind{
		Parents: par,
	}
}

func (this *UnionFind) Find(x int) int {
	if this.Parents[x] >= 0 {
		this.Parents[x] = this.Find(this.Parents[x])
		return this.Parents[x]
	}

	return x
}

func (this *UnionFind) Union(x, y int) bool {
	parentX := this.Find(x)
	parentY := this.Find(y)

	if parentX == parentY {
		return false
	}

	if this.Parents[parentX] <= this.Parents[parentY] {
		this.Parents[parentX] += this.Parents[parentY]
		this.Parents[parentY] = parentX
	} else {
		this.Parents[parentY] += this.Parents[parentX]
		this.Parents[parentX] = parentY
	}

	return true
}

func numberOfGoodPaths(vals []int, edges [][]int) int {
	adjacents:= make(map[int][]int)
	mp := make(map[int][]int)
	n := len(vals)

	for i := range vals {
		mp[vals[i]] = append(mp[vals[i]], i)
	}

	for _, e := range edges {
		u, v := e[0], e[1]

		adjacents[u] = append(adjacents[u], v)
		adjacents[v] = append(adjacents[v], u)
	}

	var keys []int
	for k, _ := range mp {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	uf := Constructor(n)
	result := n

	for _, k := range keys {
		for _, curr := range mp[k] {
			for _, adj := range adjacents[curr] {
				if vals[adj] <= vals[curr] {
					uf.Union(curr, adj)
				}
			}
		}
		
		group := make(map[int]int) 
		for _, curr := range mp[k] {
			group[uf.Find(curr)]++
		}
		
		for _, v := range group {
			if v >= 2 {
				result += v * (v - 1) / 2
			}
		}
	}
	
	return result
}

func main() {
	// result: 6
	// vals := []int{1,3,2,1,3}
	// edges := [][]int{{0,1},{0,2},{2,3},{2,4}}

	// result: 7
	// vals := []int{1,1,2,2,3}
	// edges := [][]int{{0,1},{1,2},{2,3},{2,4}}

	// result: 1
	vals := []int{1}
	edges := [][]int{}

	// result: 
	// vals := []int{}
	// edges := [][]int{}

	result := numberOfGoodPaths(vals, edges)
	fmt.Printf("result = %v\n", result)
}

