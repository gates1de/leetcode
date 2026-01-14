package main

import (
	"fmt"
	"sort"
)

type Event struct {
	Y int
	Delta int
	Xl int
	Xr int
}

type SegmentTree struct {
	Count   []int
	Covered []int
	Xs      []int
	N       int
}

func NewSegmentTree(xs []int) *SegmentTree {
	n := len(xs) - 1
	return &SegmentTree{
		Count:   make([]int, 4 * n),
		Covered: make([]int, 4 * n),
		Xs:      xs,
		N:       n,
	}
}

func (st *SegmentTree) modify(qLeft, qrRight, qVal, left, right, pos int) {
	if st.Xs[right + 1] <= qLeft || st.Xs[left] >= qrRight {
		return
	}

	if qLeft <= st.Xs[left] && st.Xs[right + 1] <= qrRight {
		st.Count[pos] += qVal
	} else {
		mid := (left + right) / 2
		st.modify(qLeft, qrRight, qVal, left, mid, pos * 2 + 1)
		st.modify(qLeft, qrRight, qVal, mid + 1, right, pos * 2 + 2)
	}

	if st.Count[pos] > 0 {
		st.Covered[pos] = st.Xs[right+1] - st.Xs[left]
	} else {
		if left == right {
			st.Covered[pos] = 0
		} else {
			st.Covered[pos] = st.Covered[pos * 2 + 1] + st.Covered[pos * 2 + 2]
		}
	}
}

func (st *SegmentTree) Update(qLeft, qrRight, qVal int) {
	st.modify(qLeft, qrRight, qVal, 0, st.N - 1, 0)
}

func (st *SegmentTree) Query() int {
	return st.Covered[0]
}

func separateSquares(squares [][]int) float64 {
	events := []Event{}
	xsSet := make(map[int]bool)

	for _, sq := range squares {
		x := sq[0]
		y := sq[1]
		l := sq[2]
		xr := x + l

		events = append(events, Event{y, 1, x, xr})
		events = append(events, Event{y + l, -1, x, xr})
		xsSet[x] = true
		xsSet[xr] = true
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Y < events[j].Y
	})

	xs := make([]int, 0, len(xsSet))
	for x := range xsSet {
		xs = append(xs, x)
	}
	sort.Ints(xs)

	segTree := NewSegmentTree(xs)

	pSum := make([]float64, 0)
	widths := make([]float64, 0)
	totalArea := float64(0)
	prev := events[0].Y

	for _, event := range events {
		y := event.Y
		delta := event.Delta
		xl := event.Xl
		xr := event.Xr
		length := segTree.Query()
		totalArea += float64(length) * float64(y - prev)
		segTree.Update(xl, xr, delta)
		pSum = append(pSum, totalArea)
		widths = append(widths, float64(segTree.Query()))
		prev = y
	}

	target := int64(totalArea + 1) / 2
	i := sort.Search(len(pSum), func(i int) bool {
		return pSum[i] >= float64(target)
	})
	i--

	area := pSum[i]
	width := widths[i]
	height := events[i].Y

	return float64(height) + (totalArea - area * 2) / (float64(width) * 2.0)
}

func main() {
	// result: 1.00000
	// squares := [][]int{{0,0,1}, {2,2,1}}

	// result: 1.00000
	squares := [][]int{{0,0,2}, {1,1,1}}

	// result: 
	// squares := [][]int{}

	result := separateSquares(squares)
	fmt.Printf("result = %v\n", result)
}

