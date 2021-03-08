package main
import (
	"fmt"
)

type Iterator struct {
	nums []int
	cursor int
}

func (this *Iterator) hasNext() bool {
	return this.cursor < len(this.nums)
}

func (this *Iterator) next() int {
	result := this.nums[this.cursor]
	this.cursor++
	return result
}


type PeekingIterator struct {
	iterator *Iterator
	nums []int
	cursor int
}

func Constructor(iter *Iterator) *PeekingIterator {
	nums := []int{}
	for iter.hasNext() {
		nums = append(nums, iter.next())
	}
	return &PeekingIterator{iterator: iter, nums: nums, cursor: 0}
}

func (this *PeekingIterator) hasNext() bool {
	return this.cursor < len(this.nums)
}

func (this *PeekingIterator) next() int {
	if !this.hasNext() {
		return 0
	}
	result := this.nums[this.cursor]
	this.cursor++
	return result
}

func (this *PeekingIterator) peek() int {
	return this.nums[this.cursor]
}

func main() {
	nums := []int{1, 2, 3}
	iterator := &Iterator{nums: nums}
	peekingIterator := Constructor(iterator)
	fmt.Printf("next = %v\n", peekingIterator.next())
	fmt.Printf("peek = %v\n", peekingIterator.peek())
	fmt.Printf("next = %v\n", peekingIterator.next())
	fmt.Printf("next = %v\n", peekingIterator.next())
	result := peekingIterator.hasNext()
	fmt.Printf("result = %v\n", result)
}

