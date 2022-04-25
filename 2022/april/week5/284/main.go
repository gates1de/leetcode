package main
import (
	"fmt"
)

type Iterator struct {
	Nums []int
	Cursor int
}

func (this *Iterator) hasNext() bool {
	return this.Cursor < len(this.Nums)
}

func (this *Iterator) next() int {
	result := this.Nums[this.Cursor]
	this.Cursor++
	return result
}

type PeekingIterator struct {
	Iterator *Iterator
	Nums []int
	Cursor int
}

func Constructor(iter *Iterator) *PeekingIterator {
	nums := []int{}
	for iter.hasNext() {
		nums = append(nums, iter.next())
	}
	return &PeekingIterator{Iterator: iter, Nums: nums, Cursor: 0}
}

func (this *PeekingIterator) hasNext() bool {
	return this.Cursor < len(this.Nums)
}

func (this *PeekingIterator) next() int {
	if !this.hasNext() {
		return 0
	}
	result := this.Nums[this.Cursor]
	this.Cursor++
	return result
}

func (this *PeekingIterator) peek() int {
	return this.Nums[this.Cursor]
}

func main() {
	// result: [null, 1, 2, 2, 3, false]
	nums := []int{1, 2, 3}
	operations := []int{1, 2, 1, 1, 3}

	iterator := &Iterator{Nums: nums}
	peekingIterator := Constructor(iterator)

	for _, ope := range operations {
		if ope == 1 {
			fmt.Printf("next = %v\n", peekingIterator.next())
		} else if ope == 2 {
			fmt.Printf("peek = %v\n", peekingIterator.peek())
		} else if ope == 3 {
			fmt.Printf("hasNext = %v\n", peekingIterator.hasNext())
		}
	}
}

