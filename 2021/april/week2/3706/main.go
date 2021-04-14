package main
import (
	"fmt"
	"math"
)

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	Integer int
	Integers []int
	Next *NestedInteger
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {
	return len(this.Integers) > 0
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {
	if this.IsInteger() {
		result := this.Integer
		this = *this.Next
		return result
	}
	result := this.Integers[0]
	if len(this.Integers) == 1 {
		this = *this.Next
	} else {
		this.Integers = this.Integers[1:]
	}
	return result
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {
	n.Integer = value
}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {
	if this.Next == nil {
		this.Next = &elem
		return
	}
	this.Next.Add(elem)
}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {
	result := []*NestedInteger{&this}
	for this.Next != nil {
		result = append(result, this.Next)
		this.Next = this.Next.Next
	}
	return result
}

type NestedIterator struct {
    NestedList []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    return &NestedIterator{NestedList: nestedList}
}

func (this *NestedIterator) Next() int {
    if !this.HasNext() {
        return math.MinInt32
    }
    head := this.NestedList[0]

    if head.IsInteger() {
        this.NestedList = this.NestedList[1:]
        return head.GetInteger()
    }

    list := head.GetList()
    if len(list) > 0 {
        current := list[0]
        // fmt.Printf("current = %v\n", current)
        for len(list) > 0 {
            if current.IsInteger() {
                break
            } else {
                list = current.List
                current = list[0]
                continue
            }
            current = list[0]
            list = list[1:]
        }
        if len(list) > 1 {
            list = list[1:]
        } else {
            list = []*NestedInteger{}
            if len(this.NestedList) > 0 {
                this.NestedList = this.NestedList[1:]
            } else {
                this.NestedList = []*NestedInteger{}
            }
        }
        head.List = list
        return current.GetInteger()
    }

    if len(this.NestedList) > 0 {
        this.NestedList = this.NestedList[1:]
    } else {
        this.NestedList = []*NestedInteger{}
    }
    return this.Next()
}

func (this *NestedIterator) HasNext() bool {
    // fmt.Printf("this.NestedList = %v\n", this.NestedList)
    if len(this.NestedList) == 0 {
        return false
    }
    head := this.NestedList[0]

    if head.IsInteger() {
        return true
    }
    list := head.GetList()
    if len(list) > 0 {
        // fmt.Printf("list = %v\n", list[0])
        if len(list) == 1 && !list[0].IsInteger() && len(list[0].GetList()) == 0 {
            return false
        }
        return true
    }
    if len(this.NestedList) == 1 {
        return false
    }
    return true
}

func main() {
	ni1 := NestedInteger{Integers: []int{1, 1}}
	ni2 := NestedInteger{Integer: 1}
	ni3 := NestedInteger{Integers: []int{1, 1}}
	ni1.Add(ni2)
	ni2.Add(ni3)
	iter := NestedIterator{NestedList: ni1.GetList()}
	fmt.Printf("iter = %v\n", iter)
	for iter.HasNext() {
		fmt.Printf("iter.Next() = %v\n", iter.Next())
	}
}

