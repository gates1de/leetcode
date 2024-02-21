package main
import (
	"fmt"
)

func rangeBitwiseAnd(left int, right int) int {
    num := int(0)
    for n := 1; n <= left; n *= 2 {
        num = n
    }
    if right >= num * 2 {
        return 0
    }

    return num + rangeBitwiseAnd(left - num, right - num)
}

func main() {
	// result: 4
	// left := int(5)
	// right := int(7)

	// result: 0
	// left := int(0)
	// right := int(0)

	// result: 0
	left := int(1)
	right := int(2147483647)

	// result: 
	// left := int(0)
	// right := int(0)

	result := rangeBitwiseAnd(left, right)
	fmt.Printf("result = %v\n", result)
}

