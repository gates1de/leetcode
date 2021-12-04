package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
		return head
	}

    odd := head
    even := head.Next
    evenHead := head.Next
    for even != nil && even.Next != nil {
        odd.Next = even.Next
        odd = odd.Next
        even.Next = odd.Next
        even = even.Next
    }
    odd.Next = evenHead
    return head
}

// Accepted, but not O(1) extra space complexity
func mySolution(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	odds := []*ListNode{}
	evens := []*ListNode{}
	i := int(1)
	for head != nil {
		if i % 2 == 0 {
			evens = append(evens, &ListNode{Val: head.Val})
		} else {
			odds = append(odds, &ListNode{Val: head.Val})
		}
		i++
		head = head.Next
	}

	result := odds[0]
	for _, node := range odds[1:] {
		add(result, node)
	}
	for _, node := range evens {
		add(result, node)
	}
	return result
}

func add(head *ListNode, target *ListNode) {
	if head == nil {
		head = target
	}

	for head.Next != nil {
		head = head.Next
	}
	head.Next = target
}

func makeList1() *ListNode {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 2}
	root.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next = &ListNode{Val: 4}
	root.Next.Next.Next.Next = &ListNode{Val: 5}
	return root
}

func makeList2() *ListNode {
	root := &ListNode{Val: 2}
	root.Next = &ListNode{Val: 1}
	root.Next.Next = &ListNode{Val: 3}
	root.Next.Next.Next = &ListNode{Val: 5}
	root.Next.Next.Next.Next = &ListNode{Val: 6}
	root.Next.Next.Next.Next.Next = &ListNode{Val: 4}
	root.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}
	return root
}

func makeList3() *ListNode {
	var root *ListNode
	return root
}

func makeList4() *ListNode {
	root := &ListNode{Val: 1}
	root.Next = &ListNode{Val: 7}
	root.Next.Next = &ListNode{Val: 7}
	root.Next.Next.Next = &ListNode{Val: 7}
	return root
}

func printNode(head *ListNode) {
	if head == nil {
		return
	}
	fmt.Printf("val = %v\n", head.Val)
	printNode(head.Next)
}

func main() {
	// result: [1,3,5,2,4]
	// head := makeList1()

	// result: [2,3,6,7,1,5,4]
	// head := makeList2()

	// result: []
	// head := makeList3()

	// result: [1,7,7,7]
	head := makeList4()

	// result: 
	// head := makeList()

	result := oddEvenList(head)
	printNode(result)
}

