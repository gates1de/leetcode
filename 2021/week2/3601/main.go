package main
import (
	"fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return &ListNode{Val: 0}
	}
	l1RDL := reverseDigitList(l1, []int{})
	l2RDL := reverseDigitList(l2, []int{})
	sum := digitToInt(l1RDL) + digitToInt(l2RDL)
	DL := toDigitList(sum, []int{})
	// fmt.Printf("l1RDL = %v, l2RDL = %v, sum = %v, DL = %v\n", l1RDL, l2RDL, sum, DL)
	var result *ListNode
	result = makeListNode(DL, result)
	return result
}

func reverseDigitList(l *ListNode, result []int) []int {
	if l.Next != nil {
		result = append(result, reverseDigitList(l.Next, result)...)
	}
	result = append(result, l.Val)
	return result
}

func toDigitList(target int64, result []int) []int {
	// fmt.Printf("target = %v, result = %v\n", target, result)
    if target > 0 {
        return toDigitList(target / 10, append(result, int(target % 10)))
    }
    return result
}

func digitToInt(target []int) int64 {
	i := int64(1)
	result := int64(0)
	for j, _ := range target {
		result += int64(target[len(target) - j - 1]) * i
		i *= 10
	}
	return result
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

func main() {
	// l1Arr := []int{2, 4, 3}
	// l2Arr := []int{5, 6, 4}

	// l1Arr := []int{0}
	// l2Arr := []int{0}

	// l1Arr := []int{9, 9, 9, 9, 9, 9, 9}
	// l2Arr := []int{9, 9, 9, 9}

	l1Arr := []int{1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1}
	l2Arr := []int{5, 6, 4}

	var l1, l2 *ListNode
	l1 = makeListNode(l1Arr, l1)
	l2 = makeListNode(l2Arr, l2)
	result := addTwoNumbers(l1, l2)
	fmt.Printf("result = %v\n", result)
}

