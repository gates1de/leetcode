package main
import (
	"fmt"
)

func findJudge(n int, trust [][]int) int {
	in := make([]int, n + 1)
	out := make([]int, n + 1)

	for _, t := range trust {
		out[t[0]]++
		in[t[1]]++
	}

	for i := 1; i <= n; i++ {
		if in[i] == n - 1 && out[i] == 0 {
			return i
		}
	}

	return -1
}

func main() {
	// result: 2
	// n := int(2)
	// trust := [][]int{{1,2}}

	// result: 3
	// n := int(3)
	// trust := [][]int{{1,3},{2,3}}

	// result: -1
	n := int(3)
	trust := [][]int{{1,3},{2,3},{3,1}}

	// result: 
	// n := int(0)
	// trust := [][]int{}

	result := findJudge(n, trust)
	fmt.Printf("result = %v\n", result)
}

