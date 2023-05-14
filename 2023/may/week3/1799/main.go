package main
import (
	"fmt"
	"strconv"
	"strings"
)

func maxScore(nums []int) int {
	n := len(nums)
	maxStates := 1 << n
	finalMask := maxStates - 1

	dp := make([]int, maxStates)

	for state := finalMask; state >= 0; state-- {
		if state == finalMask {
			dp[state] = 0
			continue
		}

		stateBinary := strconv.FormatInt(int64(state), 2)
		numbersTaken := strings.Count(fmt.Sprintf("%v", stateBinary), "1")
		pairsFormed := numbersTaken / 2

		if numbersTaken % 2 != 0 {
			continue
		}

		for firstIndex := 0; firstIndex < n; firstIndex++ {
			for secondIndex := firstIndex + 1; secondIndex < n; secondIndex++ {
				if ((state >> firstIndex) & 1) == 1 || ((state >> secondIndex) & 1) == 1 {
					continue
				}
				currentScore := (pairsFormed + 1) * gcd(nums[firstIndex], nums[secondIndex])
				stateAfterPickingCurrentPair := state | (1 << firstIndex) | (1 << secondIndex)
				remainingScore := dp[stateAfterPickingCurrentPair]
				dp[state] = max(dp[state], currentScore + remainingScore)
			}
		}
	}

	return dp[0]
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a % b)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 1
	// nums := []int{1,2}

	// result: 11
	// nums := []int{3,4,6,8}

	// result: 14
	nums := []int{1,2,3,4,5,6}

	// result: 
	// nums := []int{}

	result := maxScore(nums)
	fmt.Printf("result = %v\n", result)
}

