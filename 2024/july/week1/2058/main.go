package main
import (
	"fmt"
	"math"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	result := []int{-1, -1}
	minDistance := math.MaxInt32

	prev := head
	current := head.Next
	i := int(1)
	firstCriticalIndex := int(0)
	prevCriticalIndex := int(0)

	for current.Next != nil {
		if current.Val < prev.Val && current.Val < current.Next.Val || 
		  current.Val > prev.Val && current.Val > current.Next.Val {
			  if prevCriticalIndex == 0 {
				  prevCriticalIndex = i
				  firstCriticalIndex = i
			  } else {
				  minDistance = min(minDistance, i - prevCriticalIndex)
				  prevCriticalIndex = i
			  }

			  prevCriticalIndex = i
		  }

		  i++
		  prev = current
		  current = current.Next
	}

	if minDistance != math.MaxInt32 {
		maxDistance := prevCriticalIndex - firstCriticalIndex
		result = []int{minDistance, maxDistance}
	}

	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func makeList1() *ListNode {
	head := &ListNode{Val: 3}
	head.Next = &ListNode{Val: 1}
	return head
}

func makeList2() *ListNode {
	head := &ListNode{Val: 5}
	head.Next = &ListNode{Val: 3}
	head.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	return head
}

func makeList3() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 3}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}
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
	// result: [-1,-1]
	// head := makeList1()

	// result: [1,3]
	// head := makeList2()

	// result: [3,3]
	head := makeList3()

	// result: []
	// head := makeList()

	result := nodesBetweenCriticalPoints(head)
	fmt.Printf("result = %v\n", result)
}

