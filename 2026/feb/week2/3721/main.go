package main

import (
	"fmt"
)

type LazyTag struct {
	ToAdd int
}

func (l *LazyTag) Add(other *LazyTag) *LazyTag {
	l.ToAdd += other.ToAdd
	return l
}

func (l *LazyTag) HasTag() bool {
	return l.ToAdd != 0
}

func (l *LazyTag) Clear() {
	l.ToAdd = 0
}

type SegmentTreeNode struct {
	MinValue int
	MaxValue int
	LazyTag  *LazyTag
}

func NewSegmentTreeNode() *SegmentTreeNode {
	return &SegmentTreeNode{
			MinValue: 0,
			MaxValue: 0,
			LazyTag:  &LazyTag{},
	}
}

type SegmentTree struct {
	N    int
	Tree []*SegmentTreeNode
}

func NewSegmentTree(data []int) *SegmentTree {
	n := len(data)
	tree := make([]*SegmentTreeNode, n*4+1)
	for i := range tree {
			tree[i] = NewSegmentTreeNode()
	}

	seg := &SegmentTree{N: n, Tree: tree}
	seg.Build(data, 1, n, 1)
	return seg
}

func (seg *SegmentTree) Add(l, r, val int) {
	tag := &LazyTag{ToAdd: val}
	seg.Update(l, r, tag, 1, seg.N, 1)
}

func (seg *SegmentTree) FindLast(start, val int) int {
	if start > seg.N {
			return -1
	}
	return seg.Find(start, seg.N, val, 1, seg.N, 1)
}

func (seg *SegmentTree) ApplyTag(i int, tag *LazyTag) {
	seg.Tree[i].MinValue += tag.ToAdd
	seg.Tree[i].MaxValue += tag.ToAdd
	seg.Tree[i].LazyTag.Add(tag)
}

func (seg *SegmentTree) Pushdown(i int) {
	if seg.Tree[i].LazyTag.HasTag() {
			tag := &LazyTag{ToAdd: seg.Tree[i].LazyTag.ToAdd}
			seg.ApplyTag(i<<1, tag)
			seg.ApplyTag((i<<1)|1, tag)
			seg.Tree[i].LazyTag.Clear()
	}
}

func (seg *SegmentTree) Pushup(i int) {
	left := seg.Tree[i << 1]
	right := seg.Tree[(i << 1)|1]
	seg.Tree[i].MinValue = min(left.MinValue, right.MinValue)
	seg.Tree[i].MaxValue = max(left.MaxValue, right.MaxValue)
}

func (seg *SegmentTree) Build(data []int, l, r, i int) {
	if l == r {
			seg.Tree[i].MinValue = data[l - 1]
			seg.Tree[i].MaxValue = data[l - 1]
			return
	}

	mid := l + ((r - l) >> 1)
	seg.Build(data, l, mid, i << 1)
	seg.Build(data, mid + 1, r, (i << 1) | 1)
	seg.Pushup(i)
}

func (seg *SegmentTree) Update(targetL, targetR int, tag *LazyTag, l, r, i int) {
	if targetL <= l && r <= targetR {
			seg.ApplyTag(i, tag)
			return
	}

	seg.Pushdown(i)
	mid := l + ((r - l) >> 1)
	if targetL <= mid {
			seg.Update(targetL, targetR, tag, l, mid, i<<1)
	}
	if targetR > mid {
			seg.Update(targetL, targetR, tag, mid+1, r, (i<<1)|1)
	}
	seg.Pushup(i)
}

func (seg *SegmentTree) Find(targetL, targetR, val, l, r, i int) int {
	if seg.Tree[i].MinValue > val || seg.Tree[i].MaxValue < val {
			return -1
	}

	if l == r {
			return l
	}

	seg.Pushdown(i)
	mid := l + ((r - l) >> 1)

	if targetR >= mid+1 {
			res := seg.Find(targetL, targetR, val, mid + 1, r, (i << 1) | 1)
			if res != -1 {
					return res
			}
	}

	if l <= targetR && mid >= targetL {
			return seg.Find(targetL, targetR, val, l, mid, i << 1)
	}

	return -1
}

func longestBalanced(nums []int) int {
	occurrences := make(map[int][]int)

	sgn := func(x int) int {
		if x % 2 == 0 {
			return 1
		}

		return -1
	}

	length := int(0)
	prefixSum := make([]int, len(nums))
	prefixSum[0] = sgn(nums[0])
	occurrences[nums[0]] = append(occurrences[nums[0]], 1)

	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i - 1]
		occ := occurrences[nums[i]]
		if len(occ) == 0 {
			prefixSum[i] += sgn(nums[i])
		}
		occurrences[nums[i]] = append(occ, i + 1)
	}

	seg := NewSegmentTree(prefixSum)
	for i := range len(nums) {
		length = max(length, seg.FindLast(i + length, 0) - i)
		nextPos := len(nums) + 1
		occurrences[nums[i]] = occurrences[nums[i]][1:]
		if len(occurrences[nums[i]]) > 0 {
				nextPos = occurrences[nums[i]][0]
		}

		seg.Add(i + 1, nextPos - 1, -sgn(nums[i]))
	}

	return length
}

func main() {
	// result: 4
	// nums := []int{2,5,4,3}

	// result: 5
	// nums := []int{3,2,2,5,4}

	// result: 3
	nums := []int{1,2,3,2}

	// result: 
	// nums := []int{}

	result := longestBalanced(nums)
	fmt.Printf("result = %v\n", result)
}

