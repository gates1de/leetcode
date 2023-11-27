package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

func knightDialer(n int) int {
	jumps := [][]int{
		{4, 6},
		{6, 8},
		{7, 9},
		{4, 8},
		{3, 9, 0},
		{},
		{1, 7, 0},
		{2, 6},
		{1, 3},
		{2, 4},
	}

	dp := make([]int, 10)
	prevDp := make([]int, 10)
	for i, _ := range prevDp {
		prevDp[i] = 1
	}

	for remain := 1; remain < n; remain++ {
		dp = make([]int, 10)

		for square := 0; square < 10; square++ {
			num := int(0)
			for _, nextSquare := range jumps[square] {
				num = (num + prevDp[nextSquare]) % modulo
			}

			dp[square] = num
		}

		prevDp = dp
	}

	result := int(0)
	for square := 0; square < 10; square++ {
		result = (result + prevDp[square]) % modulo
	}

	return result
}

func main() {
	// result: 10
	// n := int(1)

	// result: 20
	// n := int(2)

	// result: 136006598
	n := int(3131)

	// result: 
	// n := int(0)

	result := knightDialer(n)
	fmt.Printf("result = %v\n", result)
}

