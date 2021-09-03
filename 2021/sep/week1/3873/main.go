package main
import (
	"fmt"
	"sort"
)

func outerTrees(trees [][]int) [][]int {
	sort.Slice(trees, func (a, b int) bool { return trees[a][1] < trees[b][1] })
	return convexHull(trees, len(trees))
}

func orientationMatch(p1 []int, p2 []int, p3 []int) int {
	val := (p2[1] - p1[1]) * (p3[0] - p2[0]) - (p2[0] - p1[0]) * (p3[1] - p2[1])
    if val == 0 {
		return 0
	} else if val > 0 {
		return 1
	}
	return 2
}

func convexHull(trees [][]int, length int) [][]int {
	if length < 3 {
		return trees
	}

	leftmost := int(0)
	for i := 1; i < length; i++ {
		if trees[i][0] < trees[leftmost][0] {
			leftmost = i
		}
	}

	result := [][]int{}
	p := leftmost
	q := int(0)
	for {
		result = append(result, trees[p])
		q = (p + 1) % length
		for i := 0; i < length; i++ {
			if orientationMatch(trees[p], trees[i], trees[q]) == 2 {
				q = i
			}
		}
		p = q

		if p == leftmost {
			break
		}
	}

	return result
}

func main() {
	// result: [[1,1],[2,0],[3,3],[2,4],[4,2]]
	// trees := [][]int{{1, 1}, {2, 2}, {2, 0}, {2, 4}, {3, 3}, {4, 2}}

	// result: [[4,2],[2,2],[1,2]]
	trees := [][]int{{1, 2}, {2, 2}, {4, 2}}

	// result: 
	// trees := [][]int{}

	// result: 
	// trees := [][]int{}

	result := outerTrees(trees)
	fmt.Printf("result = %v\n", result)
}

