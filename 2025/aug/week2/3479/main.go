package main

import (
	"fmt"
	"math"
)

type Tree struct {
	Nodes []int
	Baskets []int
}

func (this *Tree) Build(p, l, r int) {
	if l == r {
		this.Nodes[p] = this.Baskets[l]
		return
	}

	mid := (l + r) >> 1
	this.Build(p << 1, l, mid)
	this.Build(p << 1 | 1, mid + 1, r)
	this.Nodes[p] = max(this.Nodes[p << 1], this.Nodes[p << 1 | 1])
}

func (this *Tree) Query(p, l, r, ql, qr int) int {
	if ql > r || qr < l {
		return math.MinInt32
	}

	if ql <= l && r <= qr {
		return this.Nodes[p]
	}

	mid := (l + r) >> 1
	return max(this.Query(p << 1, l, mid, ql, qr), this.Query(p << 1 | 1, mid + 1, r, ql, qr))
}

func (this *Tree) Update(p, l, r, pos, val int) {
	if l == r {
		this.Nodes[p] = val
		return
	}

	mid := (l + r) >> 1
	if pos <= mid {
		this.Update(p << 1, l, mid, pos, val)
	} else {
		this.Update(p << 1 | 1, mid + 1, r, pos, val)
	}

	this.Nodes[p] = max(this.Nodes[p << 1], this.Nodes[p << 1 | 1])
}

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	m := len(baskets)
	if m == 0 {
		return len(fruits)
	}

	tree := Tree{
		Nodes: make([]int, 4 * m + 7),
		Baskets: baskets,
	}

	tree.Build(1, 0, m - 1)

	result := int(0)
	for i := 0; i < len(fruits); i++ {
		l := int(0)
		r := m - 1
		temp := int(-1)

		for l <= r {
			mid := (l + r) >> 1

			if tree.Query(1, 0, m - 1, 0, mid) >= fruits[i] {
				temp = mid
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

		if temp != -1 && tree.Baskets[temp] >= fruits[i] {
			tree.Update(1, 0, m - 1, temp, math.MinInt32)
		} else {
			result++
		}
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 1
	// fruits := []int{4,2,5}
	// baskets := []int{3,5,4}

	// result: 0
	fruits := []int{3,6,1}
	baskets := []int{6,4,7}

	// result: 
	// fruits := []int{}
	// baskets := []int{}

	result := numOfUnplacedFruits(fruits, baskets)
	fmt.Printf("result = %v\n", result)
}

