package main
import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func allPossibleFBT(n int) []*TreeNode {
	memo := make(map[int][]*TreeNode)
	return helper(n, memo)
}

func helper(n int, memo map[int][]*TreeNode) []*TreeNode {
	result := make([]*TreeNode, 0)
	if n % 2 == 0 {
		return result
	}
	if n == 1 {
		result = append(result, &TreeNode{Val: 0})
		return result
	}

	if memo[n] != nil {
		return memo[n]
	}

	for i := 1; i < n; i += 2 {
		left := helper(i, memo)
		right := helper(n - i - 1, memo)

		for _, l := range left {
			for _, r := range right {
				root := &TreeNode{Val: 0, Left: l, Right: r}
				result = append(result, root)
			}
		}
	}

	memo[n] = result
	return result
}

func main() {
	// result: [[0,0,0,null,null,0,0,null,null,0,0],[0,0,0,null,null,0,0,0,0],[0,0,0,0,0,0,0],[0,0,0,0,0,null,null,null,null,0,0],[0,0,0,0,0,null,null,0,0]]
	// n := int(7)

	// result: [[0,0,0]]
	// n := int(3)

	// result: [[0]]
	n := int(1)

	// result: 
	// n := int(0)

	result := allPossibleFBT(n)
	fmt.Printf("result = %v\n", result)
}

