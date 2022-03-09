package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	m := map[int]int{}
	values := []int{}
	for head != nil {
		values = append(values, head.Val)
		m[head.Val]++
		head = head.Next
	}

	var result *ListNode
	for _, v := range values {
		if m[v] >= 2 {
			continue
		}
		result = insert(result, v)
	}

	return result
}

func insert(head *ListNode, target int) *ListNode {
	if head == nil {
		return &ListNode{Val: target}
	}

	last := head
	for last.Next != nil {
		last = last.Next
	}
	last.Next = &ListNode{Val: target}

	return head
}

func makeList1() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 5}
	return head
}

func makeList2() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next = &ListNode{Val: 3}
	return head
}

func makeList3() *ListNode {
	head := &ListNode{Val: 1}
	return head
}

func makeList4() *ListNode {
	return nil
}

func printList(head *ListNode) {
	if head == nil {
		fmt.Println("end")
		return
	}
	fmt.Println(head)
	printList(head.Next)
}

func main() {
	// result: [1,2,5]
	// head := makeList1()

	// result: [2,3]
	// head := makeList2()

	// result: [1]
	// head := makeList3()

	// result: []
	head := makeList4()

	// result: 
	// head := makeList()

	result := deleteDuplicates(head)
	printList(result)
}

