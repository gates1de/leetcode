package main
import (
	"fmt"
	"sort"
)

type UnionFind struct {
	Group []int
	Rank []int
}

func constructor(size int) UnionFind {
	group := make([]int, size)
	rank := make([]int, size)
	for i, _ := range group {
		group[i] = i
	}
	return UnionFind{Group: group, Rank: rank}
}

func (this *UnionFind) Find(node int) int {
	if this.Group[node] != node {
		this.Group[node] = this.Find(this.Group[node])
	}
	return this.Group[node]
}


func (this *UnionFind) Join(node1, node2 int) {
	group1 := this.Find(node1)
	group2 := this.Find(node2)

	if group1 == group2 {
		return
	}

	if this.Rank[group1] > this.Rank[group2] {
		this.Group[group2] = group1
	} else if this.Rank[group1] < this.Rank[group2] {
		this.Group[group1] = group2
	} else {
		this.Group[group1] = group2
		this.Rank[group2]++
	}
}

func (this *UnionFind) AreConnected(node1, node2 int) bool {
	return this.Group[this.Find(node1)] == this.Group[this.Find(node2)]
}

func distanceLimitedPathsExist(n int, edgeList [][]int, queries [][]int) []bool {
	uf := constructor(n)
	queriesCount := len(queries)
	result := make([]bool, queriesCount)

	queriesWithIndex := make([][]int, queriesCount)
	for i := 0; i < queriesCount; i++ {
		queriesWithIndex[i] = make([]int, 4)
		queriesWithIndex[i][0] = queries[i][0]
		queriesWithIndex[i][1] = queries[i][1]
		queriesWithIndex[i][2] = queries[i][2]
		queriesWithIndex[i][3] = i
	}

	sort.Slice(edgeList, func(a, b int) bool {
		return edgeList[a][2] < edgeList[b][2]
	})
	sort.Slice(queriesWithIndex, func(a, b int) bool {
		return queriesWithIndex[a][2] < queriesWithIndex[b][2]
	})
	edgesIndex := int(0)

	for queryIndex := 0; queryIndex < queriesCount; queryIndex++ {
		p := queriesWithIndex[queryIndex][0]
		q := queriesWithIndex[queryIndex][1]
		limit := queriesWithIndex[queryIndex][2]
		queryOriginalIndex := queriesWithIndex[queryIndex][3]

		for edgesIndex < len(edgeList) && edgeList[edgesIndex][2] < limit {
			node1 := edgeList[edgesIndex][0]
			node2 := edgeList[edgesIndex][1]
			uf.Join(node1, node2)
			edgesIndex++
		}

		result[queryOriginalIndex] = uf.AreConnected(p, q)
	}

	return result
}

func main() {
	// result: [false,true]
	// n := int(3)
	// edgeList := [][]int{{0,1,2},{1,2,4},{2,0,8},{1,0,16}}
	// queries := [][]int{{0,1,2},{0,2,5}}

	// result: [true,false]
	n := int(5)
	edgeList := [][]int{{0,1,10},{1,2,5},{2,3,9},{3,4,13}}
	queries := [][]int{{0,4,14},{1,4,13}}

	// result: 
	// n := int(0)
	// edgeList := [][]int{}
	// queries := [][]int{}

	result := distanceLimitedPathsExist(n, edgeList, queries)
	fmt.Printf("result = %v\n", result)
}

