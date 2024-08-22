package main
import (
	"fmt"
)

func findComplement(num int) int {
    rangeMax := int(1)
    n := num
    for n > 0 {
        rangeMax *= 2
        n /= 2
    }

    return rangeMax - num - 1
}

func main() {
	// result: 2
	// num := int(5)

	// result: 0
	num := int(1)

	// result: 
	// num := int(0)

	result := findComplement(num)
	fmt.Printf("result = %v\n", result)
}

