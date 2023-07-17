package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Nums := make([]int, 0, 100)
	l2Nums := make([]int, 0, 100)

	temp := l1
	for temp != nil {
		l1Nums = append(l1Nums, 0)
		copy(l1Nums[1:], l1Nums)
		l1Nums[0] = temp.Val
		temp = temp.Next
	}

	temp = l2
	for temp != nil {
		l2Nums = append(l2Nums, 0)
		copy(l2Nums[1:], l2Nums)
		l2Nums[0] = temp.Val
		temp = temp.Next
	}

	if len(l1Nums) < len(l2Nums) {
		l1Nums, l2Nums = l2Nums, l1Nums
	}

	for i, num1 := range l1Nums {
		num := num1
		if i < len(l2Nums) {
			num += l2Nums[i]
		}

		if num >= 10 {
			if i + 1 >= len(l1Nums) {
				l1Nums = append(l1Nums, 1)
			} else {
				l1Nums[i + 1] += 1
			}
			num -= 10
		}
		l1Nums[i] = num
	}

	var result *ListNode
	for i := len(l1Nums) - 1; i >= 0; i-- {
		num := l1Nums[i]
		result = addNode(result, num)
	}
	return result
}

func addNode(node *ListNode, target int) *ListNode {
	if node == nil {
		return &ListNode{Val: target}
	}

	node.Next = addNode(node.Next, target)
	return node
}

func makeList1() (*ListNode, *ListNode) {
	l1 := &ListNode{Val: 7}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 4}
	l1.Next.Next.Next = &ListNode{Val: 3}

	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}

	return l1, l2
}

func makeList2() (*ListNode, *ListNode) {
	l1 := &ListNode{Val: 2}
	l1.Next = &ListNode{Val: 4}
	l1.Next.Next = &ListNode{Val: 3}

	l2 := &ListNode{Val: 5}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}

	return l1, l2
}

func makeList3() (*ListNode, *ListNode) {
	l1 := &ListNode{Val: 0}

	l2 := &ListNode{Val: 0}
	return l1, l2
}

func makeList4() (*ListNode, *ListNode) {
	l1 := &ListNode{Val: 9}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 4}
	l1.Next.Next.Next = &ListNode{Val: 3}

	l2 := &ListNode{Val: 8}
	l2.Next = &ListNode{Val: 6}
	l2.Next.Next = &ListNode{Val: 4}

	return l1, l2
}

func makeList() (*ListNode, *ListNode) {
	var l1, l2 *ListNode
	return l1, l2
}

func printList(node *ListNode) {
	if node == nil {
		return
	}

	fmt.Println(node)
	printList(node.Next)
}

func main() {
	// result: [7,8,0,7]
	// l1, l2 := makeList1()
	// result := addTwoNumbers(l1, l2)

	// result: [8,0,7]
	// l1, l2 := makeList2()
	// result := addTwoNumbers(l1, l2)

	// result: [0]
	// l1, l2 := makeList3()
	// result := addTwoNumbers(l1, l2)

	// result: [1,0,1,0,7]
	l1, l2 := makeList4()
	result := addTwoNumbers(l1, l2)

	// result: 
	// l1, l2 := makeList()
	// result := addTwoNumbers(l1, l2)

	printList(result)
}

