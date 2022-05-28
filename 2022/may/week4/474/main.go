package main
import (
	"fmt"
	"strings"
)

func findMaxForm(strs []string, m int, n int) int {
    memo := make([][]int, m + 1)
    for i := range memo {
        memo[i] = make([]int, n + 1)
    }

    for _, s := range strs {
        zeroCount := strings.Count(s, "0")
        oneCount := len(s) - zeroCount

        for i := m; i >= zeroCount; i-- {
            for j := n; j >= oneCount; j-- {
                t := memo[i - zeroCount][j - oneCount] + 1
                if t > memo[i][j] {
                    memo[i][j] = t
                }
            }
        }
    }

    return memo[m][n]
}

func main() {
	// result: 4
	// strs := []string{"10","0001","111001","1","0"}
	// m := int(5)
	// n := int(3)

	// result: 2
	strs := []string{"10","0","1"}
	m := int(1)
	n := int(1)

	// result: 
	// strs := []string{}
	// m := int(0)
	// n := int(0)

	result := findMaxForm(strs, m, n)
	fmt.Printf("result = %v\n", result)
}

