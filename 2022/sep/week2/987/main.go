package main
import (
	"fmt"
	"reflect"
	"sort"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func makeTree1() *TreeNode {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	return root
}

func makeTree2() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}
	return root
}

func makeTree3() *TreeNode {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 6}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 7}
	return root
}

func makeTree4() *TreeNode {
	root := &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 13}
	root.Right.Left = &TreeNode{Val: 14}
	root.Right.Right = &TreeNode{Val: 5}
	root.Right.Left.Left = &TreeNode{Val: 9}
	root.Right.Left.Right = &TreeNode{Val: 10}
	root.Right.Right.Left = &TreeNode{Val: 8}
	root.Right.Right.Right = &TreeNode{Val: 15}
	return root
}

func makeTree() *TreeNode {
	var root *TreeNode
	return root
}

func verticalTraversal(root *TreeNode) [][]int {
    result := [][]int{}
    dict := map[int]map[int][]int{}
    traversal(root, 0, 0, dict)

    keys := sortedKeys(dict)
    for _, col := range keys {
        val := dict[col]
        valKeys := sortedKeys(val)
        rowVals := []int{}
        for _, row := range valKeys {
            rowVal := dict[col][row]
            rowVals = append(rowVals, rowVal...)
        }
        result = append(result, rowVals)
    }
    return result
}

func traversal(node *TreeNode, row int, col int, dict map[int]map[int][]int) *TreeNode {
    if node == nil {
        return nil
    }
    traversal(node.Left, row + 1, col - 1, dict)
    if len(dict[col]) == 0 {
        dict[col] = map[int][]int{row: {node.Val}}
    } else if len(dict[col]) > 0 && len(dict[col][row]) == 0 {
        dict[col][row] = []int{node.Val}
    } else if len(dict[col]) > 0 && len(dict[col][row]) > 0 {
        dict[col][row] = append(dict[col][row], node.Val)
        sort.Slice(dict[col][row], func (i, j int) bool { return dict[col][row][i] < dict[col][row][j] })
    }
    traversal(node.Right, row + 1, col + 1, dict)
    return nil
}

func sortedKeys(m interface{}) []int {
    v := reflect.ValueOf(m)
    if v.Kind() != reflect.Map {
        return []int{}
    }

    keys := v.MapKeys()

    result := make([]int, len(keys))
    index := int(0)
    for _, k := range keys {
        key := k.Convert(v.Type().Key())
        result[index] = key.Interface().(int)
        index++
    }
    sort.Slice(result, func (i, j int) bool { return result[i] < result[j] })
    return result
}

func ngSolution(root *TreeNode) [][]int {
	resultMap := map[int][]int{}
	coordMap := map[*TreeNode][]int{}
	coordMap[root] = []int{0, 0}
	stack := []*TreeNode{root}

	for len(stack) > 0 {
		node := stack[len(stack) - 1]
		stack = stack[:len(stack) - 1]

		coord := coordMap[node]
		y := coord[0]
		x := coord[1]

		if resultMap[x] == nil {
			resultMap[x] = []int{}
		}
		resultMap[x] = append(resultMap[x], node.Val)

		if node.Left != nil {
			stack = append(stack, node.Left)
			coordMap[node.Left] = []int{y + 1, x - 1}
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
			coordMap[node.Right] = []int{y + 1, x + 1}
		}
	}

	result := [][]int{}
	for i := -500; i <= 500; i++ {
		if _, ok := resultMap[i]; !ok {
			continue
		}
		result = append(result, resultMap[i])
	}
	return result
}

func main() {
	// result: [[9],[3,15],[20],[7]]
	// root := makeTree1()

	// result: [[4],[2],[1,5,6],[3],[7]]
	root := makeTree2()

	// result: [[4],[2],[1,5,6],[3],[7]]
	// root := makeTree3()

	// result: [[9],[2,14],[3,10,18],[5],[15]]
	// root := makeTree4()

	// result: 
	// root := makeTree()

	result := verticalTraversal(root)
	fmt.Printf("result = %v\n", result)
}

