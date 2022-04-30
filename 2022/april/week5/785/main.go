package main
import (
	"fmt"
)

func isBipartite(graph [][]int) bool {
    colors := make([]int, len(graph))
    for i, _ := range graph {
        if colors[i] == 0 && !dfs(i, graph, colors, 1) {
            return false
        }
    }
    return true
}

func dfs(i int, graph [][]int, colors []int, color int) bool {
    if colors[i] != 0 {
        return colors[i] == color
    }

    colors[i] = color
    for _, to := range graph[i] {
        if !dfs(to, graph, colors, -color) {
            return false
        }
    }
    return true
}

// Wrong Answer
func ngSolution(graph [][]int) bool {
	if len(graph) <= 1 {
		return true
	}

	left := map[int]bool{}
	right := map[int]bool{}

	leftQueue := []int{0}
	rightQueue := []int{}
	isLeft := true
	unknowns := make([]int, len(graph))
	unknowns[0] = -1

	for len(leftQueue) != 0 || len(rightQueue) != 0 {
		if isLeft {
			if len(leftQueue) == 0 {
				isLeft = false
				continue
			}

			target := leftQueue[0]
			leftQueue = leftQueue[1:]
			left[target] = true
			for _, node := range graph[target] {
				if left[node] {
					return false
				} else if right[node] {
					continue
				}
				rightQueue = append(rightQueue, node)
				right[node] = true
				unknowns[node] = -1
			}

			isLeft = false
		} else {
			if len(rightQueue) == 0 {
				isLeft = true
				continue
			}

			target := rightQueue[0]
			rightQueue = rightQueue[1:]
			right[target] = true
			for _, node := range graph[target] {
				if right[node] {
					return false
				} else if left[node] {
					continue
				}
				leftQueue = append(leftQueue, node)
				left[node] = true
				unknowns[node] = -1
			}

			isLeft = true
		}
	}

	nextNode := int(-1)
	for i, unknown := range unknowns {
		if unknown == 0 {
			nextNode = i
			break
		}
	}

	if nextNode == -1 {
		return true
	}

	if helper(nextNode, true, left, right, unknowns, graph) {
		return true
	}

	return helper(nextNode, false, left, right, unknowns, graph)
}

func helper(node int, isLeft bool, left map[int]bool, right map[int]bool, unknowns []int, graph [][]int) bool {
	connected := graph[node]
	newLeft := cloneMap(left)
	newRight := cloneMap(right)
	newUnknowns := make([]int, len(unknowns))
	copy(newUnknowns, unknowns)
	newUnknowns[node] = -1

	for _, node := range connected {
		if isLeft {
			if right[node] {
				return false
			}
			newRight[node] = true
			newUnknowns[node] = -1
		} else {
			if left[node] {
				return false
			}
			newLeft[node] = true
			newUnknowns[node] = -1
		}
	}

	nextNode := int(-1)
	for i, unknown := range newUnknowns {
		if unknown == 0 {
			nextNode = i
			break
		}
	}

	if nextNode == -1 {
		return true
	}

	return helper(nextNode, !isLeft, newLeft, newRight, newUnknowns, graph)
}

func cloneMap(target map[int]bool) map[int]bool {
	result := map[int]bool{}
	for key, val := range target {
		result[key] = val
	}
	return result
}

func main() {
	// result: false
	// graph := [][]int{{1,2,3},{0,2},{0,1,3},{0,2}}

	// result: true
	// graph := [][]int{{1,3},{0,2},{1,3},{0,2}}

	// result: true
	// graph := [][]int{{1}, {0, 3}, {3}, {1, 2}}

	// result: false
	graph := [][]int{{1,2,3},{0,3,4},{0,3},{0,1,2},{1}}

	// result: 
	// graph := [][]int{}

	result := isBipartite(graph)
	fmt.Printf("result = %v\n", result)
}

