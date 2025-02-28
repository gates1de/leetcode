package main
import (
	"fmt"
)

func shortestCommonSupersequence(str1 string, str2 string) string {
	str1Len := len(str1)
	str2Len := len(str2)

	dp := make([][]int, str1Len + 1)
	for row, _ := range dp {
		dp[row] = make([]int, str2Len + 1)
		dp[row][0] = row

		for col, _ := range dp[row] {
			dp[0][col] = col
		}
	}

	for row := 1; row <= str1Len; row++ {
		for col := 1; col <= str2Len; col++ {
			if str1[row - 1] == str2[col - 1] {
				dp[row][col] = dp[row - 1][col - 1] + 1
			} else {
				dp[row][col] = min(dp[row - 1][col], dp[row][col - 1]) + 1
			}
		}
	}

	chars := make([]byte, 0, str1Len + str2Len)
	row := str1Len
	col := str2Len

	for row > 0 && col > 0 {
		if str1[row - 1] == str2[col - 1] {
			chars = append(chars, str1[row - 1])
			row--
			col--
		} else if dp[row - 1][col] < dp[row][col - 1] {
			chars = append(chars, str1[row - 1])
			row--
		} else {
			chars = append(chars, str2[col - 1])
			col--
		}
	}

	for row > 0 {
		chars = append(chars, str1[row - 1])
		row--
	}
	for col > 0 {
		chars = append(chars, str2[col - 1])
		col--
	}

	reverse(chars)
	return string(chars)
}

func reverse(chars []byte) {
	for i := 0; i < len(chars) / 2; i++ {
		chars[i], chars[len(chars) - i - 1] = chars[len(chars) - i - 1], chars[i]
	}
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: "cabac"
	// str1 := "abac"
	// str2 := "cab"

	// result: "aaaaaaaa"
	str1 := "aaaaaaaa"
	str2 := "aaaaaaaa"

	// result: ""
	// str1 := ""
	// str2 := ""

	result := shortestCommonSupersequence(str1, str2)
	fmt.Printf("result = %v\n", result)
}

