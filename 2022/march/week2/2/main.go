package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	nums := make([]int, 101)
	length := int(0)
	for l1 != nil {
		nums[length] += l1.Val
		l1 = l1.Next
		length++
	}

	maxLen := length
	length = 0
	for l2 != nil {
		nums[length] += l2.Val
		l2 = l2.Next
		length++
	}

	if length > maxLen {
		maxLen = length
	}
	nums = nums[:maxLen]
	// fmt.Println(nums)

	var result *ListNode
	isAdvance := false
	for _, num := range nums {
		if isAdvance {
			num++
			isAdvance = false
		}
		if num >= 10 {
			isAdvance = true
			num -= 10
		}
		result = insert(result, num)
	}

	if isAdvance {
		result = insert(result, 1)
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

func makeLists1() (*ListNode, *ListNode) {
	node1 := &ListNode{Val: 2}
	node2 := &ListNode{Val: 4}
	node3 := &ListNode{Val: 3}
	node1.Next = node2
	node1.Next.Next = node3

	node4 := &ListNode{Val: 5}
	node5 := &ListNode{Val: 6}
	node6 := &ListNode{Val: 4}
	node4.Next = node5
	node4.Next.Next = node6

	return node1, node4
}

func makeLists2() (*ListNode, *ListNode) {
	return &ListNode{Val: 0}, &ListNode{Val: 0}
}

func makeLists3() (*ListNode, *ListNode) {
	node1 := &ListNode{Val: 9}
	node2 := &ListNode{Val: 9}
	node3 := &ListNode{Val: 9}
	node4 := &ListNode{Val: 9}
	node5 := &ListNode{Val: 9}
	node6 := &ListNode{Val: 9}
	node7 := &ListNode{Val: 9}
	node1.Next = node2
	node1.Next.Next = node3
	node1.Next.Next.Next = node4
	node1.Next.Next.Next.Next = node5
	node1.Next.Next.Next.Next.Next = node6
	node1.Next.Next.Next.Next.Next.Next = node7

	node8 := &ListNode{Val: 9}
	node9 := &ListNode{Val: 9}
	node10 := &ListNode{Val: 9}
	node11 := &ListNode{Val: 9}
	node8.Next = node9
	node8.Next.Next = node10
	node8.Next.Next.Next = node11

	return node1, node8
}

func printList(node *ListNode) {
	if node == nil {
		fmt.Println("end")
		return
	}
	fmt.Println(node)
	printList(node.Next)
}

func main() {
	// result: [7,0,8]
	// l1, l2 := makeLists1()

	// result: [0]
	// l1, l2 := makeLists2()

	// result: [8,9,9,9,0,0,0,1]
	l1, l2 := makeLists3()

	// result: 
	// l1, l2 := makeLists()

	result := addTwoNumbers(l1, l2)
	printList(result)
}

