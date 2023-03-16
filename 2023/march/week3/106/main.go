package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(inorder) == 0 {
        return nil
    }

    root := &TreeNode{Val: postorder[len(postorder) - 1]}

    if len(inorder) == 1 {
        return root
    }

    inRootIndex := int(-1)
    for i, num := range inorder {
        if num == root.Val {
            inRootIndex = i
            break
        }
    }

    if inRootIndex == -1 {
        return root
    }

    inLefts := inorder[:inRootIndex]
    inRights := inorder[inRootIndex + 1:]

    postLefts := postorder[:len(inLefts)]
    postRights := postorder[len(inLefts):len(postorder) - 1]

    root.Left = buildTree(inLefts, postLefts)
    root.Right = buildTree(inRights, postRights)

    return root
}

func main() {
	// result: [3,9,20,null,null,15,7]
	// inorder := []int{9,3,15,20,7}
	// postorder := []int{9,15,7,20,3}

	// result: [-1]
	inorder := []int{-1}
	postorder := []int{-1}

	// result: 
	// inorder := []int{}
	// postorder := []int{}

	result := buildTree(inorder, postorder)
	fmt.Printf("result = %v\n", result)
}

