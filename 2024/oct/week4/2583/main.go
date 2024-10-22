package main
import (
	"fmt"
	heap "container/heap"
)

type Heap []int

func (h Heap) Len() int { return len(h) }

func (h Heap) Less(a, b int) bool {
	return h[a] > h[b]
}

func (h Heap) Swap(a, b int) { h[a], h[b] = h[b], h[a] }

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n - 1]
	*h = old[0 : n - 1]
	return x
}

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func kthLargestLevelSum(root *TreeNode, k int) int64 {
	levels := make(map[*TreeNode]int)
	level := int(1)
	levels[root] = level
	nodes := make([]*TreeNode, 0)
	nodes = append(nodes, root)

	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		level = levels[node]

		if node.Left != nil {
			nodes = append(nodes, node.Left)
			levels[node.Left] = level + 1
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
			levels[node.Right] = level + 1
		}
	}

	if level < k {
		return -1
	}

	sums := make([]int, level + 1)
	for node, level := range levels {
		sums[level] += node.Val
	}

	priorityQueue := &Heap{}
	heap.Init(priorityQueue)

	for _, sum := range sums {
		heap.Push(priorityQueue, sum)
	}

	result := int64(0)
	for i := 0; i < k; i++ {
		result = int64(heap.Pop(priorityQueue).(int))
	}

	return result
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 8}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Right = &TreeNode{Val: 1}
	root.Left.Left.Left = &TreeNode{Val: 4}
	root.Left.Left.Right = &TreeNode{Val: 6}
	root.Right = &TreeNode{Val: 9}
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 7}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 3}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func main() {
	// resutl: 13
	// root := makeTree1()
	// k := int(2)

	// resutl: 3
	root := makeTree2()
	k := int(1)

	// resutl: 
	// root := makeTree()
	// k := int(0)

	result := kthLargestLevelSum(root, k)
	fmt.Printf("result = %v\n", result)
}

