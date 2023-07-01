package main
import (
	"fmt"
	"math"
)

func distributeCookies(cookies []int, k int) int {
	distribute := make([]int, k)
	return dfs(0, distribute, cookies, k, k)
}

func dfs(i int, distribute []int, cookies []int, k int, zeroCount int) int {
	n := len(cookies)
	if n - i < zeroCount {
		return math.MaxInt32
	}

	if i == n {
		unfairness := math.MinInt32
		for _, val := range distribute {
			unfairness = max(unfairness, val)
		}
		return unfairness
	}

	result := math.MaxInt32
	for j := 0; j < k; j++ {
		if distribute[j] == 0 {
			zeroCount -= 1
		}
		distribute[j] += cookies[i]

		result = min(result, dfs(i + 1, distribute, cookies, k, zeroCount))

		distribute[j] -= cookies[i]
		if distribute[j] == 0 {
			zeroCount += 1
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

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 31
	// cookies := []int{8,15,10,20,8}
	// k := int(2)

	// result: 7
	cookies := []int{6,1,3,2,2,4,1,2}
	k := int(3)

	// result: 
	// cookies := []int{}
	// k := int(0)

	result := distributeCookies(cookies, k)
	fmt.Printf("result = %v\n", result)
}

