package main

import (
	"container/heap"
	"fmt"
)

type Entry struct {
	Val      	int
	HeapIndex int
	HeapMark 	int
}

type HeapMin []*Entry

func (this *HeapMin) Len() int {
	return len(*this)
}

func (this *HeapMin) Less(i, j int) bool {
	s := *this
	return s[i].Val < s[j].Val
}

func (this *HeapMin) Swap(i, j int) {
	s := *this
	s[i], s[j] = s[j], s[i]
	s[i].HeapIndex = int(i)
	s[j].HeapIndex = int(j)
}

func (this *HeapMin) Push(x any) {
	e := x.(*Entry)
	e.HeapIndex = int(len(*this))
	e.HeapMark = 0
	*this = append(*this, e)
}

func (this *HeapMin) Pop() any {
	s := *this
	l := len(s) - 1
	result := s[l]
	result.HeapIndex = -1
	result.HeapMark = -1
	*this = s[:l]

	return result
}

type HeapMax struct {
	S   []*Entry
	Sum int64
}

func (this *HeapMax) Len() int {
	return len(this.S)
}

func (this *HeapMax) Less(i, j int) bool {
	return this.S[j].Val < this.S[i].Val
}

func (this *HeapMax) Swap(i, j int) {
	this.S[i], this.S[j] = this.S[j], this.S[i]
	this.S[i].HeapIndex = int(i)
	this.S[j].HeapIndex = int(j)
}

func (this *HeapMax) Push(x any) {
	e := x.(*Entry)
	e.HeapIndex = int(len(this.S))
	e.HeapMark = 1
	this.S = append(this.S, e)
	this.Sum += int64(e.Val)
}

func (this *HeapMax) Pop() any {
	l := len(this.S) - 1
	result := this.S[l]
	result.HeapIndex = -1
	result.HeapMark = -1
	this.S = this.S[:l]
	this.Sum -= int64(result.Val)

	return result
}

func minimumCost(nums []int, k int, dist int) int64 {
	k2 := k - 2
	if dist == k2 {
		minSum := int64(0)
		for i := 1; i < k; i++ {
			minSum += int64(nums[i])
		}

		for i, sum := 2, minSum; i + k2 < len(nums); i++ {
			sum += int64(nums[i + k2]) - int64(nums[i - 1])
			minSum = min(minSum, sum)
		}

		return int64(nums[0]) + minSum
	}

	entries := make([]Entry, len(nums))
	for i := range entries {
		entries[i].Val = int(nums[i])
		entries[i].HeapIndex = -1
		entries[i].HeapMark = -1
	}

	heapMax := HeapMax{S: make([]*Entry, 0, k - 2)}
	for i := 2; i < k; i++ {
		heap.Push(&heapMax, &entries[i])
	}

	heapMin := make(HeapMin, 0, dist - k + 2)
	for i := k; i < dist + 2; i++ {
		if heapMax.S[0].Val <= entries[i].Val {
			heap.Push(&heapMin, &entries[i])
		} else {
			heap.Push(&heapMin, heap.Pop(&heapMax))
			heap.Push(&heapMax, &entries[i])
		}
	}

	minSum := int64(nums[1]) + heapMax.Sum
	for i := 3; i+dist <= len(entries); i++ {
		if entries[i - 1].HeapMark == 1 {
			heap.Remove(&heapMax, int(entries[i - 1].HeapIndex))
			heap.Push(&heapMax, heap.Pop(&heapMin))
		} else {
			heap.Remove(&heapMin, int(entries[i - 1].HeapIndex))
		}

		heap.Push(&heapMin, &entries[i + dist - 1])
		if heapMin[0].Val < heapMax.S[0].Val {
			heap.Push(&heapMax, heap.Pop(&heapMin))
			heap.Push(&heapMin, heap.Pop(&heapMax))
		}

		minSum = min(minSum, int64(nums[i - 1])+ heapMax.Sum)
	}

	for i := len(entries) - dist + 1; 0 < len(heapMin); i++ {
		if entries[i - 1].HeapMark == 1 {
			heap.Remove(&heapMax, int(entries[i - 1].HeapIndex))
			heap.Push(&heapMax, heap.Pop(&heapMin))
		} else {
			heap.Remove(&heapMin, int(entries[i - 1].HeapIndex))
		}

		if 0 < len(heapMin) && heapMin[0].Val < heapMax.S[0].Val {
			heap.Push(&heapMax, heap.Pop(&heapMin))
			heap.Push(&heapMin, heap.Pop(&heapMax))
		}

		minSum = min(minSum, int64(nums[i - 1])+ heapMax.Sum)
	}

	return int64(nums[0]) + minSum
}

func main() {
	// result: 5
	// nums := []int{1,3,2,6,4,2}
	// k := int(3)
	// dist := int(3)

	// result: 15
	// nums := []int{10,1,2,2,2,1}
	// k := int(4)
	// dist := int(3)

	// result: 36
	nums := []int{10,8,18,9}
	k := int(3)
	dist := int(1)

	// result: 
	// nums := []int{}
	// k := int(0)
	// dist := int(0)

	result := minimumCost(nums, k, dist)
	fmt.Printf("result = %v\n", result)
}

