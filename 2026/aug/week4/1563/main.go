package main

import (
	"fmt"
)

func stoneGameV(stoneValue []int) int {
	n := len(stoneValue)
	if n < 2 {
		return 0
	}

	prefix := make([]int, n+1)
	for i, value := range stoneValue {
		prefix[i+1] = prefix[i] + value
	}

	dp := make([][]int, n)
	leftBest := make([][]int, n)
	rightBest := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		leftBest[i] = make([]int, n)
		rightBest[i] = make([]int, n)
		leftBest[i][i] = stoneValue[i]
		rightBest[i][i] = stoneValue[i]
	}

	sum := func(left, right int) int {
		return prefix[right+1] - prefix[left]
	}

	for length := 2; length <= n; length++ {
		for left := 0; left+length <= n; left++ {
			right := left + length - 1
			total := sum(left, right)
			low, high, split := left, right-1, left-1
			for low <= high {
				middle := low + (high-low)/2
				if sum(left, middle)*2 <= total {
					split = middle
					low = middle + 1
				} else {
					high = middle - 1
				}
			}

			best := 0
			if split >= left {
				best = leftBest[left][split]
				if sum(left, split)*2 == total {
					equalScore := sum(left, split) + dp[split+1][right]
					if equalScore > best {
						best = equalScore
					}
				}
			}
			if split+2 <= right && rightBest[split+2][right] > best {
				best = rightBest[split+2][right]
			}
			dp[left][right] = best

			leftValue := sum(left, right) + dp[left][right]
			leftBest[left][right] = leftBest[left][right-1]
			if leftValue > leftBest[left][right] {
				leftBest[left][right] = leftValue
			}

			rightValue := sum(left, right) + dp[left][right]
			rightBest[left][right] = rightBest[left+1][right]
			if rightValue > rightBest[left][right] {
				rightBest[left][right] = rightValue
			}
		}
	}

	return dp[0][n-1]
}

func main() {
	// result: 18
	// stoneValue := []int{6,2,3,4,5,5}

	// result: 0
	// stoneValue := []int{4}

	// result: 28
	stoneValue := []int{7, 7, 7, 7, 7, 7, 7}

	// result:
	// stoneValue := []int{}

	result := stoneGameV(stoneValue)
	fmt.Printf("result = %v\n", result)
}
