package main
import (
	"fmt"
	"sort"
)

const modulo = int(1e9 + 7)

func numFactoredBinaryTrees(arr []int) int {
    sort.Ints(arr)
    n := len(arr)
    dp := make([]int, n)
    for i, _ := range arr {
        dp[i] = 1
    }

    indexes := map[int]int{}
    for i, num := range arr {
        indexes[num] = i
    }
    for i, num1 := range arr {
        for j, num2 := range arr {
            if num1 % num2 != 0 {
                continue
            }

            right := num1 / num2
            if _, ok := indexes[right]; ok {
                dp[i] = (dp[i] + dp[j] * dp[indexes[right]]) % modulo
            }
        }
    }

    result := int(0)
    for _, val := range dp {
        result += val
    }
    return result % modulo
}

func main() {
	// result: 3
	// arr := []int{2,4}

	// result: 7
	arr := []int{2,4,5,10}

	// result: 
	// arr := []int{}

	result := numFactoredBinaryTrees(arr)
	fmt.Printf("result = %v\n", result)
}

