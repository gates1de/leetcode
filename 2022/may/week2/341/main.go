package main
import (
	"fmt"
	"math"
)

type NestedInteger struct {
	Integer int
	List []*NestedInteger
}

func (this NestedInteger) IsInteger() bool {
	return this.Integer != math.MinInt32
}

func (this NestedInteger) GetInteger() int {
	return this.Integer
}

func (n *NestedInteger) SetInteger(value int) {
	n.Integer = value
}

func (this *NestedInteger) Add(elem NestedInteger) {
	if this.List == nil || len(this.List) == 0 {
		this.List = []*NestedInteger{&elem}
		return
	}

	this.List = append(this.List, &elem)
}

func (this NestedInteger) GetList() []*NestedInteger {
	return this.List
}


type NestedIterator struct {
	Values []int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	values := []int{}
	if len(nestedList) == 0 {
		return &NestedIterator{Values: values}
	}

	for i := 0; i < len(nestedList); i++ {
		target := nestedList[i]
		if target.IsInteger() {
			values = append(values, target.GetInteger())
			continue
		}

		newNestedList := target.GetList()
		newNestedList = append(newNestedList, nestedList[i + 1:]...)
		nestedList = newNestedList
		i = -1
	}

	return &NestedIterator{Values: values}
}

func (this *NestedIterator) Next() int {
	if !this.HasNext() {
		return math.MinInt32
	}

	result := this.Values[0]
	this.Values = this.Values[1:]
	return result
}

func (this *NestedIterator) HasNext() bool {
	return len(this.Values) > 0
}

func main() {
	// result: [1,1,2,1,1]
	// ni1 := &NestedInteger{}
	// ni1.SetInteger(math.MinInt32)
	// ni1.Add(NestedInteger{Integer: 1})
	// ni1.Add(NestedInteger{Integer: 1})
	// ni2 := &NestedInteger{}
	// ni2.SetInteger(2)
	// ni3 := &NestedInteger{}
	// ni3.SetInteger(math.MinInt32)
	// ni3.Add(NestedInteger{Integer: 1})
	// ni3.Add(NestedInteger{Integer: 1})
	// list := []*NestedInteger{ni1, ni2, ni3}

	// result: [1,4,6]
	// ni1 := &NestedInteger{}
	// ni1.SetInteger(1)
	// ni2 := &NestedInteger{}
	// ni2.SetInteger(math.MinInt32)
	// ni2.Add(NestedInteger{Integer: 4})
	// ni3 := &NestedInteger{}
	// ni3.SetInteger(math.MinInt32)
	// ni3.Add(NestedInteger{Integer: 6})
	// ni2.Add(*ni3)
	// list := []*NestedInteger{ni1, ni2}

	// result: [1]
	// ni1 := &NestedInteger{}
	// ni1.SetInteger(1)
	// ni2 := &NestedInteger{}
	// ni2.SetInteger(math.MinInt32)
	// list := []*NestedInteger{ni1, ni2}

	// result: [8,4]
	ni1 := &NestedInteger{}
	ni1.SetInteger(math.MinInt32)
	ni2 := &NestedInteger{}
	ni2.SetInteger(math.MinInt32)
	ni3 := &NestedInteger{}
	ni3.SetInteger(math.MinInt32)
	ni3.Add(NestedInteger{Integer: 8})
	ni4 := &NestedInteger{}
	ni4.SetInteger(4)
	ni2.Add(*ni3)
	ni2.Add(*ni4)
	ni1.Add(*ni2)
	list := []*NestedInteger{ni1}

	// result: 
	// ni1 := &NestedInteger{}
	// list := []*NestedInteger{ni1}

	obj := Constructor(list)
	for obj.HasNext() {
		fmt.Printf("obj.Next() = %v\n", obj.Next())
	}
}

