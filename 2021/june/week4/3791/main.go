package main
import (
	"fmt"
)

// REF: Union Find Algorithm ( https://www.slideshare.net/chokudai/union-find-49066733 )

func FindRoot(p int, arr []int) int {
    for p != arr[p] {
        p = arr[p]
    }
    return p
}

func Union(p, q int, arr []int) {
    arr[q] = p
}

func findRedundantConnection(edges [][]int) []int {
    arr := make([]int, len(edges) + 1)
    for index := range arr {
        arr[index] = index
    }
    for _, edge := range edges {
        pRoot := FindRoot(edge[0], arr)
        qRoot := FindRoot(edge[1], arr)
        if pRoot == qRoot {
            return edge
        } else {
            Union(pRoot, qRoot, arr)
        }
    }
    return nil
}

// Wrong Answer
func ngSolution(edges [][]int) []int {
	result := make([]int, 2)
	visited := map[int]bool{}
	for _, edge := range edges {
		if visited[edge[0]] && visited[edge[1]] {
			result = edge
		}
		visited[edge[0]] = true
		visited[edge[1]] = true
	}
	return result
}

func main() {
	// result: [2,3]
	// edges := [][]int{{1, 2}, {1, 3}, {2, 3}}

	// result: [1,4]
	// edges := [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}

	// result: [4,10]
	edges := [][]int{{9,10},{5,8},{2,6},{1,5},{3,8},{4,9},{8,10},{4,10},{6,8},{7,9}}

	// result: 
	// edges := [][]int{}

	result := findRedundantConnection(edges)
	fmt.Printf("result = %v\n", result)
}

