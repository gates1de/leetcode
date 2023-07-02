package main
import (
	"fmt"
)

func maximumRequests(n int, requests [][]int) int {
	result := int(0)
	indegree := make([]int, n)
	maxRequest(n, 0, 0, requests, indegree, &result)
	return result
}

func maxRequest(n int, index int, count int, requests [][]int, indegree []int, result *int) {
	if index == len(requests) {
		for _, val := range indegree {
			if val != 0 {
				return
			}
		}

		*result = max(*result, count)
		return
	}

	indegree[requests[index][0]]--
	indegree[requests[index][1]]++
	maxRequest(n, index + 1, count + 1, requests, indegree, result)
	indegree[requests[index][0]]++
	indegree[requests[index][1]]--
	maxRequest(n, index + 1, count, requests, indegree, result)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 5
	// n := int(5)
	// requests := [][]int{{0,1},{1,0},{0,1},{1,2},{2,0},{3,4}}

	// result: 3
	// n := int(3)
	// requests := [][]int{{0,0},{1,2},{2,1}}

	// result: 5
	n := int(4)
	requests := [][]int{{0,3},{3,1},{1,2},{2,0}}

	// result: 
	// n := int(0)
	// requests := [][]int{}

	result := maximumRequests(n, requests)
	fmt.Printf("result = %v\n", result)
}

