package main
import (
	"fmt"
)

type MyCircularDeque struct {
	Nums []int
	Size int
	Capacity int
	Front int
	Rear int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{Nums: make([]int, k), Size: 0, Capacity: k, Front: 0, Rear: k - 1}
}

func (this *MyCircularDeque) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}

	this.Front = (this.Front - 1 + this.Capacity) % this.Capacity
	this.Nums[this.Front] = value
	this.Size++
	return true
}

func (this *MyCircularDeque) InsertLast(value int) bool {
	if this.IsFull() {
		return false
	}

	this.Rear = (this.Rear + 1) % this.Capacity
	this.Nums[this.Rear] = value
	this.Size++
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}

	this.Front = (this.Front + 1) % this.Capacity
	this.Size--
	return true
}

func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}

	this.Rear = (this.Rear - 1 + this.Capacity) % this.Capacity
	this.Size--
	return true
}

func (this *MyCircularDeque) GetFront() int {
	if this.IsEmpty() {
		return -1
	}

	return this.Nums[this.Front]
}

func (this *MyCircularDeque) GetRear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.Nums[this.Rear]
}

func (this *MyCircularDeque) IsEmpty() bool {
	return this.Size == 0
}

func (this *MyCircularDeque) IsFull() bool {
	return this.Size == this.Capacity
}

func main() {
	// result: [null, true, true, true, false, 2, true, true, true, 4]
	operations := [][]int{{0, 3}, {2, 1}, {2, 2}, {1, 3}, {1, 4}, {6, 0}, {8, 0}, {4, 0}, {1, 4}, {5, 0}}

	// result: []
	// operations := [][]int{}

	// {0, n} -> init
	// {1, n} -> insertFront()
	// {2, n} -> insertLast()
	// {3, n} -> deleteFront()
	// {4, n} -> deleteLast()
	// {5, n} -> getFront()
	// {6, n} -> getRear()
	// {7, n} -> isEmpty()
	// {8, n} -> isFull()
	obj := Constructor(operations[0][1])
	for _, ope := range operations {
		function := ope[0]
		value := ope[1]

		if function == 1 {
			fmt.Printf("obj.InsertFront(%d) = %t\n", value, obj.InsertFront(value))
		} else if function == 2 {
			fmt.Printf("obj.InsertLast(%d) = %t\n", value, obj.InsertLast(value))
		} else if function == 3 {
			fmt.Printf("obj.DeleteFront() = %t\n", obj.DeleteFront())
		} else if function == 4 {
			fmt.Printf("obj.DeleteLast() = %t\n", obj.DeleteLast())
		} else if function == 5 {
			fmt.Printf("obj.GetFront() = %d\n", obj.GetFront())
		} else if function == 6 {
			fmt.Printf("obj.GetRear() = %d\n", obj.GetRear())
		} else if function == 7 {
			fmt.Printf("obj.IsEmpty() = %t\n", obj.IsEmpty())
		} else if function == 8 {
			fmt.Printf("obj.IsFull() = %t\n", obj.IsFull())
		}
	}
}

