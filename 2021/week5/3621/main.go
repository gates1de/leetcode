package main
import (
	"./treenode"
	"fmt"
	"reflect"
	"sort"
)

func verticalTraversal(root *treenode.TreeNode) [][]int {
	result := [][]int{}
	dict := map[int]map[int][]int{}
	traversal(root, 0, 0, dict)
	// fmt.Printf("dict = %v\n", dict)

	keys := sortedKeys(dict)
	for _, col := range keys {
		val := dict[col]
		valKeys := sortedKeys(val)
		rowVals := []int{}
		for _, row := range valKeys {
			rowVal := dict[col][row]
			// fmt.Printf("rowVal = %v\n", rowVal)
			rowVals = append(rowVals, rowVal...)
		}
		result = append(result, rowVals)
	}
	return result
}

func traversal(node *treenode.TreeNode, row int, col int, dict map[int]map[int][]int) *treenode.TreeNode {
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
	// fmt.Printf("(%v, %v): node.Val = %v\n", row, col, node.Val)
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

func main() {
	// rootNums := []int{3, 9, 20, -1, -1, 15, 7}
	// rootNums := []int{1, 2, 3, 4, 5, 6, 7}
	// rootNums := []int{1, 2, 3, 4, 6, 5, 7}
	rootNums := []int{3, 1, 4, 0, 2, 2}

	root := treenode.Maketree(rootNums)
	result := verticalTraversal(root)
	fmt.Printf("result = %v\n", result)
}

