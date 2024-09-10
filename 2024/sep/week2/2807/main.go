package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	node1 := head
	node2 := head.Next

	for node2 != nil {
		node := &ListNode{Val: gcd(node1.Val, node2.Val)}
		node1.Next = node
		node.Next = node2

		node1 = node2
		node2 = node2.Next
	}

	return head
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a % b)
}

func makeList1() *ListNode {
	head := &ListNode{Val: 18}
	head.Next = &ListNode{Val: 6}
	head.Next.Next = &ListNode{Val: 10}
	head.Next.Next.Next = &ListNode{Val: 3}
	return head
}

func makeList2() *ListNode {
	head := &ListNode{Val: 7}
	return head
}

func makeList() *ListNode {
	var head *ListNode
	return head
}

func printList(head *ListNode) {
	if head == nil {
		return
	}

	fmt.Println(head)
	printList(head.Next)
}

func main() {
	// result: [18,6,6,2,10,1,3]
	// head := makeList1()

	// result: [7]
	head := makeList2()

	// result: []
	// head := makeList()

	result := insertGreatestCommonDivisors(head)
	printList(result)
}

