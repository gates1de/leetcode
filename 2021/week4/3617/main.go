package main
import (
	"fmt"
	"sort"
)

type Node struct {
	X int
	Y int
	Cost int
	IsDone bool
	Ref *Node
}

func minimumEffortPath(heights [][]int) int {
	if len(heights) <= 0 || len(heights[0]) <= 0 {
		return 0
	}
	maxX := len(heights[0]) - 1
	maxY := len(heights) - 1

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	start := &Node{X: 0, Y: 0, Cost: 0, IsDone: false}
	var goal *Node
	openNodes := []*Node{start}
	closeNodes := []*Node{}
	for len(openNodes) > 0 {
		var current *Node
		current, openNodes = pop(openNodes)
		// fmt.Printf("current = %v\n", current)
		// fmt.Printf("after openNodes = %v\n", openNodes)
		if current.X == maxX && current.Y == maxY {
			refValue := heights[current.Ref.Y][current.Ref.X]
			goalValue := heights[current.Y][current.X]
			cost := max(current.Cost, abs(refValue, goalValue))
			// goal = &Node{X: maxX, Y: maxY, Cost: cost, IsDone: true, Ref: current.Ref}
			// break
			// current.IsDone = true

			if goal == nil {
				goal = &Node{X: maxX, Y: maxY, Cost: cost, IsDone: true, Ref: current.Ref}
			} else {
				// refValue := heights[current.Ref.Y][current.Ref.X]
				// goalValue := heights[current.Y][current.X]
				// cost := max(current.Cost, abs(refValue, goalValue))
				if cost < goal.Cost {
					goal = &Node{X: maxX, Y: maxY, Cost: cost, IsDone: true, Ref: current.Ref}
				}
			}
			continue
		}

		if contains(current, closeNodes) {
			continue
		}
		closeNodes = append(closeNodes, current)
		// if current.IsDone {
		// 	continue
		// }
		// current.IsDone = true
		for i := 0; i < 4; i++ {
			x := current.X + dx[i]
			y := current.Y + dy[i]
			// fmt.Printf("x = %v, y = %v\n", x, y)
			if (x < 0 || maxX < x) || (y < 0 || maxY < y) {
				continue
			}
			currentValue := heights[current.Y][current.X]
			adjacentValue := heights[y][x]
			// fmt.Printf("x = %v, y = %v, currentValue = %v, adjacentValue = %v\n", x, y, currentValue, adjacentValue)
			cost := max(current.Cost, abs(currentValue, adjacentValue))
			adjacentNode := &Node{X: x, Y: y, Cost: cost, IsDone: false, Ref: current}
			if contains(adjacentNode, openNodes) {
				continue
			}
			openNodes = append(openNodes, adjacentNode)
		}
		sort.Slice(openNodes, func (i, j int) bool { return openNodes[i].Cost < openNodes[j].Cost })
	}
	printNode(goal)
	return goal.Cost
}

func contains(target *Node, nodes []*Node) bool {
	for _, node := range nodes {
		if node.X == target.X && node.Y == target.Y {
			return true
		}
	}
	return false
}

func abs(a int, b int) int {
	if a - b < 0 {
		return b - a
	}
	return a - b
}

func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

func pop(nodes []*Node) (*Node, []*Node) {
    if len(nodes) == 0 {
        return nil, nodes
    }
    result := nodes[0]
    nodes = nodes[1:]
    return result, nodes
}

func printNode(n *Node) {
	if n == nil {
	    return
	}
	fmt.Printf("n = %v\n", n)
	printNode(n.Ref)
}

func main() {
	// heights := [][]int{{1, 2, 2}, {3, 8, 2}, {5, 3, 5}}
	heights := [][]int{{1, 2, 3}, {3, 8, 4}, {5, 3, 5}}
	result := minimumEffortPath(heights)
	fmt.Printf("result = %v\n", result)
}

