package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func makeListNode(s []int, current *ListNode) *ListNode {
    if len(s) == 0 {
        if current == nil {
            current = &ListNode{Val: 0}
        }
        return current
    }

    if current == nil {
        current = &ListNode{Val: s[0]}
        return makeListNode(s[1:], current)
    }
    current.Next = makeListNode(s[1:], &ListNode{Val: s[0]})
    return current
}

func isPalindrome(head *ListNode) bool {
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}

	i := int(0)
	j := int(len(nums) - 1)
	for i <= j {
		if nums[i] != nums[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	// headNums := []int{1, 2, 2, 1}
	headNums := []int{1, 2, 3, 2, 2}

	var head *ListNode
	head = makeListNode(headNums, head)

	result := isPalindrome(head)
	fmt.Printf("result = %v\n", result)
}

