package main
import (
	"fmt"
	"strconv"
)

const modulo = int(1e9 + 7)

func numberOfArrays(s string, k int) int {
	m := len(s)
	dp := make([]int, m + 1)
	return dfs(0, s, k, dp)
}

func dfs(start int, s string, k int, dp []int) int {
	if dp[start] != 0 {
		return dp[start]
	}

	if start == len(s) {
		return 1
	}

	if s[start] == '0' {
		return 0
	}

	count := int(0)
	for end := start; end < len(s); end++ {
		currentNumStr := s[start:end + 1]
		currentNum, _ := strconv.Atoi(currentNumStr)
		if currentNum > k {
			break
		}

		count = (count + dfs(end + 1, s, k, dp)) % modulo
	}

	dp[start] = count
	return count
}

func main() {
	// result: 1
	// s := "1000"
	// k := int(10000)

	// result: 0
	// s := "1000"
	// k := int(10)

	// result: 8
	s := "1317"
	k := int(2000)

	// result: 
	// s := ""
	// k := int(0)

	result := numberOfArrays(s, k)
	fmt.Printf("result = %v\n", result)
}

