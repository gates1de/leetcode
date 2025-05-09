package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

func countBalancedPermutations(num string) int {
	total := int(0)
	n := len(num)
	counter := make([]int, 10)

	for _, char := range num {
			val := int(char - '0')
			counter[val]++
			total += val
	}

	if total % 2 != 0 {
			return 0
	}

	target := total / 2
	maxOdd := (n + 1) / 2
	comb := make([][]int, maxOdd + 1)
	for i := range comb {
			comb[i] = make([]int, maxOdd + 1)
			comb[i][i], comb[i][0] = 1, 1
			for j := 1; j < i; j++ {
					comb[i][j] = (comb[i - 1][j] + comb[i - 1][j - 1]) % modulo
			}
	}

	f := make([][]int, target + 1)
	for i := range f {
			f[i] = make([]int, maxOdd + 1)
	}
	f[0][0] = 1

	psum := int(0)
	totalSum := int(0)
	for i := 0; i <= 9; i++ {
			psum += counter[i]
			totalSum += i * counter[i]

			for oddCount := min(psum, maxOdd); oddCount >= max(0, psum- (n - maxOdd)); oddCount-- {
					evenCount := psum - oddCount

					for curr := min(totalSum, target); curr >= max(0, totalSum - target); curr-- {
							val := int(0)
							for j := max(0, counter[i] - evenCount); j <= min(counter[i], oddCount) && i * j <= curr; j++ {
									ways := comb[oddCount][j] * comb[evenCount][counter[i] - j] % modulo
									val = (val + ways * f[curr - i * j][oddCount - j] % modulo) % modulo
							}
							f[curr][oddCount] = val % modulo
					}
			}
	}

	return f[target][maxOdd]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 2
	// num := "123"

	// result: 1
	// num := "112"

	// result: 0
	num := "12345"

	// result: 
	// num := ""

	result := countBalancedPermutations(num)
	fmt.Printf("result = %v\n", result)
}

