package main
import (
	"fmt"
)

const (
	modulo = int(1e9 + 7)
	maxN = int(10010)
	maxP = int(15)
)

func idealArrays(n int, maxValue int) int {
	counter := make([][]int, maxN + maxP)
	sieve := make([]int, maxN)
	ps := make([][]int, maxN)

	for i, _ := range counter {
		counter[i] = make([]int, maxP + 1)
	}
	for i, _ := range ps {
		ps[i] = make([]int, 0)
	}

	if counter[0][0] != 0 {
		return -1
	}

	for i := 2; i < maxN; i++ {
		if sieve[i] == 0 {
			for j := i; j < maxN; j += i {
				if sieve[j] == 0 {
					sieve[j] = i
				}
			}
		}
	}

	for i := 2; i < maxN; i++ {
		x := i
		for x > 1 {
			p := sieve[x]
			count := int(0)
			for x % p == 0 {
				x /= p
				count++
			}

			ps[i] = append(ps[i], count)
		}
	}

	counter[0][0] = 1

	for i := 1; i < maxN + maxP; i++ {
		counter[i][0] = 1
		for j := 1; j <= maxP && j <= i; j++ {
			counter[i][j] = (counter[i - 1][j] + counter[i - 1][j - 1]) % modulo
		}
	}

	result := int(0)
	for x := 1; x <= maxValue; x++ {
		mul := int(1)
		for _, p := range ps[x] {
			mul = mul * counter[n + p - 1][p] % modulo
		}

		result = (result + mul) % modulo
	}

	return result
}

func main() {
	// result: 10
	// n := int(2)
	// maxValue := int(5)

	// result: 11
	n := int(5)
	maxValue := int(3)

	// result: 
	// n := int(0)
	// maxValue := int(0)

	result := idealArrays(n, maxValue)
	fmt.Printf("result = %v\n", result)
}

