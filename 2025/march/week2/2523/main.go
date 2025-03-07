package main
import (
	"fmt"
	"math"
)

func closestPrimes(left int, right int) []int {
	primeNumbers := make([]int, 0)
	for candidate := left; candidate <= right; candidate++ {
		if candidate % 2 == 0 && candidate > 2 {
			continue
		}

		if isPrime(candidate) {
			if len(primeNumbers) > 0 && candidate <= primeNumbers[len(primeNumbers) - 1] + 2 {
				return []int{primeNumbers[len(primeNumbers) - 1], candidate}
			}
			primeNumbers = append(primeNumbers, candidate)
		}
	}

	if len(primeNumbers) < 2 {
		return []int{-1, -1}
	}

	result := []int{-1, -1}
	minDiff := int(1000000)
	for i := 1; i < len(primeNumbers); i++ {
		diff := primeNumbers[i] - primeNumbers[i - 1]
		if diff < minDiff {
			minDiff = diff
			result = []int{primeNumbers[i - 1], primeNumbers[i]}
		}
	}

	return result
}

func isPrime(num int) bool {
	if num == 1 {
		return false
	}

	for div := 2; div <= int(math.Sqrt(float64(num))); div++ {
		if num % div == 0 {
			return false
		}
	}

	return true
}

func main() {
	// result: [11,13]
	// left := int(10)
	// right := int(19)

	// result: [-1,-1]
	left := int(4)
	right := int(6)

	// result: []
	// left := int(0)
	// right := int(0)

	result := closestPrimes(left, right)
	fmt.Printf("result = %v\n", result)
}

