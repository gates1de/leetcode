package main
import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

var dir = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func spiralMatrix(m int, n int, head *ListNode) [][]int {
	cur := int(0)

	result := make([][]int, m)
	for i, _ := range result {
		result[i] = make([]int, n)
		for j, _ := range result[i] {
			result[i][j] = -1
		}
	}

	i := int(0)
	j := int(0)

	for head != nil {
		result[i][j] = head.Val
		newI := i + dir[cur][0]
		newJ := j + dir[cur][1]

		if min(newI, newJ) < 0 || newI >= m || newJ >= n || result[newI][newJ] != -1 {
			cur = (cur + 1) % 4
		}

		i += dir[cur][0]
		j += dir[cur][1]

		head = head.Next
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
	head.Next = &ListNode{Val: 0}
	head.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next = &ListNode{Val: 6}
	head.Next.Next.Next.Next = &ListNode{Val: 8}
	head.Next.Next.Next.Next.Next = &ListNode{Val: 1}
	head.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}
	head.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 9}
	head.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 2}
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 5}
	head.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next.Next = &ListNode{Val: 0}
	return head
}

func makeList2() *ListNode {
	head := &ListNode{Val: 0}
	head.Next = &ListNode{Val: 1}
	head.Next.Next = &ListNode{Val: 2}
	return head
}

func makeList() *ListNode {
	var head *ListNode
	return head
}

func main() {
	// result: [[3,0,2,6,8],[5,0,-1,-1,1],[5,2,4,9,7]]
	// m := int(3)
	// n := int(5)
	// head := makeList1()

	// result: [[0,1,2,-1]]
	m := int(1)
	n := int(4)
	head := makeList2()

	// result: []
	// m := int(0)
	// n := int(0)
	// head := makeList()

	result := spiralMatrix(m, n, head)
	fmt.Printf("result = %v\n", result)
}

