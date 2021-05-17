package main
import (
	"fmt"
    treenode "treenode/treenode"
)

func minCameraCover(root *treenode.TreeNode) int {
	result := int(0)
	covered := map[*treenode.TreeNode]bool{}
    covered[nil] = true

    dfs(root, nil, covered, &result)

    return result
}

func dfs(node, parent *treenode.TreeNode, covered map[*treenode.TreeNode]bool, result *int) {
    if node != nil {
        dfs(node.Left, node, covered, result)
        dfs(node.Right, node, covered, result)

        if (parent == nil && !covered[node]) ||
            !covered[node.Left] ||
            !covered[node.Right] {
            *result++
            covered[node] = true
            covered[parent] = true
            covered[node.Left] = true
            covered[node.Right] = true
        }
    }
}

// Wrong Answer
func ngSolution(root *treenode.TreeNode) int {
    if root == nil {
        return 0
    }
    result := int(0)
    passed := []*treenode.TreeNode{}
    traversal(root, nil, passed, &result)
    return result
}

func traversal(root *treenode.TreeNode, parent *treenode.TreeNode, passed []*treenode.TreeNode, result *int) {
    if root == nil {
        return
    }

    if root.Left != nil && root.Right != nil {
        if parent != nil && !contains(parent, passed) {
            fmt.Printf("set camera on root (no parent, exists children): %v %p\n", root, root)
            *result++
            passed = append(passed, parent)
            passed = append(passed, root)
            passed = append(passed, root.Left)
            passed = append(passed, root.Right)
        } else {
            if !hasChild(root.Left) || !hasChild(root.Right) {
                fmt.Printf("set camera on root (exists children): %v %p\n", root, root)
                *result++
                passed = append(passed, root)
                if root.Left != nil {
                    passed = append(passed, root.Left)
                }
                if root.Right != nil {
                    passed = append(passed, root.Right)
                }
            }
        }
    } else if parent != nil && !contains(parent, passed) {
        fmt.Printf("set camera on root: %v %p\n", root, root)
        *result++
        passed = append(passed, parent)
        passed = append(passed, root)
        if root.Left != nil {
            passed = append(passed, root.Left)
        }
        if root.Right != nil {
            passed = append(passed, root.Right)
        }
    } else if !contains(root, passed) {
        if root.Left == nil && root.Right == nil {
            fmt.Printf("set camera on root 2: %v %p\n", root, root)
            *result++
            passed = append(passed, root)
            return
        }
    }

    traversal(root.Left, root, passed, result)
    traversal(root.Right, root, passed, result)
}

func contains(root *treenode.TreeNode, passed []*treenode.TreeNode) bool {
    for _, node := range passed {
        if root == node {
            return true
        }
    }
    return false
}

func hasChild(root *treenode.TreeNode) bool {
    if root == nil {
        return false
    }
    return root.Left != nil || root.Right != nil
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
	// result: 1
	// rootNums := []int{0, 0, -1, 0, 0}

	// result: 2
	// rootNums := []int{0, 0, -1, 0, -1, 0, -1, -1, 0}

	// result: 1
	// rootNums := []int{0, 0, 0}

	// result: 2
	// rootNums := []int{0, 0, 0, -1, -1, -1, 0}

	// result: 2
	rootNums := []int{0, 0, -1, 0, -1, 0, -1, -1, 0}

	// result: 
	// rootNums := []int{}

	root := treenode.Maketree(rootNums)
	result := minCameraCover(root)
	fmt.Printf("result = %v\n", result)
}

