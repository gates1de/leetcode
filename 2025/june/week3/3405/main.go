package main
import (
	"fmt"
)

const (
	modulo = int64(1e9 + 7)
	nmMax = int(100000)
)

func countGoodArrays(n int, m int, k int) int {
	fact := make([]int64, nmMax)
	invFact := make([]int64, nmMax)

	initFact(fact, invFact)
	return int(comb(n - 1, k, fact, invFact) * int64(m) % modulo * qpow(int64(m - 1), n - k - 1) % modulo)
}

func initFact(fact, invFact []int64) {
	if fact[0] != 0 {
		return
	}

	fact[0] = 1
	for i := 1; i < nmMax; i++ {
		fact[i] = fact[i - 1] * int64(i) % modulo
	}

	invFact[nmMax - 1] = qpow(fact[nmMax - 1], int(modulo - 2))
	for i := nmMax - 1; i > 0; i-- {
		invFact[i - 1] = invFact[i] * int64(i) % modulo
	}
}

func comb(n, m int, fact, invFact []int64) int64 {
	return fact[n] * invFact[m] % modulo * invFact[n - m] % modulo
}

func qpow(x int64, n int) int64 {
	result := int64(1)

	for n > 0 {
		if n & 1 == 1 {
			result = result * x % modulo
		}

		x = x * x % modulo
		n >>= 1
	}

	return result
}

func main() {
	// result: 4
	// n := int(3)
	// m := int(2)
	// k := int(1)

	// result: 6
	// n := int(4)
	// m := int(2)
	// k := int(2)

	// result: 2
	n := int(5)
	m := int(2)
	k := int(0)

	// result: 
	// n := int(0)
	// m := int(0)
	// k := int(0)

	result := countGoodArrays(n, m, k)
	fmt.Printf("result = %v\n", result)
}

