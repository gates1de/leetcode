package main
import (
	"fmt"
)

func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}

    rangeMax := int(1)
    tmp := n
    for tmp > 0 {
        rangeMax *= 2
        tmp /= 2
    }

    return rangeMax - n - 1
}

func main() {
	// result: 2
	// n := int(5)

	// result: 7
	// n := int(0)

	// result: 5
	n := int(10)

	// result: 
	// n := int(0)

	result := bitwiseComplement(n)
	fmt.Printf("result = %v\n", result)
}

