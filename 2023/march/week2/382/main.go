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
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	return head
}

func makeList() *ListNode {
	var head *ListNode
	return head
}

func main() {
	// result: [null, 1, 3, 2, 2, 3]
	head := makeList1()
	operationCount := int(5)

	// result: 
	// head := makeList()
	// operationCount := int(0)

	obj := Constructor(head)
	for i := 0; i < operationCount; i++ {
		fmt.Printf("obj.GetRandom() = %v\n", obj.GetRandom())
	}
}

