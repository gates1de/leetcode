package main
import (
	"fmt"
	treenode "treenode/treenode"
)

type LevelQueue struct {
	Level int
	Node *treenode.TreeNode
}

func levelOrder(root *treenode.TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	nodes := bfs(root)
	maxLevel := int(0)
	for level := range nodes {
		if maxLevel < level {
			maxLevel = level
		}
	}
	for i := 0; i <= maxLevel; i++ {
		levelNodes := nodes[i]
		var levelNodeVals []int
		for _, node := range levelNodes {
			levelNodeVals = append(levelNodeVals, node.Val)
		}
		result = append(result, levelNodeVals)
	}
	return result
}

func bfs(root *treenode.TreeNode) map[int][]*treenode.TreeNode {
	result := map[int][]*treenode.TreeNode{}
	if root == nil {
		return result
	}

	first := &LevelQueue{Level: 0, Node: root}
	queue := []*LevelQueue{first}
	for len(queue) > 0 {
		first = queue[0]
		queue = queue[1:]
		// fmt.Printf("queue = %v\n", queue)
		if first.Node.Left != nil {
			left := &LevelQueue{Level: first.Level + 1, Node: first.Node.Left}
			queue = append(queue, left)
		}
		if first.Node.Right != nil {
			right := &LevelQueue{Level: first.Level + 1, Node: first.Node.Right}
			queue = append(queue, right)
		}
		result[first.Level] = append(result[first.Level], first.Node)
		// fmt.Printf("level %v: first = %v\n", first.Level, first.Node)
		// fmt.Printf("queue = %v\n", queue)
	}
	return result
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
	// result: [[3],[9,20],[15,7]]
    rootNums := []int{3, 9, 20, -1001, -1001, 15, 7}

	// result:
    // rootNums := []int{}

    root := treenode.Maketree(rootNums)
	result := levelOrder(root)
	fmt.Printf("result = %v\n", result)
}

