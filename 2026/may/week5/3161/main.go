package main

import (
	"fmt"
)

type Fenwick struct {
	Tree []int
	N    int
}

func NewFenwick(n int) *Fenwick {
	return &Fenwick{Tree: make([]int, n + 1), N: n}
}

func (f *Fenwick) Add(idx, delta int) {
	for idx <= f.N {
		f.Tree[idx] += delta
		idx += idx & -idx
	}
}

func (f *Fenwick) Sum(idx int) int {
	if idx > f.N {
		idx = f.N
	}
	if idx < 0 {
		return int(0)
	}
	res := int(0)
	for idx > 0 {
		res += f.Tree[idx]
		idx -= idx & -idx
	}
	return res
}

// Kth returns the smallest index such that prefix sum >= k.
func (f *Fenwick) Kth(k int) int {
	idx := int(0)
	bitMask := int(1)
	for (bitMask << 1) <= f.N {
		bitMask <<= 1
	}
	for bitMask > 0 {
		next := idx + bitMask
		if next <= f.N && f.Tree[next] < k {
			idx = next
			k -= f.Tree[next]
		}
		bitMask >>= 1
	}
	return idx + 1
}

type SegTree struct {
	Size int
	Tree []int
}

func NewSegTree(n int) *SegTree {
	size := int(1)
	for size < n {
		size <<= 1
	}
	return &SegTree{
		Size: size,
		Tree: make([]int, size*2),
	}
}

func (s *SegTree) Update(idx, val int) {
	i := idx + s.Size - 1
	s.Tree[i] = val
	for i >>= 1; i > 0; i >>= 1 {
		left, right := s.Tree[i<<1], s.Tree[i<<1|1]
		s.Tree[i] = max(left, right)
	}
}

func (s *SegTree) Query(l, r int) int {
	if l > r {
		return int(0)
	}

	l += s.Size - 1
	r += s.Size - 1
	result := int(0)
	for l <= r {
		if l & 1 == 1 {
			if s.Tree[l] > result {
				result = s.Tree[l]
			}
			l++
		}

		if r & 1 == 0 {
			if s.Tree[r] > result {
				result = s.Tree[r]
			}
			r--
		}

		l >>= 1
		r >>= 1
	}

	return result
}

func getResults(queries [][]int) []bool {
	maxCoord := int(0)
	for _, q := range queries {
		if q[0] == 1 || q[0] == 2 {
			if q[1] > maxCoord {
				maxCoord = q[1]
			}
		}
	}

	bit := NewFenwick(maxCoord + 2)
	segt := NewSegTree(maxCoord + 2)
	gap := make([]int, maxCoord+3)

	results := make([]bool, 0)

	for _, q := range queries {
		switch q[0] {
		case 1:
			x := q[1]

			leftCount := bit.Sum(x - 1)
			left := int(0)
			if leftCount > 0 {
				left = bit.Kth(leftCount)
			}

			right := int(0)
			total := bit.Sum(maxCoord + 1)
			if leftCount < total {
				right = bit.Kth(leftCount + 1)
			}

			bit.Add(x, 1)

			gap[x] = x - left
			segt.Update(x, gap[x])

			if right != 0 {
				gap[right] = right - x
				segt.Update(right, gap[right])
			}

		case 2:
			x, sz := q[1], q[2]

			leftCount := bit.Sum(x - 1)
			left := int(0)
			if leftCount > 0 {
				left = bit.Kth(leftCount)
			}

			best := segt.Query(1, x)
			if tail := x - left; tail > best {
				best = tail
			}

			results = append(results, best >= sz)
		}
	}

	return results
}

func main() {
	// result: [false,true,true]
	// queries := [][]int{{1,2},{2,3,3},{2,3,1},{2,2,2}}

	// result: [true,false,false]
	queries := [][]int{{1,7},{2,7,6},{1,2},{2,7,5},{2,7,6}}

	// result: []
	// queries := [][]int{}

	result := getResults(queries)
	fmt.Printf("result = %v\n", result)
}
