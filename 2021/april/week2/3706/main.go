package main
import (
	"fmt"
	"math"
)

type NestedIterator struct {
    list []*NestedInteger
    i int
    cur *NestedIterator
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    return &NestedIterator{list: nestedList}
}

func (this *NestedIterator) Next() int {
    if this.list[this.i].IsInteger() {
        res := this.list[this.i].GetInteger()
        this.i++
        return res
    }

    return this.cur.Next()
}

func (this *NestedIterator) HasNext() bool {
    if this.i >= len(this.list) {
        return false
    }
    if this.list[this.i].IsInteger() {
        return true
    }
    if this.cur == nil {
        this.cur = Constructor(this.list[this.i].GetList())
    }

    if this.cur.HasNext() {
        return true
    }

    this.i++
    this.cur = nil

    return this.HasNext()
}

// NG Solution (Wrong Answer)
// type NestedIterator struct {
//     NestedList []*NestedInteger
// }
// 
// func Constructor(nestedList []*NestedInteger) *NestedIterator {
//     return &NestedIterator{NestedList: nestedList}
// }
// 
// func (this *NestedIterator) Next() int {
//     if !this.HasNext() {
//         return math.MinInt32
//     }
//     head := this.NestedList[0]
// 
//     if head.IsInteger() {
//         this.NestedList = this.NestedList[1:]
//         return head.GetInteger()
//     }
// 
//     list := head.GetList()
//     if len(list) > 0 {
//         current := list[0]
//         // fmt.Printf("current = %v\n", current)
//         for len(list) > 0 {
//             if current.IsInteger() {
//                 break
//             } else {
//                 list = current.List
//                 current = list[0]
//                 continue
//             }
//             current = list[0]
//             list = list[1:]
//         }
//         if len(list) > 1 {
//             list = list[1:]
//         } else {
//             list = []*NestedInteger{}
//             if len(this.NestedList) > 0 {
//                 this.NestedList = this.NestedList[1:]
//             } else {
//                 this.NestedList = []*NestedInteger{}
//             }
//         }
//         head.List = list
//         return current.GetInteger()
//     }
// 
//     if len(this.NestedList) > 0 {
//         this.NestedList = this.NestedList[1:]
//     } else {
//         this.NestedList = []*NestedInteger{}
//     }
//     return this.Next()
// }
// 
// func (this *NestedIterator) HasNext() bool {
//     // fmt.Printf("this.NestedList = %v\n", this.NestedList)
//     if len(this.NestedList) == 0 {
//         return false
//     }
//     head := this.NestedList[0]
// 
//     if head.IsInteger() {
//         return true
//     }
//     list := head.GetList()
//     if len(list) > 0 {
//         // fmt.Printf("list = %v\n", list[0])
//         if len(list) == 1 && !list[0].IsInteger() && len(list[0].GetList()) == 0 {
//             return false
//         }
//         return true
//     }
//     if len(this.NestedList) == 1 {
//         return false
//     }
//     return true
// }

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

