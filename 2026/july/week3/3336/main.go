package main

import (
	"fmt"
)

const modulo = int64(1e9 + 7)

func subsequencePairCount(nums []int) int {
	const maxValue = int(200)
	const stride = maxValue + 1

	size := stride * stride
	dp := make([]int64, size)
	next := make([]int64, size)

	var trans [maxValue + 1][maxValue + 1]uint8
	for num := 0; num <= maxValue; num++ {
		for g := 0; g <= maxValue; g++ {
			trans[num][g] = uint8(gcd(num, g))
		}
	}

	dp[0] = 1
	for _, num := range nums {
		for i := 0; i < size; i++ {
			next[i] = 0
		}

		for g1 := 0; g1 <= maxValue; g1++ {
			row := g1 * stride
			for g2 := 0; g2 <= maxValue; g2++ {
				cur := dp[row+g2]
				if cur == 0 {
					continue
				}

				idx := row + g2
				next[idx] += cur
				if next[idx] >= modulo {
					next[idx] -= modulo
				}

				ng1 := int(trans[num][g1])
				idx = ng1 * stride + g2
				next[idx] += cur
				if next[idx] >= modulo {
					next[idx] -= modulo
				}

				ng2 := int(trans[num][g2])
				idx = row + ng2
				next[idx] += cur
				if next[idx] >= modulo {
					next[idx] -= modulo
				}
			}
		}

		dp, next = next, dp
	}

	result := int64(0)
	for g := 1; g <= maxValue; g++ {
		result += dp[g * stride + g]
		if result >= modulo {
			result -= modulo
		}
	}

	return int(result % modulo)
}

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	// result: 10
	// nums := []int{1, 2, 3, 4}

	// result: 2
	// nums := []int{10, 20, 30}

	// result: 50
	nums := []int{1, 1, 1, 1}

	// result: 
	// nums := []int{}

	result := subsequencePairCount(nums)
	fmt.Printf("result = %v\n", result)
}
