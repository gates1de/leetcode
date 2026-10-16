package main

import (
	"fmt"
)

const modulo = int64(1e9 + 7)

type segmentNode struct {
	sum    int64
	square int64
	lazy   int64
}

func apply(node *segmentNode, left int, right int, value int64) {
	length := int64(right - left + 1)
	node.square = (node.square + 2*value*node.sum + value*value%modulo*length) % modulo
	node.sum = (node.sum + value*length) % modulo
	node.lazy += value
}

func push(tree []segmentNode, index int, left int, right int) {
	if tree[index].lazy == 0 || left == right {
		return
	}
	middle := (left + right) / 2
	value := tree[index].lazy
	apply(&tree[index*2], left, middle, value)
	apply(&tree[index*2+1], middle+1, right, value)
	tree[index].lazy = 0
}

func update(tree []segmentNode, index int, left int, right int, queryLeft int, queryRight int) {
	if queryLeft <= left && right <= queryRight {
		apply(&tree[index], left, right, 1)
		return
	}

	push(tree, index, left, right)
	middle := (left + right) / 2
	if queryLeft <= middle {
		update(tree, index*2, left, middle, queryLeft, queryRight)
	}
	if middle < queryRight {
		update(tree, index*2+1, middle+1, right, queryLeft, queryRight)
	}
	tree[index].sum = (tree[index*2].sum + tree[index*2+1].sum) % modulo
	tree[index].square = (tree[index*2].square + tree[index*2+1].square) % modulo
}

func sumCounts(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tree := make([]segmentNode, len(nums)*4)
	last := make([]int, 100001)
	answer := int64(0)

	for right, value := range nums {
		left := last[value]
		update(tree, 1, 0, len(nums)-1, left, right)
		answer = (answer + tree[1].square) % modulo
		last[value] = right + 1
	}

	return int(answer)
}

func main() {
	// result: 15
	// nums := []int{1,2,1}

	// result: 3
	nums := []int{2, 2}

	// result: 0
	// nums := []int{}

	result := sumCounts(nums)
	fmt.Printf("result = %v\n", result)
}
