package main
import (
	"fmt"
)

func change(amount int, coins []int) int {
	n := len(coins)
	dp := make([]int, amount + 1)
	dp[0] = 1

	for i := n - 1; i >= 0; i-- {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j - coins[i]]
		}
	}

	return dp[amount]
}

func main() {
	// result: 4
	// amount := int(5)
	// coins := []int{1,2,5}

	// result: 0
	// amount := int(3)
	// coins := []int{2}

	// result: 1
	amount := int(10)
	coins := []int{10}

	// result: 
	// amount := int(0)
	// coins := []int{}

	result := change(amount, coins)
	fmt.Printf("result = %v\n", result)
}

