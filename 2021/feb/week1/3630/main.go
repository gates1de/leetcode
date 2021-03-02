package main
import (
	"./treenode"
	"fmt"
	"reflect"
	"sort"
)

func rightSideView(root *treenode.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	// printNode(root)

	result := []int{}
	rightNums := &map[int]int{1: root.Val}
	traversal(root, rightNums, 1)
	// fmt.Printf("rightNums = %v\n", rightNums)
	sortedKeys := sortedKeys(*rightNums)
	for _, key := range sortedKeys {
		result = append(result, (*rightNums)[key])
	}
	return result
}

func traversal(root *treenode.TreeNode, rightNums *map[int]int, depth int) {
	if root == nil {
		return
	}

	// fmt.Printf("root = %v\n", root)

	if _, ok := (*rightNums)[depth]; !ok {
		(*rightNums)[depth] = root.Val
	}
	traversal(root.Right, rightNums, depth + 1)
	traversal(root.Left, rightNums, depth + 1)
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

func printNode(root *treenode.TreeNode) {
    if root == nil {
        return
    }
    fmt.Printf("node = %v\n", root)
    printNode(root.Left)
    printNode(root.Right)
}

func main() {
	// rootNums := []int{1, 2, 3, -1, 5, -1, 4}
	// rootNums := []int{1, -1, 3}
	// rootNums := []int{}
	// rootNums := []int{1, 2}
	rootNums := []int{1, 2, 3, 4}

	root := treenode.Maketree(rootNums)
	result := rightSideView(root)
	fmt.Printf("result = %v\n", result)
}

