package treenode

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
	for i, val := range list {
        insert(parent, i + 1, 1, val)
	}
    return parent
}


func insert(parent *TreeNode, i int, parentIndex int, val int) bool {
    if parent == nil || val < -10000 {
		return false
	}

    if parentIndex * 2 == i {
        if parent.Left == nil {
			parent.Left = &TreeNode{Val: val}
            return true
		}
	} else if parentIndex * 2 + 1 == i {
        if parent.Right == nil {
			parent.Right = &TreeNode{Val: val}
            return true
		}
	}

    if parentIndex * 2 > i {
        return false
	}

	isInserted := insert(parent.Left, i, parentIndex * 2, val)
    if isInserted {
        return true
	}

    return insert(parent.Right, i, parentIndex * 2 + 1, val)
}

