package main
import (
	"fmt"
	"math"
)

func minimumFuelCost(roads [][]int, seats int) int64 {
	adjacents := make(map[int][]int)
	for _, road := range roads {
		if adjacents[road[0]] == nil {
			adjacents[road[0]] = make([]int, 0)
		}
		if adjacents[road[1]] == nil {
			adjacents[road[1]] = make([]int, 0)
		}
		adjacents[road[0]] = append(adjacents[road[0]], road[1])
		adjacents[road[1]] = append(adjacents[road[1]], road[0])
	}

	result := int64(0)
	dfs(0, -1, seats, adjacents, &result)
	return result
}

func dfs(node int, parent int, seats int, adjacents map[int][]int, result *int64) int64 {
	representatives := int64(1)
	if adjacents[node] == nil {
		return representatives
	}

	for _, child := range adjacents[node] {
		if child != parent {
			representatives += dfs(child, node, seats, adjacents, result)
		}
	}

	if node != 0 {
		*result += int64(math.Ceil(float64(representatives) / float64(seats)))
	}

	return representatives
}

func main() {
	// result: 3
	// roads := [][]int{{0,1},{0,2},{0,3}}
	// seats := int(5)

	// result: 7
	// roads := [][]int{{3,1},{3,2},{1,0},{0,4},{0,5},{4,6}}
	// seats := int(2)

	// result: 0
	roads := [][]int{}
	seats := int(1)

	// result: 
	// roads := [][]int{}
	// seats := int(0)

	result := minimumFuelCost(roads, seats)
	fmt.Printf("result = %v\n", result)
}

