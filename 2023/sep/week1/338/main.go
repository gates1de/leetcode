package main
import (
	"fmt"
)

func countBits(n int) []int {
    result := make([]int, n + 1)
    for i := 0; i <= n; i++ {
        result[i] = popcount(i)
    }
    return result
}

func popcount(x int) int {
    x = x - ((x >> 1) & 0x5555555555555555)
    x = (x & 0x3333333333333333) + ((x >> 2) & 0x3333333333333333)
    x = (x + (x >> 4)) & 0x0f0f0f0f0f0f0f0f
    x = x + (x >> 8)
    x = x + (x >> 16)
    x = x + (x >> 32)
    return x & 0x0000007f
}

func main() {
	// result: [0,1,1]
	// n := int(2)

	// result: [0,1,1,2,1,2]
	n := int(5)

	// result: 
	// n := int(0)

	result := countBits(n)
	fmt.Printf("result = %v\n", result)
}

