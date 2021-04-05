package main
import (
	"fmt"
)

type MyCircularQueue struct {
	size int
	nums []int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{size: k, nums: []int{}}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.nums = append(this.nums, value)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if len(this.nums) == 0 {
		return false
	}
	this.nums = this.nums[1:]
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.nums[0]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.nums[len(this.nums) - 1]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return len(this.nums) == 0
}

func (this *MyCircularQueue) IsFull() bool {
	return len(this.nums) >= this.size
}

func main() {
	k := int(3)
	obj := Constructor(k);
	fmt.Printf("EnQueue(1): %v\n", obj.EnQueue(1))
	fmt.Printf("EnQueue(2): %v\n", obj.EnQueue(2))
	fmt.Printf("EnQueue(3): %v\n", obj.EnQueue(3))
	fmt.Printf("EnQueue(4): %v\n", obj.EnQueue(4))
	fmt.Printf("Rear(): %v\n", obj.Rear())
	fmt.Printf("IsFull: %v\n", obj.IsFull())
	fmt.Printf("DeQueue(): %v\n", obj.DeQueue())
	fmt.Printf("Enqueue(4): %v\n", obj.EnQueue(4))
	fmt.Printf("Rear: %v\n", obj.Rear())

	// k := int(8)
	// obj := Constructor(k);
	// fmt.Printf("EnQueue(3): %v\n", obj.EnQueue(3))
	// fmt.Printf("EnQueue(9): %v\n", obj.EnQueue(9))
	// fmt.Printf("EnQueue(5): %v\n", obj.EnQueue(5))
	// fmt.Printf("EnQueue(0): %v\n", obj.EnQueue(0))
	// fmt.Printf("DeQueue(): %v\n", obj.DeQueue())
	// fmt.Printf("DeQueue(): %v\n", obj.DeQueue())
	// fmt.Printf("Rear(): %v\n", obj.Rear())
	// fmt.Printf("IsEmpty: %v\n", obj.IsEmpty())
	// fmt.Printf("Rear: %v\n", obj.Rear())
	// fmt.Printf("Rear: %v\n", obj.Rear())
	// fmt.Printf("DeQueue(): %v\n", obj.DeQueue())
}

