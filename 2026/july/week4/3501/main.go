package main

import (
	"fmt"
	"sort"
)

func maxActiveSectionsAfterTrade(s string, queries [][]int) []int {
	n := len(s)
	prefixOnes := make([]int, n + 1)
	for i := range n {
		prefixOnes[i + 1] = prefixOnes[i]
		if s[i] == '1' {
			prefixOnes[i + 1]++
		}
	}

	var start, end []int
	var one []bool
	for i := 0; i < n; {
		j := i + 1
		for j < n && s[j] == s[i] {
			j++
		}

		start = append(start, i)
		end = append(end, j)
		one = append(one, s[i] == '1')
		i = j
	}

	m := len(start)
	gain := make([]int, m)
	for i := range gain {
		gain[i] = -1
		if i > 0 && i + 1 < m && one[i] && !one[i - 1] && !one[i + 1] {
			gain[i] = (end[i - 1] - start[i - 1]) + (end[i + 1] - start[i + 1])
		}
	}

	size := int(1)
	for size < m {
		size <<= 1
	}

	tree := make([]int, size*2)
	for i := range tree {
		tree[i] = -1
	}
	for i := range m {
		tree[size + i] = gain[i]
	}

	for i := size - 1; i > 0; i-- {
		tree[i] = max(tree[i * 2], tree[i * 2 + 1])
	}

	maxRange := func(left, right int) int {
		if left >= right {
			return -1
		}

		left += size
		right += size
		result := -1
		for left < right {
			if left & 1 == 1 {
				if tree[left] > result {
					result = tree[left]
				}
				left++
			}

			if right&1 == 1 {
				right--
				if tree[right] > result {
					result = tree[right]
				}
			}

			left >>= 1
			right >>= 1
		}

		return result
	}

	answers := make([]int, len(queries))
	totalOnes := prefixOnes[n]
	for qi, query := range queries {
		l, r := query[0], query[1]
		leftRun := sort.Search(m, func(i int) bool { return start[i] > l }) - 1
		rightRun := sort.Search(m, func(i int) bool { return start[i] > r }) - 1

		bestGain := int(0)
		boundaryGain := func(c int) int {
			if c <= leftRun || c >= rightRun || !one[c] || one[c-1] || one[c+1] {
				return -1
			}

			leftZero := min(end[c-1], r + 1)
			leftZero -= start[c-1]
			if l > start[c-1] {
				leftZero -= l - start[c-1]
			}

			rightZero := min(end[c+1], r + 1)
			rightZero -= start[c+1]
			if l > start[c+1] {
				rightZero -= l - start[c+1]
			}

			return leftZero + rightZero
		}

		if candidate := boundaryGain(leftRun + 1); candidate > bestGain {
			bestGain = candidate
		}

		if candidate := boundaryGain(rightRun - 1); candidate > bestGain {
			bestGain = candidate
		}

		if candidate := maxRange(leftRun+2, rightRun-1); candidate > bestGain {
			bestGain = candidate
		}
		answers[qi] = totalOnes + bestGain
	}

	return answers
}

func main() {
	// result: [1]
	// s := "01"
	// queries := [][]int{{0,1}}

	// result: [4,3,1,1]
	// s := "0100"
	// queries := [][]int{{0,3},{0,2},{1,3},{2,3}}

	// result: [6,7,2]
	// s := "1000100"
	// queries := [][]int{{1,5},{0,6},{0,4}}

	// result: [4,4,2]
	s := "01010"
	queries := [][]int{{0, 3}, {1, 4}, {1, 3}}

	// result: []
	// s := ""
	// queries := [][]int{}

	result := maxActiveSectionsAfterTrade(s, queries)
	fmt.Printf("result = %v\n", result)
}
