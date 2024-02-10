package main
import (
	"fmt"
)

const MAX = int(10000)

func numSquares(n int) int {
    if n <= 3 {
        return n
    }

    squares := make([]int, n + 1)
    for i, _ := range squares {
        squares[i] = MAX
    }
    squares[0] = 0

    for i := 1; i <= n; i++ {
        for j := 1; j * j <= i; j++ {
            squares[i] = min(squares[i], squares[i - j * j] + 1)
        }
    }

    return squares[n]
}

func min(a, b int) int {
    if b < a {
        return b
    }
    return a
}

func main() {
	// result: 3
	// n := int(12)

	// result: 2
	n := int(13)

	// result: 
	// n := int(0)

	result := numSquares(n)
	fmt.Printf("result = %v\n", result)
}

