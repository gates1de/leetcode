package main
import (
	"fmt"
)

type MyCircularQueue struct {
    Size int
    Nums []int
}

func Constructor(k int) MyCircularQueue {
    return MyCircularQueue{Size: k, Nums: []int{}}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
    if this.IsFull() {
        return false
    }
    this.Nums = append(this.Nums, value)
    return true
}

func (this *MyCircularQueue) DeQueue() bool {
    if len(this.Nums) == 0 {
        return false
    }
    this.Nums = this.Nums[1:]
    return true
}

func (this *MyCircularQueue) Front() int {
    if this.IsEmpty() {
        return -1
    }
    return this.Nums[0]
}

func (this *MyCircularQueue) Rear() int {
    if this.IsEmpty() {
        return -1
    }
    return this.Nums[len(this.Nums) - 1]
}

func (this *MyCircularQueue) IsEmpty() bool {
    return len(this.Nums) == 0
}

func (this *MyCircularQueue) IsFull() bool {
    return len(this.Nums) >= this.Size
}

func main() {
	// operation as below
	// enqueue: [0, val], dequeue: [1, 0], front: [2, 0], rear: [3, 0], isEmpty: [4, 0], isFull: [5, 0]

	// result: [null,true,true,true,false,3,true,true,true,4]
	queue := MyCircularQueue{Size: 3}
	operations := [][2]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{0, 4},
		{3, 0},
		{5, 0},
		{1, 0},
		{0, 4},
		{3, 0},
	}

	for _, ope := range operations {
		if ope[0] == 0 {
			fmt.Printf("EnQueue(%v) = %v\n", ope[1], queue.EnQueue(ope[1]))
		} else if ope[0] == 1 {
			fmt.Printf("DeQueue() = %v\n", queue.DeQueue())
		} else if ope[0] == 2 {
			fmt.Printf("Front() = %v\n", queue.Front())
		} else if ope[0] == 3 {
			fmt.Printf("Rear() = %v\n", queue.Rear())
		} else if ope[0] == 4 {
			fmt.Printf("IsEmpty() = %v\n", queue.IsEmpty())
		} else if ope[0] == 5 {
			fmt.Printf("IsFull() = %v\n", queue.IsFull())
		}
	}
}

