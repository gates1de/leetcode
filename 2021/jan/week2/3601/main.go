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
	DL := addDigit(l1RDL, l2RDL)
	// fmt.Printf("l1RDL = %v, l2RDL = %v, DL = %v\n", l1RDL, l2RDL, DL)
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

func toDigitList(target int, result []int) []int {
	// fmt.Printf("target = %v, result = %v\n", target, result)
    if target > 0 {
        return toDigitList(target / 10, append(result, target % 10))
    }
    return result
}

func addDigit(d1 []int, d2 []int) []int {
	maxLen := len(d1)
	if len(d1) < len(d2) {
		maxLen = len(d2)
	}

	result := []int{}
	carryD := int(0)
	for i := 0; i < maxLen; i++ {
		d := int(carryD)
		if len(d1) > i {
			d += d1[len(d1) - i - 1]
		}
		if len(d2) > i {
			d += d2[len(d2) - i - 1]
		}

		if d >= 10 {
			carryD = d / 10
			d -= 10
		} else {
			carryD = 0
		}
		// fmt.Printf("d = %v, carryD = %v\n", d, carryD)
		result = append(result, d)
	}
	if carryD > 0 {
		result = append(result, carryD)
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

