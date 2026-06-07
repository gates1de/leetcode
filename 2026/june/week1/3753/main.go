package main

import (
	"fmt"
)

type State struct {
	Ways int64
	Sum  int64
}

func totalWaviness(num1 int64, num2 int64) int64 {
	result := countUpTo(num2) - countUpTo(num1 - 1)
	return result
}

func dfs(digits []int, memo [][][][]State, seen [][][][]bool, pos int, prev2 int, prev1 int, started int, tight int) State {
	if pos == len(digits) {
		return State{Ways: int64(1), Sum: int64(0)}
	}

	if tight == 0 && seen[pos][prev2][prev1][started] {
		return memo[pos][prev2][prev1][started]
	}

	limitDigit := int(9)
	if tight == 1 {
		limitDigit = digits[pos]
	}

	result := State{Ways: int64(0), Sum: int64(0)}
	for digit := int(0); digit <= limitDigit; digit++ {
		nextTight := int(0)
		if tight == 1 && digit == limitDigit {
			nextTight = int(1)
		}

		nextPrev2 := prev2
		nextPrev1 := prev1
		nextStarted := started
		contribution := int64(0)

		if started == 0 {
			if digit == int(0) {
				nextPrev2 = 10
				nextPrev1 = 10
			} else {
				nextStarted = 1
				nextPrev2 = 10
				nextPrev1 = digit
			}
		} else if prev2 == 10 {
			nextPrev2 = prev1
			nextPrev1 = digit
		} else {
			if (prev1 > prev2 && prev1 > digit) || (prev1 < prev2 && prev1 < digit) {
				contribution = int64(1)
			}
			nextPrev2 = prev1
			nextPrev1 = digit
		}

		child := dfs(digits, memo, seen, pos+1, nextPrev2, nextPrev1, nextStarted, nextTight)
		result.Ways += child.Ways
		result.Sum += child.Sum + child.Ways*contribution
	}

	if tight == 0 {
		seen[pos][prev2][prev1][started] = true
		memo[pos][prev2][prev1][started] = result
	}

	return result
}

func countUpTo(limit int64) int64 {
	if limit < int64(0) {
		return int64(0)
	}

	digits := make([]int, 0)
	if limit == int64(0) {
		digits = append(digits, 0)
	} else {
		for limit > 0 {
			digits = append(digits, int(limit%10))
			limit /= 10
		}
		for left, right := int(0), len(digits) - 1; left < right; left, right = left + 1, right - 1 {
			digits[left], digits[right] = digits[right], digits[left]
		}
	}

	memo := make([][][][]State, len(digits)+1)
	seen := make([][][][]bool, len(digits) + 1)
	for i := range memo {
		memo[i] = make([][][]State, 11)
		seen[i] = make([][][]bool, 11)
		for j := range memo[i] {
			memo[i][j] = make([][]State, 11)
			seen[i][j] = make([][]bool, 11)

			for k := range memo[i][j] {
				memo[i][j][k] = make([]State, 2)
				seen[i][j][k] = make([]bool, 2)
			}
		}
	}

	return dfs(digits, memo, seen, int(0), 10, 10, 0, 1).Sum
}

func main() {
	// result: 3
	// num1 := int64(120)
	// num2 := int64(130)

	// result: 3
	// num1 := int64(198)
	// num2 := int64(202)

	// result: 2
	num1 := int64(4848)
	num2 := int64(4848)

	// result: 
	// num1 := int64(0)
	// num2 := int64(0)

	result := totalWaviness(num1, num2)
	fmt.Printf("result = %v\n", result)
}
