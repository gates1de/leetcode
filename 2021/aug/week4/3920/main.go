package main
import (
	"fmt"
	"strings"
)

type TreeNode struct {
    Val string
    Left *TreeNode
    Right *TreeNode
}

func isValidSerialization(preorder string) bool {
	values := strings.Split(preorder, ",")
	if len(values) == 0 {
		return false
	} else if len(values) == 1 {
		return values[0] == "#"
	}

	var root *TreeNode
	helper(root, &values)
	return len(values) == 0
}

func helper(root *TreeNode, values *[]string) {
	if len(*values) == 0 {
		return
	}

	value := (*values)[0]
	// fmt.Printf("value = %v\n", value)
	if value != "#" {
		root = &TreeNode{Val: value}
		// fmt.Printf("root = %v\n", root)
	}
	*values = (*values)[1:]

	if root != nil {
		helper(root.Left, values)
		if len(*values) == 0 {
			*values = append(*values, " ")
			return
		}
		helper(root.Right, values)
	}
}

func main() {
	// result: true
	// preorder := "9,3,4,#,#,1,#,#,2,#,6,#,#"

	// result: false
	// preorder := "1,#"

	// result: false
	// preorder := "9,#,#,1"

	// result: true
	// preorder := "#"

	// result: false
	// preorder := "1,1,#"

	// result: false
	preorder := "91,13,14,#,#,10"

	// result: 
	// preorder := ""

	result := isValidSerialization(preorder)
	fmt.Printf("result = %v\n", result)
}

