package main
import (
	"fmt"
	treenode "treenode/treenode"
)

func lowestCommonAncestor(root, p, q *treenode.TreeNode) *treenode.TreeNode {
	// printNode(root)
	m := map[*treenode.TreeNode]*treenode.TreeNode{}
	var pRoot *treenode.TreeNode
	var qRoot *treenode.TreeNode
	saveParent(root, m, p, q, &pRoot, &qRoot)
	// fmt.Printf("pRoot = %v, qRoot = %v\n", pRoot, qRoot)

	pParents := []*treenode.TreeNode{}
	qParents := []*treenode.TreeNode{}
	for pRoot != nil {
		pParents = append(pParents, pRoot)
		pRoot = m[pRoot]
	}
	for qRoot != nil {
		qParents = append(qParents, qRoot)
		qRoot = m[qRoot]
	}
	// fmt.Printf("pParents = %v, qParents = %v\n", pParents, qParents)
	for _, pNode := range pParents {
		for _, qNode := range qParents {
			if pNode == qNode {
				return pNode
			}
		}
	}

	return root
	// return traverse(root, m, p.Val, q.Val)
}

func saveParent(
	root *treenode.TreeNode,
	m map[*treenode.TreeNode]*treenode.TreeNode,
	p *treenode.TreeNode,
	q *treenode.TreeNode,
	pRoot **treenode.TreeNode,
	qRoot **treenode.TreeNode,
) {
    if root == nil {
        return
    }

	if root.Val == p.Val {
		*pRoot = root
	}
	if root.Val == q.Val {
		*qRoot = root
	}

	if root.Left != nil {
		m[root.Left] = root
		saveParent(root.Left, m, p, q, pRoot, qRoot)
	}
	if root.Right != nil {
		m[root.Right] = root
		saveParent(root.Right, m, p, q, pRoot, qRoot)
	}
}

func printNode(root *treenode.TreeNode) {
    if root == nil {
        return
    }
    fmt.Printf("node %p = %v\n", root, root)
    printNode(root.Left)
    printNode(root.Right)
}

func main() {
	// result: 3
	// rootNums := []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}
	// p := &treenode.TreeNode{Val: 5}
	// q := &treenode.TreeNode{Val: 1}

	// result: 5
	// rootNums := []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}
	// p := &treenode.TreeNode{Val: 5}
	// q := &treenode.TreeNode{Val: 4}

	// result: 1
	rootNums := []int{1, 2}
	p := &treenode.TreeNode{Val: 1}
	q := &treenode.TreeNode{Val: 2}

	// result: 
	// rootNums := []int{}
	// p := &treenode.TreeNode{Val: }
	// q := &treenode.TreeNode{Val: }

	root := treenode.Maketree(rootNums)
	result := lowestCommonAncestor(root, p, q)
	fmt.Printf("result = %v\n", result)
}

