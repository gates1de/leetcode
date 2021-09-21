package main
import (
	"fmt"
)

type Node struct {
    Val int
    Children []*Node
}

func levelOrder(root *Node) [][]int {
    res := [][]int{}
    recursive(root, 0, &res)
    return res
}

func recursive(root *Node, level int, res *[][]int) {
    if root == nil {
        return
    }

    for _, v := range root.Children {
        recursive(v, level+1, res)
    }

    for len(*res) <= level {
        *res = append(*res, []int{})
    }

    (*res)[level] = append((*res)[level], root.Val)
}

// Slow & Use more memory
func mySolution(root *Node) [][]int {
	levelVals := map[int][]int{}
	bfs(root, levelVals)
	result := [][]int{}
	for i := 0; i < 1000; i++ {
		if _, ok := levelVals[i]; !ok {
			break
		}
		result = append(result, levelVals[i])
	}
	return result
}

func bfs(root *Node, levelVals map[int][]int) {
	if root == nil {
		return
	}

	queue := []*Node{root}
	levels := map[*Node]int{}
	levelVals[0] = []int{root.Val}
	for len(queue) > 0 {
		node := queue[0]
		currentLevel := levels[node]
		if node.Children != nil && len(node.Children) > 0 {
			for _, n := range node.Children {
				levels[n] = currentLevel + 1
				levelVals[levels[n]] = append(levelVals[levels[n]], n.Val)
				queue = append(queue, n)
			}
		}
		queue = queue[1:]
		// fmt.Printf("currentLevel = %v, node.Val = %v, queue = %v\n", currentLevel, node.Val, queue)
	}
}

func makeNode1() *Node {
	root := &Node{Val: 1}
	child1 := &Node{Val: 3}
	child2 := &Node{Val: 2}
	child3 := &Node{Val: 4}
	child4 := &Node{Val: 5}
	child5 := &Node{Val: 6}
	root.Children = []*Node{child1, child2, child3}
	child1.Children = []*Node{child4, child5}
	return root
}

func main() {
	// result: 
	root := makeNode1()

	// result: 
	// root := makeNode()

	// result: 
	// root := makeNode()

	result := levelOrder(root)
	fmt.Printf("result = %v\n", result)
}

