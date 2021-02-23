package main
import (
	"./treenode"
	"fmt"
)

func trimBST(root *treenode.TreeNode, low int, high int) *treenode.TreeNode {
	fmt.Printf("root = %v\n", root)
	root = traversalTrim(root, low, high)
	return root
}

func traversalTrim(target *treenode.TreeNode, low int, high int) *treenode.TreeNode {
	fmt.Printf("target = %v\n", target)
	if target == nil {
		return nil
	}

	if target.Val < low || high < target.Val {
		if target.Left == nil && target.Right == nil {
			return nil
		} else if target.Left == nil {
			target = target.Right
		} else if target.Right == nil {
			target = target.Left
		} else {
			minNode := getMinNode(target.Right)
			target.Val = minNode.Val
			minNode.Val = -1
			target.Right = traversalTrim(target.Right, low, high)
		}
		target = traversalTrim(target, low, high)
		return target
	}
	target.Left = traversalTrim(target.Left, low, high)
	target.Right = traversalTrim(target.Right, low, high)
	return target
}

func getMinNode(target *treenode.TreeNode) *treenode.TreeNode {
	current := target
	for current.Left != nil {
		current = current.Left
	}
	// fmt.Printf("getMinNode = %v\n", current)
	return current
}

func printNode(root *treenode.TreeNode, index int) {
	if root == nil {
		return
	}
	fmt.Printf("%v: node = %v\n", index, root)
	printNode(root.Left, index * 2)
    printNode(root.Right, index * 2 + 1)
}

func main() {
	// rootNums := []int{1, 0, 2}
	// low := int(1)
	// high := int(2)

	// rootNums := []int{3, 0, 4, -1, 2, -1, -1, 1}
	// low := int(1)
	// high := int(3)

	// rootNums := []int{1}
	// low := int(1)
	// high := int(2)

	// rootNums := []int{1, -1, 2}
	// low := int(2)
	// high := int(4)

	// rootNums := []int{2, 1, 3}
	// low := int(3)
	// high := int(4)

	rootNums := []int{41,37,44,24,39,42,48,1,35,38,40,-1,43,46,49,0,2,30,36,-1,-1,-1,-1,-1,-1,45,47,-1,-1,-1,-1,-1,4,29,32,-1,-1,-1,-1,-1,-1,3,9,26,-1,31,34,-1,-1,7,11,25,27,-1,-1,33,-1,6,8,10,16,-1,-1,-1,28,-1,-1,5,-1,-1,-1,-1,-1,15,19,-1,-1,-1,-1,12,-1,18,20,-1,13,17,-1,-1,22,-1,14,-1,-1,21,23}
	low := int(25)
	high := int(55)

	root := treenode.Maketree(rootNums)
	// printNode(root, 1)
	result := trimBST(root, low, high)
	printNode(result, 1)
	// result := trimBST(root, low, high)
	// fmt.Printf("result = %v\n", result)
}

