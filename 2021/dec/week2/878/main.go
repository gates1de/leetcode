package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

func nthMagicalNumber(n int, a int, b int) int {
	if b < a {
		a, b = b, a
	}

	if b % a == 0 {
		return n * a % modulo
	}

	currentA := a
	currentB := b
	count := int(0)
	for currentA <= a * b {
		// fmt.Printf("currentA = %v, currentB = %v\n", currentA, currentB)
		count++

		if n == count {
			if currentA <= currentB {
				return currentA
			}
			return currentB
		}

		if currentA == currentB && currentA < a * b {
			count--
		}

		if currentA <= currentB {
			currentA += a
		} else {
			currentB += b
		}
	}

	result := int(0)
	for n > count {
		result += (a * b) % modulo
		n -= count
	}
	result += nthMagicalNumber(n, a, b)
	return result % modulo
}

func main() {
	// result: 2
	// n := int(1)
	// a := int(2)
	// b := int(3)

	// result: 6
	// n := int(4)
	// a := int(2)
	// b := int(3)

	// result: 10
	// n := int(5)
	// a := int(2)
	// b := int(4)

	// result: 8
	// n := int(3)
	// a := int(6)
	// b := int(4)

	// result: 33
	// n := int(13)
	// a := int(3)
	// b := int(11)

	// result: 999999993
	// n := int(1000000000)
	// a := int(2)
	// b := int(40000)

	// result: 8
	n := int(3)
	a := int(8)
	b := int(3)

	// result: 
	// n := int(0)
	// a := int(0)
	// b := int(0)

	result := nthMagicalNumber(n, a, b)
	fmt.Printf("result = %v\n", result)
}

