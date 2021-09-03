package main
import (
	"fmt"
	"sort"
)

func outerTrees(points [][]int) [][]int {
    if len(points) <= 1 {
        return points
    }

    bm := bottomLeft(points)
    sort.Slice(points, func (i int, j int) bool {
        diff := orientation(bm, points[i], points[j])
        if diff == 0 {
            return distance(bm, points[i]) < distance(bm, points[j])
        } else {
            return diff < 0
        }
    })

    i := len(points) - 1
    for i >= 0 && orientation(bm, points[len(points) - 1], points[i]) == 0 {
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
    res := points[0]
    for _, v := range points {
        if v[1] < res[1] {
            res = v
        }
    }
    return res
}

func orientation(p []int, q []int, r []int) int {
    return (q[1] - p[1]) * (r[0] - q[0]) - (q[0] - p[0]) * (r[1] - q[1])
}

func distance(p []int, q []int) int {
    return (p[0] - q[0]) * (p[0] - q[0]) + (p[1] - q[1]) * (p[1] - q[1])
}

// Wrong Answer
func ngSolution(trees [][]int) [][]int {
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

