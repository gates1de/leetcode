package main

import (
	"fmt"
)

func minFlips(s string) int {
	n := len(s)
	pre := make([][2]int, n)
	for i := range n {
		if i == 0 {
			pre[i][0] = check(s[i], 1)
			pre[i][1] = check(s[i], 0)
		} else {
			pre[i][0] = pre[i - 1][1] + check(s[i], 1)
			pre[i][1] = pre[i - 1][0] + check(s[i], 0)
		}
	}

	result := min(pre[n - 1][0], pre[n - 1][1])
	if n % 2 == 1 {
		suf := make([][2]int, n)
		for i := n - 1; i >= 0; i-- {
			if i == n - 1 {
				suf[i][0] = check(s[i], 1)
				suf[i][1] = check(s[i], 0)
			} else {
				suf[i][0] = suf[i + 1][1] + check(s[i], 1)
				suf[i][1] = suf[i + 1][0] + check(s[i], 0)
			}
		}

		for i := range n - 1 {
			result = min(result, pre[i][0]+suf[i + 1][0])
			result = min(result, pre[i][1]+suf[i + 1][1])
		}
	}

	return result
}

func check(ch byte, x int) int {
	if int(ch - '0') == x {
		return 1
	}

	return 0
}

func main() {
	// result: 2
	// s := "111000"

	// result: 0
	// s := "010"

	// result: 1
	s := "1110"

	// s := ""

	result := minFlips(s)
	fmt.Printf("result = %v\n", result)
}

