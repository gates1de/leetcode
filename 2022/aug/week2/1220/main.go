package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

func countVowelPermutation(n int) int {
	if n <= 0 {
		return 0
	}

	dp := make([][]int, n + 1)
	for i, _ := range dp {
		dp[i] = make([]int, 5)
	}
	for i, _ := range dp[0] {
		dp[0][i] = 1
	}

	result := int(5)
	for i := 1; i < n; i++ {
		result = count(i, dp) % modulo
	}

	return result
}

func count(n int, dp [][]int) int {
	result := int(0)
	for i := 0; i < 5; i++ {
		numOfV := dp[n - 1][i] % modulo
        if i == 0 {
            dp[n][1] += numOfV
            result += numOfV
        } else if i == 1 {
            dp[n][0] += numOfV
            dp[n][2] += numOfV
            result += numOfV * 2
        } else if i == 2 {
            dp[n][0] += numOfV
            dp[n][1] += numOfV
            dp[n][3] += numOfV
            dp[n][4] += numOfV
            result += numOfV * 4
        } else if i == 3 {
            dp[n][2] += numOfV
            dp[n][4] += numOfV
            result += numOfV * 2
        } else if i == 4 {
            dp[n][0] += numOfV
            result += numOfV
        }
	}

	return result
}

func main() {
	// result: 5
	// n := int(1)

	// result: 10
	// n := int(2)

	// result: 19
	// n := int(3)

	// result: 35
	// n := int(4)

	// result: 68
	// n := int(5)

	// result: 76428576
	n := int(10000)

	// result: 
	// n := int(0)

	result := countVowelPermutation(n)
	fmt.Printf("result = %v\n", result)
}

