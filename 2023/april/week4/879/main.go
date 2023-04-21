package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

func profitableSchemes(n int, minProfit int, group []int, profit []int) int {
	memo := make([][][]int, 101)
	for i, _ := range memo {
		memo[i] = make([][]int, 101)
		for j, _ := range memo[i] {
			memo[i][j] = make([]int, 101)
			for k, _ := range memo[i][j] {
				memo[i][j][k] = -1
			}
		}
	}

	return find(0, 0, 0, n, minProfit, group, profit, memo)
}

func find(position int, count int, profit int, n int, minProfit int, group []int, profits []int, memo [][][]int) int {
	if position == len(group) {
		if profit >= minProfit {
			return 1
		}
		return 0
	}

	if memo[position][count][profit] != -1 {
		return memo[position][count][profit]
	}

	totalWays := find(position + 1, count, profit, n, minProfit, group, profits, memo)
	if count + group[position] <= n {
		totalWays += find(position + 1, count + group[position], min(minProfit, profit + profits[position]), n, minProfit, group, profits, memo)
	}

	memo[position][count][profit] = totalWays % modulo
	return memo[position][count][profit]
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 2
	// n := int(5)
	// minProfit := int(3)
	// group := []int{2,2}
	// profit := []int{2,3}

	// result: 7
	n := int(10)
	minProfit := int(5)
	group := []int{2,3,5}
	profit := []int{6,7,8}

	// result: 
	// n := int(0)
	// minProfit := int(0)
	// group := []int{}
	// profit := []int{}

	result := profitableSchemes(n, minProfit, group, profit)
	fmt.Printf("result = %v\n", result)
}

