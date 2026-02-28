package main

import (
	"fmt"
)

const modulo = int(1e9 + 7)

func concatenatedBinary(n int) int {
    result := int(1)
    length := int(4)
    for i := 2; i <= n; i++ {
        if i == length {
            length *= 2
        }
        result = (result * length + i) % modulo
    }
    return result
}

func main() {
	// result: 1
	// n := int(1)

	// result: 27
	// n := int(3)

	// result: 505379714
	n := int(12)

	// result: 
	// n := int(0)

	result := concatenatedBinary(n)
	fmt.Printf("result = %v\n", result)
}

