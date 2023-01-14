package main
import (
	"fmt"
)

type Node struct {
	Index int
	HasApple bool
	Children []*Node
}

func minTime(n int, edges [][]int, hasApple []bool) int {
	adjacents := make([][]int, n)
	for _, edge := range edges {
		n1, n2 := edge[0], edge[1]
		adjacents[n1] = append(adjacents[n1], n2)
		adjacents[n2] = append(adjacents[n2], n1)
	}

	visited := make([]bool, n)
	return dfs(0, visited, adjacents, hasApple) * 2
}

func dfs(index int, visited []bool, adjacents [][]int, hasApple []bool) (count int) {
	visited[index] = true
	for _, n := range adjacents[index] {
		if !visited[n] {
			count += dfs(n, visited, adjacents, hasApple)
		}
	}

	if (hasApple[index] || count > 0) && index != 0 {
		count += 1
	}
	return count
}

// Wrong Answer
func ngSolution(n int, edges [][]int, hasApple []bool) int {
	nodes := make([]*Node, n)
	for i := 0; i < n; i++ {
		nodes[i] = &Node{Index: i, HasApple: hasApple[i]}
	}
	for _, edge := range edges {
		node := nodes[edge[0]]
		child := nodes[edge[1]]
		if node.Children == nil {
			node.Children = make([]*Node, 0, n - 1)
		}
		node.Children = append(node.Children, child)
	}

	return preorder(nodes[0])
}

func preorder(node *Node) int {
	if node == nil {
		return 0
	}

	time := int(0)
	for _, child := range node.Children {
		time += preorder(child)
	}

	if node.Index != 0 && (node.HasApple || time > 0) {
		time += 2
	}
	return time
}

func printNode(node *Node) {
	if node == nil {
		return
	}

	fmt.Println(node, node.Index)
	for _, child := range node.Children {
		printNode(child)
	}
}

func main() {
	// result: 8
	// n := int(7)
	// edges := [][]int{{0,1},{0,2},{1,4},{1,5},{2,3},{2,6}}
	// hasApple := []bool{false,false,true,false,true,true,false}

	// result: 6
	// n := int(7)
	// edges := [][]int{{0,1},{0,2},{1,4},{1,5},{2,3},{2,6}}
	// hasApple := []bool{false,false,true,false,false,true,false}

	// result: 0
	// n := int(7)
	// edges := [][]int{{0,1},{0,2},{1,4},{1,5},{2,3},{2,6}}
	// hasApple := []bool{false,false,false,false,false,false,false}

	// result: 4
	n := int(4)
	edges := [][]int{{0,2},{0,3},{1,2}}
	hasApple := []bool{false,true,false,false}

	// result: 
	// n := int(0)
	// edges := [][]int{}
	// hasApple := []bool{}

	result := minTime(n, edges, hasApple)
	fmt.Printf("result = %v\n", result)
}

