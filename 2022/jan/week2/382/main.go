package main
import (
	"fmt"
	"math/rand"
)

type ListNode struct {
    Val int
    Next *ListNode
}

type Solution struct {
    Nodes []*ListNode
}

func Constructor(head *ListNode) Solution {
	nodes := []*ListNode{}
	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}
    return Solution{Nodes: nodes}
}

func (this *Solution) GetRandom() int {
	random := rand.Int() % len(this.Nodes)
    return this.Nodes[random].Val
}

func makeList1() *ListNode {
    root := &ListNode{Val: 1}
    root.Next = &ListNode{Val: 2}
    root.Next.Next = &ListNode{Val: 3}
    return root
}

func makeList2() *ListNode {
    root := &ListNode{Val: 7}
    root.Next = &ListNode{Val: 2}
    root.Next.Next = &ListNode{Val: 3}
    root.Next.Next.Next = &ListNode{Val: 1}
    root.Next.Next.Next.Next = &ListNode{Val: -10}
    root.Next.Next.Next.Next.Next = &ListNode{Val: 14}
    return root
}

func makeList3() *ListNode {
    root := &ListNode{Val: 1}
    return root
}

func makeList4() *ListNode {
    root := &ListNode{Val: 10}
    root.Next = &ListNode{Val: 1}
    root.Next.Next = &ListNode{Val: 10}
    root.Next.Next.Next = &ListNode{Val: 20}
    root.Next.Next.Next.Next = &ListNode{Val: 100}
    return root
}

func main() {

	// result: [null, 1, 3, 2, 2, 3]
	// head := makeList1()
	// obj := Constructor(head)
	// fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
	// fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
	// fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
	// fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
	// fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())

	// result: [null, 3, 2, 1, 14, 14, 3, 7, 3]
	// head := makeList2()
	// obj := Constructor(head)
	// fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
	// fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
	// fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
	// fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
	// fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
	// fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
	// fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
	// fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())

	// result: [null, 1, 10, 1, 1, 10, 1, 100, 10, 1, 20]
	// head := makeList4()
	// obj := Constructor(head)
	// fmt.Println(obj)
	// fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
	// fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
	// fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
	// fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
	// fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
	// fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
	// fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
	// fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
	// fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
	// fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())

	// result: [null, 1, 1, 1]
	head := makeList3()
	obj := Constructor(head)
	fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
	fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
	fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())

	// result: 
	// head := makeList()
	// obj := Constructor(head)
	// fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
}

