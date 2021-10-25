package main
import (
	"fmt"
	// "math"
	// "sort"
)

type MinStack struct {
    Stack[]int
    Mins []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{Stack: make([]int , 0)}
}


func (this *MinStack) Push(val int)  {
    this.Stack = append(this.Stack, val)
    if len(this.Mins) == 0 || val <= this.Mins[len(this.Mins) - 1]{
        this.Mins = append(this.Mins, val)
    }
}


func (this *MinStack) Pop()  {
    val := this.Top()
    this.Stack = this.Stack[:len(this.Stack) - 1]

    if val == this.Mins[len(this.Mins) - 1]{
        this.Mins = this.Mins[:len(this.Mins) - 1]
    }
}


func (this *MinStack) Top() int {
    return this.Stack[len(this.Stack) - 1]
}


func (this *MinStack) GetMin() int {
    if len(this.Mins) == 0 {
        return 0
    }

    return this.Mins[len(this.Mins) - 1]
}

// Slow & Use more memory
// type MinStack struct {
// 	Nums []int
// 	SortedNums []int
// }
// 
// func Constructor() MinStack {
// 	return MinStack{Nums: []int{}}
// }
// 
// func (this *MinStack) Push(val int)  {
// 	this.Nums = append(this.Nums, val)
// 
// 	if len(this.SortedNums) == 0 {
// 		this.SortedNums = append(this.SortedNums, val)
// 		return
// 	}
// 
//     index := sort.SearchInts(this.SortedNums, val)
// 	// fmt.Printf("this.SortedNums = %v, index = %v\n", this.SortedNums, index)
// 	heads := make([]int, len(this.SortedNums[:index]))
// 	copy(heads, this.SortedNums[:index])
// 	heads = append(heads, val)
// 	if len(this.SortedNums) == index {
// 		this.SortedNums = heads
// 	} else {
// 		this.SortedNums = append(heads, this.SortedNums[index:]...)
// 	}
// 	// fmt.Printf("after this.Nums = %v, this.SortedNums = %v\n", this.Nums, this.SortedNums)
// }
// 
// func (this *MinStack) Pop()  {
// 	if len(this.Nums) == 0 {
// 		return
// 	}
// 
// 	target := this.Nums[len(this.Nums) - 1]
// 	this.Nums = this.Nums[:len(this.Nums) - 1]
//     index := sort.SearchInts(this.SortedNums, target)
// 	// fmt.Printf("pop index = %v\n", index)
// 	if index == 0 {
// 		this.SortedNums = this.SortedNums[1:]
// 		// fmt.Printf("Pop: after this.Nums = %v, this.SortedNums = %v\n", this.Nums, this.SortedNums)
// 		return
// 	}
// 	this.SortedNums = append(this.SortedNums[:index], this.SortedNums[index + 1:]...)
// 	// fmt.Printf("Pop: after this.Nums = %v, this.SortedNums = %v\n", this.Nums, this.SortedNums)
// }
// 
// func (this *MinStack) Top() int {
// 	if len(this.Nums) == 0 {
// 		return math.MinInt32
// 	}
// 
// 	return this.Nums[len(this.Nums) - 1]
// }
// 
// func (this *MinStack) GetMin() int {
// 	if len(this.SortedNums) == 0 {
// 		return math.MinInt32
// 	}
// 
// 	return this.SortedNums[0]
// }

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	fmt.Printf("obj.GetMin() = %v\n", obj.GetMin())
	obj.Pop()
	fmt.Printf("obj.Top() = %v\n", obj.Top())
	fmt.Printf("obj.GetMin() = %v\n", obj.GetMin())
	// obj.Pop()
	// fmt.Printf("obj.Top() = %v\n", obj.Top())
	// fmt.Printf("obj.GetMin() = %v\n", obj.GetMin())
	// obj.Pop()
	// fmt.Printf("obj.Top() = %v\n", obj.Top())
	// fmt.Printf("obj.GetMin() = %v\n", obj.GetMin())

	// obj.Push(0)
	// obj.Pop()
	// fmt.Printf("obj.Pop() = %v\n", obj.Pop())
	// fmt.Printf("obj.Top() = %v\n", obj.Top())
	// fmt.Printf("obj.GetMin() = %v\n", obj.GetMin())
}

