package main
import (
	"fmt"
	"sort"
)

func kClosest(points [][]int, k int) [][]int {
	if len(points) == k {
		return points
	}

	distances := make([][3]int, len(points))
	for i, point := range points {
		dist := distance(point[0], point[1])
		distances[i] = [3]int{dist, point[0], point[1]}
	}
	sort.Slice(distances, func(a, b int) bool {
		return distances[a][0] < distances[b][0]
	})

	result := make([][]int, k)
	for i := 0; i < k; i++ {
		distance := distances[i]
		result[i] = []int{distance[1], distance[2]}
	}
    return result
}

func distance(x int, y int) int {
	return (x * x) + (y * y)
}

func main() {
	// result: [[-2,2]]
	// points := [][]int{{1,3},{-2,2}}
	// k := int(1)

	// result: [[3,3],[-2,4]]
	// points := [][]int{{3,3},{5,-1},{-2,4}}
	// k := int(2)

	// result: [[1,1]]
	// points := [][]int{{1,1}}
	// k := int(1)

	// result: [[1,1][-3,2],[4,4],[3,5]]
	points := [][]int{{1,1},{3,5},{-7,3},{4,4},{9,-1},{-3,2}}
	k := int(4)

	// result: 
	// points := [][]int{}
	// k := int(0)

	result := kClosest(points, k)
	fmt.Printf("result = %v\n", result)
}

