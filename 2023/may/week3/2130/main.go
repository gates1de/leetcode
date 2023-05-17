package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func pairSum(head *ListNode) int {
	nums := make([]int, 0, 100000)
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	result := int(0)
	i := int(0)
	j := len(nums) - 1

	for i <= j {
		if i == j {
			result = max(result, nums[i])
		} else {
			result = max(result, nums[i] + nums[j])
		}
		i++
		j--
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func makeList1() *ListNode {
	head := &ListNode{Val: 5}
	head.Next = &ListNode{Val: 4}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 1}
	return head
}

func makeList2() *ListNode {
	head := &ListNode{Val: 4}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 3}
	return head
}

func makeList3() *ListNode {
	head := &ListNode{Val: 1}
	head.Next = &ListNode{Val: 100000}
	return head
}

func makeList() *ListNode {
	var head *ListNode
	return head
}

func main() {
	// result: 6
	// head := makeList1()

	// result: 7
	// head := makeList2()

	// result: 100001
	head := makeList3()

	// result: 
	// head := makeList()

	result := pairSum(head)
	fmt.Printf("result = %v\n", result)
}

