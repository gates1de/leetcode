package main
import (
	"fmt"
)

type MyStack struct {
    Queue []int
}

func Constructor() MyStack {
    return MyStack{Queue: []int{}}
}

func (this *MyStack) Push(x int)  {
    this.Queue = append(this.Queue, x)
}

func (this *MyStack) Pop() int {
    if this.Empty() {
        return -1
    }

    result := this.Queue[len(this.Queue) - 1]

    if len(this.Queue) == 1 {
        this.Queue = []int{}
    } else {
        this.Queue = this.Queue[:len(this.Queue) - 1]
    }

    return result
}

func (this *MyStack) Top() int {
    if this.Empty() {
        return -1
    }

    return this.Queue[len(this.Queue) - 1]
}

func (this *MyStack) Empty() bool {
    return len(this.Queue) == 0
}

func main() {
	// result: [null, null, null, 2, 2, false]
	operations := [][]int{{0,1},{0,2},{1,0},{2,0},{3,0}}

	// result: []
	// operations := [][]int{{0,0},{1,0},{2,0},{3,0}}

	obj := Constructor()
	for _, ope := range operations {
		if ope[0] == 0 {
			obj.Push(ope[1])
		} else if ope[0] == 1 {
			fmt.Printf("obj.Pop() = %v\n", obj.Pop())
		} else if ope[0] == 2 {
			fmt.Printf("obj.Top() = %v\n", obj.Top())
		} else if ope[0] == 3 {
			fmt.Printf("obj.Empty() = %v\n", obj.Empty())
		}
	}
}

