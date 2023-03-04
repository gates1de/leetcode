package main
import (
	"fmt"
)

type Node struct {
    Val bool
    IsLeaf bool
    TopLeft *Node
    TopRight *Node
    BottomLeft *Node
    BottomRight *Node
}

func construct(grid [][]int) *Node {
	return helper(len(grid), 0, 0, grid)
}

func helper(n int, x int, y int, grid [][]int) *Node {
	if n == 1 {
		return &Node{Val: grid[x][y] == 1, IsLeaf: true}
	}

	topLeft := helper(n / 2, x, y, grid)
	topRight := helper(n / 2, x, y + (n / 2), grid)
	bottomLeft := helper(n / 2, x + (n / 2), y, grid)
	bottomRight := helper(n / 2, x + (n / 2), y + (n / 2), grid)

	if topLeft.IsLeaf &&
		topRight.IsLeaf &&
		bottomLeft.IsLeaf &&
		bottomRight.IsLeaf &&
		topLeft.Val == topRight.Val &&
		topRight.Val == bottomLeft.Val &&
		bottomLeft.Val == bottomRight.Val {
		  return &Node{Val: topLeft.Val, IsLeaf: true}
	}

	return &Node{
		Val: false,
		IsLeaf: false,
		TopLeft: topLeft,
		TopRight: topRight,
		BottomLeft: bottomLeft,
		BottomRight: bottomRight,
	}
}

func main() {
	// result: [[0,1],[1,0],[1,1],[1,1],[1,0]]
	// grid := [][]int{{0,1},{1,0}}

	// result: [[0,1],[1,1],[0,1],[1,1],[1,0],null,null,null,null,[1,0],[1,0],[1,1],[1,1]]
	grid := [][]int{
		{1,1,1,1,0,0,0,0},
		{1,1,1,1,0,0,0,0},
		{1,1,1,1,1,1,1,1},
		{1,1,1,1,1,1,1,1},
		{1,1,1,1,0,0,0,0},
		{1,1,1,1,0,0,0,0},
		{1,1,1,1,0,0,0,0},
		{1,1,1,1,0,0,0,0},
	}

	// result: 
	// grid := [][]int{}

	result := construct(grid)
	fmt.Printf("result = %v\n", result)
}

