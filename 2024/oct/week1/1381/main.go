package main
import (
	"fmt"
)

type CustomStack struct {
	Stack []int
	Increments []int
	TopIndex int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{Stack: make([]int, maxSize), Increments: make([]int, maxSize), TopIndex: -1}
}

func (this *CustomStack) Push(x int)  {
	if this.TopIndex >= len(this.Stack) - 1 {
		return
	}

	this.TopIndex++
	this.Stack[this.TopIndex] = x
}

func (this *CustomStack) Pop() int {
	if this.TopIndex < 0 {
		return -1
	}

	result := this.Stack[this.TopIndex] + this.Increments[this.TopIndex]
	if this.TopIndex > 0 {
		this.Increments[this.TopIndex - 1] += this.Increments[this.TopIndex]
	}

	this.Increments[this.TopIndex] = 0
	this.TopIndex--
	return result
}

func (this *CustomStack) Increment(k int, val int)  {
	if this.TopIndex < 0 {
		return
	}

	incrementIndex := min(this.TopIndex, k - 1)
	this.Increments[incrementIndex] += val
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: [null,null,null,2,null,null,null,null,null,103,202,201,-1]
	operations := [][]int{{0, 3}, {1, 1}, {1, 2}, {2, 0}, {1, 2}, {1, 3}, {1, 4}, {3, 5, 100}, {3, 2, 100}, {2, 0}, {2, 0}, {2, 0}, {2, 0}}

	obj := Constructor(operations[0][1])
	for _, ope := range operations {
		function := ope[0]

		if function == 1 {
			obj.Push(ope[1])
		} else if function == 2 {
			fmt.Printf("obj.Pop() = %d\n", obj.Pop())
		} else if function == 3 {
			obj.Increment(ope[1], ope[2])
		}
	}
}

