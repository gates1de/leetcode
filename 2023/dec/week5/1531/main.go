package main
import (
	"fmt"
)

func getLengthOfOptimalCompression(s string, k int) int {
    n := len(s)
    dp := make([][]int, 101)
    for i, _ := range dp {
        dp[i] = make([]int, 101)
        for j, _ := range dp[i] {
            dp[i][j] = 10000
        }
    }

    dp[0][0] = 0
    for i := 1; i <= n; i++ {
        for j := 0; j <= k; j++ {
            count := 0
            del := 0
            for l := i; l >= 1; l-- {
                if s[l - 1] == s[i - 1] {
                    count++
                } else {
                    del++
                }

                if j - del >= 0 {
                    numsLength := int(0)
                    if count >= 100 {
                        numsLength = 3
                    } else if count >= 10 {
                        numsLength = 2
                    } else if count >= 2 {
                        numsLength = 1
                    }
                    dp[i][j] = min(dp[i][j], dp[l - 1][j - del] + 1 + numsLength)
                }
            }

            if j > 0 {
                dp[i][j] = min(dp[i][j], dp[i - 1][j - 1])
            }
        }
    }

    return dp[n][k]
}

func min(a, b int) int {
    if b < a {
        return b
    }
    return a
}

func main() {
	// result: 4
	// s := "aaabcccd"
	// k := int(2)

	// result: 2
	// s := "aabbaa"
	// k := int(2)

	// result: 3
	s := "aaaaaaaaaaa"
	k := int(0)

	// result: 
	// s := ""
	// k := int(0)

	result := getLengthOfOptimalCompression(s, k)
	fmt.Printf("result = %v\n", result)
}

