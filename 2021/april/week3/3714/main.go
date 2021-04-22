package main
import (
	"fmt"
)

type Node struct {
    Val int
    Children []*Node
}

func preorder(root *Node) []int {
	result := []int{}
	helper(root, &result)
	return result
}

func makeNode(nums []int, current *Node) *Node {
    if len(nums) == 0 {
		return nil
    }
    parent := &Node{Val: nums[0]}
	nodes := []*Node{}
	parentIndex := int(0)
	nodes = append(nodes, parent)
	for _, val := range nums[2:] {
		if val < 0 {
			parentIndex++
			continue
		}
		node := insert(val, parentIndex, nodes)
		if node != nil {
			nodes = append(nodes, node)
		}
    }
    return parent
}

func insert(val int, parentIndex int, nodes []*Node) *Node {
    if val < 0 {
		return nil
    }
	parent := nodes[parentIndex]
	node := &Node{Val: val}
	parent.Children = append(parent.Children, node)
	return node
}

func helper(root *Node, result *[]int) {
	if root == nil {
		return
	}
	*result = append(*result, root.Val)
	for _, child := range root.Children {
		helper(child, result)
	}
}

func main() {
	// result: [1, 3, 5, 6, 2, 4]
	// rootNums := []int{1, -1, 3, 2, 4, -1, 5, 6}

	// result: [1,2,3,6,7,11,14,4,8,12,5,9,13,10]
	rootNums := []int{1,-1,2,3,4,5,-1,-1,6,7,-1,8,-1,9,10,-1,-1,11,-1,12,-1,13,-1,-1,14}

	// result: []
	// rootNums := []int{}

	var root *Node
	root = makeNode(rootNums, root)
	result := preorder(root)
	fmt.Printf("result = %v\n", result)
}

