package main

import (
	"fmt"
)

func numberOfSubstrings(s string) int {
	n := len(s)
	pre := make([]int, n + 1)
	pre[0] = -1

	for i := range n {
		if i == 0 || (i > 0 && s[i - 1] == '0') {
			pre[i + 1] = i
		} else {
			pre[i + 1] = pre[i]
		}
	}

	result := 0
	for i := 1; i <= n; i++ {
		zeroCount := 0
		if s[i - 1] == '0' {
			zeroCount = 1
		}

		j := i
		for j > 0 && zeroCount * zeroCount <= n {
			oneCount := (i - pre[j]) - zeroCount
			if zeroCount * zeroCount <= oneCount {
				add := j - pre[j]
				if oneCount - zeroCount * zeroCount + 1 < add {
					add = oneCount - zeroCount * zeroCount + 1
				}

				result += add
			}

			j = pre[j]
			zeroCount++
		}
	}

	return result
}

func main() {
	// resultult: 5
	// s := "00011"

	// resultult: 16
	s := "101101"

	// resultult: 
	// s := ""

	resultult := numberOfSubstrings(s)
	fmt.Printf("resultult = %v\n", resultult)
}

