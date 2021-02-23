package treenode

// import "fmt"

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func Maketree(list []int) *TreeNode {
    if len(list) == 0 {
		return nil
	}

	parent := &TreeNode{Val: list[0]}
	for i, val := range list[1:] {
        insert(i + 1, parent, val)
	}
    return parent
}


func insert(i int, parent *TreeNode, val int) bool {
	// fmt.Printf("%v: parent = %v, val = %v\n", i, parent, val)
    if parent == nil || val < 0 {
		return false
	}

	if val < parent.Val {
		if parent.Left != nil {
			return insert(i, parent.Left, val)
		}
		parent.Left = &TreeNode{Val: val}
		return true
	}

	if parent.Right != nil {
		return insert(i, parent.Right, val)
	}
	parent.Right = &TreeNode{Val: val}
    return true
}

