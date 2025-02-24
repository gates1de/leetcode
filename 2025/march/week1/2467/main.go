package main
import (
	"fmt"
	"math"
)

func mostProfitablePath(edges [][]int, bob int, amount []int) int {
	n := len(amount)
	tree  := make([][]int, n)
	for i, _ := range tree {
		tree[i] = make([]int, 0)
	}
	distanceFromBob := make([]int, n)

	for _, edge := range edges {
		tree[edge[0]] = append(tree[edge[0]], edge[1])
		tree[edge[1]] = append(tree[edge[1]], edge[0])
	}

	return findPaths(0, 0, 0, bob, amount, tree, distanceFromBob)
}

func findPaths(
	sourceNode int,
	parentNode int,
	time int,
	bob int,
	amount []int,
	tree [][]int,
	distanceFromBob []int,
) int {
	n := len(amount)
	maxIncome := int(0)
	maxChild := math.MinInt32

	if sourceNode == bob {
		distanceFromBob[sourceNode] = 0
	} else {
		distanceFromBob[sourceNode] = n
	}

	for _, adjacent := range tree[sourceNode] {
		if adjacent != parentNode {
			maxChild = max(maxChild, findPaths(adjacent, sourceNode, time + 1, bob, amount, tree, distanceFromBob))
			distanceFromBob[sourceNode] = min(distanceFromBob[sourceNode], distanceFromBob[adjacent] + 1)
		}
	}

	if distanceFromBob[sourceNode] > time {
		maxIncome += amount[sourceNode]
	} else if distanceFromBob[sourceNode] == time {
		maxIncome += amount[sourceNode] / 2
	}

	if maxChild == math.MinInt32 {
		return maxIncome
	}

	return maxIncome + maxChild
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 6
	// edges := [][]int{{0,1},{1,2},{1,3},{3,4}}
	// bob := int(3)
	// amount := []int{-2,4,2,-4,6}

	// result: -7280
	edges := [][]int{{0,1}}
	bob := int(1)
	amount := []int{-7280,2350}

	// result: 
	// edges := [][]int{}
	// bob := int(0)
	// amount := []int{}

	result := mostProfitablePath(edges, bob, amount)
	fmt.Printf("result = %v\n", result)
}

