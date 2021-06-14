package main
import (
	"fmt"
)

func stoneGameVII(stones []int) int {
	memo := make([][]int, len(stones))
	for i, _ := range memo {
		memo[i] = make([]int, len(stones))
	}
	scoreSum := sum(stones)
	return pick(0, len(stones) - 1, scoreSum, memo, stones)
}

func pick(i int, j int, score int, memo [][]int, stones []int) int {
	if i > j {
		return 0
	}
	if memo[i][j] != 0 {
		return memo[i][j]
	}
	n := len(stones)
	currentTurn := n + i - j

	var finalScore int
    rightScore := pick(i + 1, j, score - stones[i], memo, stones)
    leftScore := pick(i, j - 1, score - stones[j], memo, stones)

	if currentTurn % 2 == 1 {
		// Alice's turn
        finalScore = max(rightScore - stones[i], leftScore - stones[j]) + score
	} else {
		// Bob's turn
        finalScore = min(rightScore + stones[i], leftScore + stones[j]) - score
	}
    memo[i][j] = finalScore
	// fmt.Printf("memo = %v\n", memo)
    return memo[i][j]
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func sum(nums []int) int {
	result := int(0)
	for _, num := range nums {
		result += num
	}
	return result
}

func main() {
	// result: 6
	// stones := []int{5, 3, 1, 4, 2}

	// result: 122
	stones := []int{7, 90, 5, 1, 100, 10, 10, 2}

	// result: 
	// stones := []int{}

	result := stoneGameVII(stones)
	fmt.Printf("result = %v\n", result)
}

