package main
import (
	"fmt"
)


type ListNode struct {
    Val int
    Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	i := int(1)
	nums := []int{}
	reverseTargets := make([]int, right - left + 1)
	for head != nil {
		if i < left || right < i {
			nums = append(nums, head.Val)
		} else if left <= i && i <= right {
			reverseTargets[i - left] = head.Val
			if i == right {
				reverse(reverseTargets)
				nums = append(nums, reverseTargets...)
			}
		}
		head = head.Next
		i++
	}

	var result *ListNode
	result = makeListNode(nums, result)
	return result
}

func reverse(nums []int) {
	for i := 0; i < len(nums) / 2; i++ {
		nums[i], nums[len(nums) - i - 1] = nums[len(nums) - i - 1], nums[i]
	}
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

func printListNode(l *ListNode) {
    if l == nil {
        return
    }
    fmt.Printf("l = %v\n", l)
    printListNode(l.Next)
}

func main() {
	// result: [1,4,3,2,5]
	// left := int(2)
	// right := int(4)
	// nums := []int{1, 2, 3, 4, 5}

	// result: [5]
	left := int(1)
	right := int(1)
	nums := []int{5}

	// result: 
	// left := int(0)
	// right := int(0)
	// nums := []int{}

	var head *ListNode
	head = makeListNode(nums, head)
	result := reverseBetween(head, left, right)
	printListNode(result)
}

