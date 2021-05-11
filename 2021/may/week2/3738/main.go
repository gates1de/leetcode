package main
import (
	"fmt"
)

func countPrimes(n int) int {
	return eratosthenes(n)
}

func eratosthenes(n int) int {
	if n <= 1 {
		return 0
	}
	primes := make([]bool, n)
	for i := 2; i < n; i++ {
		primes[i] = true
	}
	for i := 2; i < n; i++ {
		if !primes[i] {
			continue
		}
		primes[i] = true
		multiple := i * i
		for multiple < n {
			primes[multiple] = false
			multiple += i
		}
	}
	// fmt.Printf("primes = %v\n", primes)
	result := int(0)
	for _, isPrime := range primes {
		if isPrime {
			result++
		}
	}
	return result
}

func main() {
	// result: 4
	// n := int(10)

	// result: 8
	// n := int(20)

	// result: 0
	// n := int(1)

	// result: 1
	n := int(3)

	// result: 
	// n := int()

	result := countPrimes(n)
	fmt.Printf("result = %v\n", result)
}

