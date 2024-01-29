package main
import (
	"fmt"
)

type MyQueue struct {
    Queue []int
}

func Constructor() MyQueue {
    return MyQueue{Queue: make([]int, 0, 100)}
}

func (this *MyQueue) Push(x int)  {
    this.Queue = append(this.Queue, x)
}

func (this *MyQueue) Pop() int {
    if this.Empty() {
        return -1
    }

    result := this.Queue[0]
    this.Queue = this.Queue[1:]
    return result
}

func (this *MyQueue) Peek() int {
    return this.Queue[0]
}

func (this *MyQueue) Empty() bool {
    return len(this.Queue) == 0
}

func main() {
	// result: [null, null, null, 1, 1, false]
	operations := [][]int{
		{0, 1},
		{0, 2},
		{2, 0},
		{1, 0},
		{3, 0},
	}

	// result: 
	// operations := [][]int{
	// 	{0, 0}, // push
	// 	{1, 0}, // pop
	// 	{2, 0}, // peak
	// 	{3, 0}, // empty
	// }

	obj := Constructor()
	for _, ope := range operations {
		switch ope[0] {
		case 0:
			obj.Push(ope[1])
			fmt.Printf("obj.Push(%v)\n", ope[1])
		case 1:
			fmt.Printf("obj.Pop = %v\n", obj.Pop())
		case 2:
			fmt.Printf("obj.Peek = %v\n", obj.Peek())
		case 3:
			fmt.Printf("obj.Empty = %v\n", obj.Empty())
		}
	}
}

