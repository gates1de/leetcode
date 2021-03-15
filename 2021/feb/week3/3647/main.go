package main
import (
	"fmt"
)

func brokenCalc(X int, Y int) int {
	if X == Y {
		return 0
	}
	if X > Y {
		return X - Y
	}
	return traversal(Y, 0, X)
}

func traversal(current int, count int, X int) int {
	if current == X {
		return count
	}
	if current > X {
		count++
		if current % 2 == 0 {
			return traversal(current / 2, count, X)
		} else {
			return traversal(current + 1, count, X)
		}
	}
	return count + X - current
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// X := 2
	// Y := 3

	// X := 5
	// Y := 8

	// X := 3
	// Y := 10

	// X := 1024
	// Y := 1

	X := 1
	Y := 1000000000

	result := brokenCalc(X, Y)
	fmt.Printf("result = %v\n", result)
}

