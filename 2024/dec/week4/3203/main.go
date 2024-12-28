package main
import (
	"fmt"
	"math"
)

type Pair struct {
	First int
	Second int
}

func minimumDiameterAfterMerge(edges1 [][]int, edges2 [][]int) int {
	n := len(edges1) + 1
	m := len(edges2) + 1

	adjacents1 := getAdjacents(n, edges1)
	adjacents2 := getAdjacents(m, edges2)

	diameter1 := findDiameter(adjacents1, 0, -1).First
	diameter2 := findDiameter(adjacents2, 0, -1).First
	combinedDiameter := int(math.Ceil(float64(diameter1) / 2)) + int(math.Ceil(float64(diameter2) / 2)) + 1

	return max(max(diameter1, diameter2), combinedDiameter)
}

func getAdjacents(size int, edges [][]int) [][]int {
	result := make([][]int, 0)

	for i := 0; i < size; i++ {
		result = append(result, make([]int, 0))
	}

	for _, edge := range edges {
		result[edge[0]] = append(result[edge[0]], edge[1])
		result[edge[1]] = append(result[edge[1]], edge[0])
	}

	return result
}

func findDiameter(adjacents [][]int, node int, parent int) Pair {
	maxDepth1 := int(0)
	maxDepth2 := int(0)
	diameter := int(0)

	for _, neighbor := range adjacents[node] {
		if neighbor == parent {
			continue
		}

		result := findDiameter(adjacents, neighbor, node)
		child := result.First
		depth := result.Second + 1

		diameter = max(diameter, child)

		if depth > maxDepth1 {
			maxDepth2 = maxDepth1
			maxDepth1 = depth
		} else if depth > maxDepth2 {
			maxDepth2 = depth
		}
	}

	diameter = max(diameter, maxDepth1 + maxDepth2)
	return Pair{First: diameter, Second: maxDepth1}
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// edges1 := [][]int{{0,1},{0,2},{0,3}}
	// edges2 := [][]int{{0,1}}

	// result: 5
	edges1 := [][]int{{0,1},{0,2},{0,3},{2,4},{2,5},{3,6},{2,7}}
	edges2 := [][]int{{0,1},{0,2},{0,3},{2,4},{2,5},{3,6},{2,7}}

	// result: 
	// edges1 := [][]int{}
	// edges2 := [][]int{}

	result := minimumDiameterAfterMerge(edges1, edges2)
	fmt.Printf("result = %v\n", result)
}

