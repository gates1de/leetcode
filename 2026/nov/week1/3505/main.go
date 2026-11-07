package main

import (
	"fmt"
	"sort"
)

func minOperations(nums []int, x int, k int) int64 {
	n := len(nums)
	if n == 0 || x <= 0 || k <= 0 {
		return 0
	}

	values := append([]int(nil), nums...)
	sort.Ints(values)
	unique := values[:1]
	for _, value := range values[1:] {
		if unique[len(unique)-1] != value {
			unique = append(unique, value)
		}
	}

	ranks := make([]int, n)
	for i, value := range nums {
		ranks[i] = sort.SearchInts(unique, value) + 1
	}

	countTree := newFenwick(len(unique))
	sumTree := newFenwick(len(unique))
	for i := range ranks[:x] {
		countTree.add(ranks[i], 1)
		sumTree.add(ranks[i], int64(nums[i]))
	}

	windowCosts := make([]int64, n-x+1)
	for start := range windowCosts {
		medianRank := countTree.kth(int64((x + 1) / 2))
		median := int64(unique[medianRank-1])
		leftCount := countTree.query(medianRank)
		leftSum := sumTree.query(medianRank)
		totalSum := sumTree.query(len(unique))
		windowCosts[start] = median*leftCount - leftSum + totalSum - leftSum - median*(int64(x)-leftCount)

		if start+1 < len(windowCosts) {
			countTree.add(ranks[start], -1)
			sumTree.add(ranks[start], -int64(nums[start]))
			countTree.add(ranks[start+x], 1)
			sumTree.add(ranks[start+x], int64(nums[start+x]))
		}
	}

	const inf int64 = 1 << 62
	previous := make([]int64, n+1)
	current := make([]int64, n+1)
	for group := 1; group <= k; group++ {
		for i := range current {
			current[i] = inf
		}

		for length := 1; length <= n; length++ {
			current[length] = current[length-1]
			if length < x || previous[length-x] == inf {
				continue
			}

			candidate := previous[length-x] + windowCosts[length-x]
			if candidate < current[length] {
				current[length] = candidate
			}
		}

		previous, current = current, previous
	}

	return previous[n]
}

type fenwick struct {
	tree []int64
}

func newFenwick(size int) fenwick {
	return fenwick{tree: make([]int64, size+1)}
}

func (f *fenwick) add(index int, value int64) {
	for index < len(f.tree) {
		f.tree[index] += value
		index += index & -index
	}
}

func (f *fenwick) query(index int) int64 {
	result := int64(0)
	for index > 0 {
		result += f.tree[index]
		index -= index & -index
	}
	return result
}

func (f *fenwick) kth(target int64) int {
	index := 0
	step := 1
	for step<<1 < len(f.tree) {
		step <<= 1
	}

	for step > 0 {
		next := index + step
		if next < len(f.tree) && f.tree[next] < target {
			target -= f.tree[next]
			index = next
		}
		step >>= 1
	}

	return index + 1
}

func main() {
	// result: 8
	// nums := []int{5,-2,1,3,7,3,6,4,-1}
	// x := 3
	// k := 2

	// result: 3
	nums := []int{9, -2, -2, -2, 1, 5}
	x := 2
	k := 2

	// result: 0
	// nums := []int{}
	// x := 0
	// k := 0

	result := minOperations(nums, x, k)
	fmt.Printf("result = %v\n", result)
}
