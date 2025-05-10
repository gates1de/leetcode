package main
import (
	"fmt"
)

const modulo = int(1e9 + 7)

func numTilings(n int) int {
	m := make(map[int]int, 1000)
	for i := 1; i <= n; i++ {
		helper(i, m)
	}
	return m[n]
}

func helper(n int, m map[int]int) {
	if n <= 1 {
		m[n] = n
		return
	} else if n == 2 {
		m[n] = 2
		return
	} else if n == 3 {
		m[n] = 5
		return
	}

	m[n] = m[n - 1] * 2 + m[n - 3]
	m[n] %= modulo
}

func main() {
	// result: 5
	// n := int(3)

	// result: 1
	// n := int(1)

	// result: 2
	// n := int(2)

	// result: 258
	// n := int(8)

	// result: 603582422
	n := int(500)

	// result: 
	// n := int(0)

	result := numTilings(n)
	fmt.Printf("result = %v\n", result)
}
