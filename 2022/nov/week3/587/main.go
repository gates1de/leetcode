package main
import (
	"fmt"
	"sort"
)

func outerTrees(points [][]int) [][]int {
    if len(points) <= 1 {
        return points
    }

    bl := bottomLeft(points)
    sort.Slice(points, func (i int, j int) bool {
        diff := orientation(bl, points[i], points[j])
        if diff == 0 {
            return distance(bl, points[i]) < distance(bl, points[j])
        } else {
            return diff < 0
        }
    })

    i := len(points) - 1
    for i >= 0 && orientation(bl, points[len(points) - 1], points[i]) == 0 {
        i--
    }

    j, k := i + 1, len(points) - 1
    for j < k {
        points[j], points[k] = points[k], points[j]
        j++
        k--
    }

    stack := [][]int{points[0], points[1]}

    for p := 2; p < len(points); p++ {
        top := stack[len(stack) - 1]
        stack = stack[:len(stack) - 1]
        for orientation(stack[len(stack) - 1], top, points[p]) > 0 {
            top = stack[len(stack) - 1]
            stack = stack[:len(stack) - 1]
        }
        stack = append(stack, top)
        stack = append(stack, points[p])
    }
    return stack
}

func bottomLeft(points [][]int) []int {
    result := points[0]
    for _, point := range points {
        if point[1] < result[1] {
            result = point
        }
    }
    return result
}

func orientation(p []int, q []int, r []int) int {
    return (q[1] - p[1]) * (r[0] - q[0]) - (q[0] - p[0]) * (r[1] - q[1])
}

func distance(p []int, q []int) int {
    return (p[0] - q[0]) * (p[0] - q[0]) + (p[1] - q[1]) * (p[1] - q[1])
}

func main() {
	// result: [[1,1],[2,0],[3,3],[2,4],[4,2]]
	// points := [][]int{{1,1},{2,2},{2,0},{2,4},{3,3},{4,2}}

	// result: [[4,2],[2,2],[1,2]]
	points := [][]int{{1,2},{2,2},{4,2}}

	// result: 
	// points := [][]int{}

	result := outerTrees(points)
	fmt.Printf("result = %v\n", result)
}

