package main

import (
	"fmt"
	"math"

	"github.com/emirpasic/gods/trees/redblacktree"
)

func minOperations(s string, k int) int {
	n := len(s)
	m := int(0)
	dist := make([]int, n + 1)
	for i := range dist {
		dist[i] = math.MaxInt32
	}

	nodeSets := make([]*redblacktree.Tree, 2)
	nodeSets[0] = redblacktree.NewWithIntComparator()
	nodeSets[1] = redblacktree.NewWithIntComparator()
	for i := 0; i <= n; i++ {
		nodeSets[i % 2].Put(i, struct{}{})
		if i < n && s[i] == '0' {
			m++
		}
	}

	q := []int{m}
	dist[m] = 0
	nodeSets[m % 2].Remove(m)

	for len(q) > 0 {
		m := q[0]
		q = q[1:]
		c1 := max(k - n + m, 0)
		c2 := min(k, m)
		lnode := m + k - 2 * c2
		rnode := m + k - 2 * c1
		nodeSet := nodeSets[lnode % 2]

		for m2, found := nodeSet.Ceiling(lnode); found && m2.Key.(int) <= rnode; m2, found = nodeSet.Ceiling(lnode) {
			dist[m2.Key.(int)] = dist[m] + 1
            q = append(q, m2.Key.(int))
			nodeSet.Remove(m2.Key.(int))
		}
	}

	if dist[0] == math.MaxInt32 {
		return -1
	}

	return dist[0]
}

func main() {
	// result: 1
	// s := "110"
	// k := int(1)

	// result: 2
	// s := "0101"
	// k := int(3)

	// result: -1
	s := "101"
	k := int(2)

	// result: 
	// s := ""
	// k := int(0)

	result := minOperations(s, k)
	fmt.Printf("result = %v\n", result)
}

