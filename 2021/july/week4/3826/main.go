package main
import (
	"fmt"
)

func findIntegers(n int) int {
	digit := int(0)
	m := map[int]int{}
	m[0] = 1
	m[1] = 2
	for i := n; i > 0; i /= 2  {
		digit++
	}
	// fmt.Printf("digit = %v\n", digit)
	last := int(0)
	// limit := pow(digit - 1) + pow(digit - 2)
	for i := 2; i <= digit; i++ {
		for j := 0; j < pow(i - 2); j++ {
			m[pow(i - 1) + j] = m[pow(i - 1) - 1] + m[j]
			// fmt.Printf("m[%v] = %v\n", pow(i - 1) + j, m[pow(i - 1) + j])
			last = m[pow(i - 1) + j]
			if pow(i - 1) + j >= n {
				return last
			}
		}
		for j := pow(i - 1) + 1; j < pow(i + 1); j++ {
			m[j] = last
			if j >= n {
				break
			}
		}
		// fmt.Printf("m = %v\n", m)
	}

	return m[n]
}

func pow(n int) int {
	if n < 0 {
		return 0
	}
	result := int(1)
	for i := 0; i < n; i++ {
		result *= 2
	}
	return result
}

func main() {
	// result: 5
	// n := int(5)

	// result: 2
	// n := int(1)

	// result: 3
	// n := int(2)

	// result: 6
	// n := int(8)

	// result: 14
	// n := int(32)

	// result: 22
	n := int(64)

	// result: 
	// n := int(0)

	result := findIntegers(n)
	fmt.Printf("result = %v\n", result)
}

