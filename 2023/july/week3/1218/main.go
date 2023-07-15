package main
import (
	"fmt"
)

func longestSubsequence(arr []int, difference int) int {
	result := int(1)
	dp := make(map[int]int)
	for _, num := range arr {
		before := dp[num - difference]
		dp[num] = before + 1
		result = max(result, dp[num])
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
	// result: 4
	// arr := []int{1,2,3,4}
	// difference := int(1)

	// result: 1
	// arr := []int{1,3,5,7}
	// difference := int(1)

	// result: 4
	arr := []int{1,5,7,8,5,3,4,2,1}
	difference := int(-2)

	// result: 
	// arr := []int{}
	// difference := int(0)

	result := longestSubsequence(arr, difference)
	fmt.Printf("result = %v\n", result)
}

