package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	queue := []*ListNode{}
	stack := []*ListNode{}
	tmp := head
	count := int(1)
	for tmp != nil {
		if left <= count && count <= right {
			stack = append(stack, tmp)
		} else {
			queue = append(queue, tmp)
		}
		tmp = tmp.Next
		count++
	}

	var result *ListNode
	var resultHead *ListNode
	for i := 1; i < count; i++ {
		var targetNode *ListNode
		if left <= i && i <= right {
			targetNode = stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
		} else {
			targetNode = queue[0]
			queue = queue[1:]
		}

		if result == nil {
			result = targetNode
			resultHead = result
		} else {
			result.Next = targetNode
			result = result.Next
		}
		result.Next = nil
	}

	return resultHead
}

func makeList1() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	return head
}

func makeList2() *ListNode {
	head := &ListNode{Val: 5}
	return head
}

func makeList() *ListNode {
	head := &ListNode{Val: 0}
	return head
}

func printList(head *ListNode) {
	if head == nil {
		fmt.Println("end")
		return
	}
	fmt.Printf("%p: %v\n", head, head)
	printList(head.Next)
}

func main() {
	// result: [1,4,3,2,5]
	// head := makeList1()
	// left := int(2)
	// right := int(4)

	// result: [5]
	head := makeList2()
	left := int(1)
	right := int(1)

	// result: 
	// head := makeList()
	// left := int(0)
	// right := int(0)

	result := reverseBetween(head, left, right)
	printList(result)
}

