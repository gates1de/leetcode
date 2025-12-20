package main

import (
	"fmt"
	"sort"
)

type UnionFind struct {
	N int
	P []int
}

func (this UnionFind) Init() {
	for i := 0; i < len(this.P); i++ {
		this.P[i] = i
	}
}
func (this UnionFind) Union(x, y int) {
	px := this.Find(x)
	py := this.Find(y)
	if px != py {
		this.P[px] = py
	}
}

func (this UnionFind) Find(x int) int {
	tx := x
	for x != this.P[x] {
		x = this.P[x]
	}

	for tx != x {
		tmp := this.P[tx]
		this.P[tx] = x
		tx = tmp
	}

	return x
}

func (this UnionFind) Unset(x int) {
	this.P[x] = x
}

type ByThird [][]int

func (b ByThird) Len() int {
	return len(b)
}

func (b ByThird) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByThird) Less(i, j int) bool {
	return b[i][2] < b[j][2]
}

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	this := UnionFind{n, make([]int, n)}
	this.Init()
	this.Union(0, firstPerson)
	sort.Sort(ByThird(meetings))

	i := int(0)
	for i < len(meetings) {
		time := meetings[i][2]
		line := make([]int, 0)
		for i < len(meetings) && meetings[i][2] == time {
			line = append(line, meetings[i][0])
			line = append(line, meetings[i][1])
			this.Union(meetings[i][0], meetings[i][1])
			i += 1
		}

		root := this.Find(0)
		for _, e := range line {
			if this.Find(e) != root {
				this.Unset(e)
			}
		}
	}

	root := this.Find(0)
	result := make([]int, 0)
	for i := 0; i < n; i++ {
		if this.Find(i) == root {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	// result: [0,1,2,3,5]
	// n := int(6)
	// meetings := [][]int{{1,2,5},{2,3,8},{1,5,10}}
	// firstPerson := int(1)

	// result: [0,1,3]
	// n := int(4)
	// meetings := [][]int{{3,1,3},{1,2,2},{0,3,3}}
	// firstPerson := int(3)

	// result: [0,1,2,3,4]
	n := int(5)
	meetings := [][]int{{3,4,2},{1,2,1},{2,3,1}}
	firstPerson := int(1)

	// result: 
	// n := int(0)
	// meetings := [][]int{}
	// firstPerson := int(0)

	result := findAllPeople(n, meetings, firstPerson)
	fmt.Printf("result = %v\n", result)
}

