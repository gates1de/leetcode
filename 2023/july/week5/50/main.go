package main
import (
	"fmt"
)

func myPow(x float64, n int) float64 {
	return binaryExp(x, int64(n))
}

func binaryExp(x float64, n int64) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return float64(1) / binaryExp(x, -1 * n)
	}
       
	if n % 2 == 1 {
		return x * binaryExp(x * x, (n - 1) / 2)
	}

	return binaryExp(x * x, n / 2)
}

func main() {
	// result: 1024.00000
	// x := float64(2.00000)
	// n := int(10)

	// result: 9.26100
	// x := float64(2.10000)
	// n := int(3)

	// result: 0.25000
	x := float64(2.00000)
	n := int(-2)

	// result: 
	// x := float64(0)
	// n := int(0)

	result := myPow(x, n)
	fmt.Printf("result = %v\n", result)
}

