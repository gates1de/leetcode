package main
import (
	"fmt"
)

func longestPath(parent []int, s string) int {
	n := len(parent)
	childrenCounts := make([]int, n)
	for node := 1; node < n; node++ {
		childrenCounts[parent[node]]++
	}

	queue := make([]int, 0, n)
	result := int(1)
	longestChains := make([][2]int, n)
	for node := 1; node < n; node++ {
		if childrenCounts[node] == 0 {
			longestChains[node][0] = 1
			queue = append(queue, node)
		}
	}

	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		p := parent[currentNode]
		longestChainStartingFromCurrentNode := longestChains[currentNode][0]
		if s[currentNode] != s[p] {
			if longestChainStartingFromCurrentNode > longestChains[p][0] {
				longestChains[p][1] = longestChains[p][0]
				longestChains[p][0] = longestChainStartingFromCurrentNode
			} else if longestChainStartingFromCurrentNode > longestChains[p][1] {
				longestChains[p][1] = longestChainStartingFromCurrentNode
			}
		}

		result = max(result, longestChains[p][0] + longestChains[p][1] + 1)
		childrenCounts[p]--

		if childrenCounts[p] == 0 && p != 0 {
			longestChains[p][0]++
			queue = append(queue, p)
		}
	}

	return result
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// parent := []int{-1,0,0,1,1,2}
	// s := "abacbe"

	// result: 3
	parent := []int{-1,0,0,0}
	s := "aabc"

	// result: 
	// parent := []int{}
	// s := ""

	result :=longestPath(parent, s)
	fmt.Printf("result = %v\n", result)
}

