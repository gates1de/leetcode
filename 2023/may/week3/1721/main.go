package main
import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    n := int(0)
    tmp := head
    for tmp != nil {
        n++
        tmp = tmp.Next
    }

    if k > n / 2 {
        k = n - k
    } else {
        k--
    }

    nodes := make([]*ListNode, n)
    i := int(0)
    for head != nil {
        nodes[i] = head
        i++
        head = head.Next
    }

    if n == 2 {
        nodes[0], nodes[1] = nodes[1], nodes[0]
        nodes[0].Next = nodes[1]
        nodes[1].Next = nil
        return nodes[0]
    }

    k1 := nodes[k]
    k2 := nodes[n - 1 - k]
    k1Next := k1.Next
    k2Next := k2.Next
    var result *ListNode
    if k == 0 {
        result = k2
        result.Next = nodes[1]
        k1.Next = nil
        nodes[n - 2].Next = k1
    } else {
        result = nodes[0]
        nodes[k - 1].Next = k2
        if n - 2 - k == k {
            k2.Next = k1
        } else {
            k2.Next = k1Next
        }
        nodes[n - 2 - k].Next = k1
        k1.Next = k2Next
    }

    return result
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
	head := &ListNode{Val: 7}
	head.Next = &ListNode{Val: 9}
	head.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next.Next = &ListNode{Val: 7}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 8}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 0}
	head.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 9}
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 5}
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
	// result: [1,4,3,2,5]
	// head := makeList1()
	// k := int(2)

	// result: [7,9,6,6,8,7,3,0,9,5]
	head := makeList2()
	k := int(5)

	// result: 
	// head := makeList()
	// k := int(0)

	result := swapNodes(head, k)
	printList(result)
}

