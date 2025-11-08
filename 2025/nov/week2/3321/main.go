package main

import (
	"container/heap"
	"fmt"
)

type (
	PQMin = [0]bool
	PQMax = [1]bool
)

type PQKinds interface {
	PQMin | PQMax
}

type Freq struct {
	Val   	int
	Count		int
	Index 	int
	isTopX  bool
}

type PQFreqs[K PQKinds] []*Freq

func (this *PQFreqs[_]) Len() int {
	return len(*this)
}

func (this *PQFreqs[K]) Less(i, j int) bool {
	s := *this
	if isMax[K]() {
		return isLess(s[j], s[i])
	}

	return isLess(s[i], s[j])
}

func (this *PQFreqs[_]) Swap(i, j int) {
	s := *this
	s[i], s[j] = s[j], s[i]
	s[i].Index = i
	s[j].Index = j
}

func (this *PQFreqs[_]) Push(x any) {
	f := x.(*Freq)
	f.Index = len(*this)
	*this = append(*this, f)
}

func (this *PQFreqs[_]) Pop() any {
	s := *this
	l := len(s) - 1
	result := s[l]
	*this = s[:l]
	return result
}

type XSum struct {
	X     int
	Sum   int64
	Freqs map[int]*Freq
	TopX  PQFreqs[PQMin]
	Other PQFreqs[PQMax]
}

func newXSum(x int) *XSum {
	return &XSum{X: x, Freqs: make(map[int]*Freq)}
}

func (this *XSum) RemoveFromTopX(f *Freq) {
	this.Sum -= int64(f.Val) * int64(f.Count)
	f.isTopX = false
	heap.Remove(&this.TopX, int(f.Index))
}

func (this *XSum) TryToRemoveFromOther(f *Freq) {
	if 0 < f.Count {
		heap.Remove(&this.Other, int(f.Index))
	}
}

func (this *XSum) TryToRemove(f *Freq) {
	if f.isTopX {
		this.RemoveFromTopX(f)
	} else {
		this.TryToRemoveFromOther(f)
	}
}

func (this *XSum) PushToTopX(f *Freq) {
	f.isTopX = true
	this.Sum += int64(f.Val) * int64(f.Count)
	heap.Push(&this.TopX, f)
}

func (this *XSum) TryToPushToOther(f *Freq) {
	if 0 < f.Count {
		f.isTopX = false
		heap.Push(&this.Other, f)
	}
}

func (this *XSum) TryToPropagate(f *Freq) {
	this.TryToPushToOther(f)

	if 0 < len(this.Other) && this.X <= len(this.TopX) && isLess(this.TopX[0], this.Other[0]) {
		f = this.TopX[0]
		this.RemoveFromTopX(f)
		this.TryToPushToOther(f)
	}

	if 0 < len(this.Other) && len(this.TopX) < this.X {
		f = this.Other[0]
		this.TryToRemoveFromOther(f)
		this.PushToTopX(f)
	}
}

func (this *XSum) Push(n int) {
	f := this.Freqs[n]
	if f == nil {
		f = &Freq{Val: n}
		this.Freqs[n] = f
	} else {
		this.TryToRemove(f)
	}

	f.Count++
	this.TryToPropagate(f)
}

func (this *XSum) Pop(n int) {
	f := this.Freqs[n]
	this.TryToRemove(f)
	f.Count--
	this.TryToPropagate(f)
}

func isLess(f1, f2 *Freq) bool {
	if f1.Count == f2.Count {
		return f1.Val < f2.Val
	}

	return f1.Count < f2.Count
}

func isMax[K PQKinds]() bool {
	var kind K
	return 0 < len(kind)
}

func findXSum(nums []int, k int, x int) []int64 {
	n := len(nums)
	s := newXSum(x)
	for _, n := range nums[:k] {
		s.Push(n)
	}

	result := make([]int64, n - k + 1)
	result[0] = s.Sum
	for i := k; i < n; i++ {
		n1 := nums[i - k]
		n2 := nums[i]

		if n1 != n2 {
			s.Pop(n1)
			s.Push(n2)
		}
		result[i - k + 1] = s.Sum
	}

	return result
}

func main() {
	// result: [6,10,12]
	// nums := []int{1,1,2,2,3,4,2,3}
	// k := int(6)
	// x := int(2)

	// result: [11,15,15,15,12]
	nums := []int{3,8,7,8,7,5}
	k := int(2)
	x := int(2)

	// result: []
	// nums := []int{}
	// k := int(0)
	// x := int(0)

	result := findXSum(nums, k, x)
	fmt.Printf("result = %v\n", result)
}

