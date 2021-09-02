package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

// Not my solution (REF: https://leetcode.com/problems/unique-binary-search-trees-ii/discuss/523744/Go-golang-0ms-solution)
func generateTrees(n int) []*TreeNode {
    if n == 0 {
		return []*TreeNode{}
	}
    return helper(1, n)
}

func helper(start int, end int) []*TreeNode {
    if start > end {
		return []*TreeNode{nil}
	}

    result := []*TreeNode{}
    for s := start; s <= end; s++ {
        left := helper(start, s - 1)
        right := helper(s + 1, end)
        for _, l := range left {
            for _, r := range right {
				node := &TreeNode{Val: s, Left: l, Right: r}
                result = append(result, node)
            }
        }
    }
    return result
}


func main() {
	// result: [[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
	n := int(3)

	// result: [[1]]
	// n := int(1)

	// result: 
	// n := int(0)

	// result: 
	// n := int(0)

	// result: 
	// n := int(0)

	result := generateTrees(n)
	fmt.Printf("result = %v\n", result)
}

