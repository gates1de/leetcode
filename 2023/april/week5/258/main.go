package main
import (
	"fmt"
)

func addDigits(num int) int {
    result := int(10)
    for result >= 10 {
        result = 0
        for num > 0 {
            result += num % 10
            num /= 10
        }
        num = result
    }
    return result
}

func main() {
	// result: 2
	// num := int(38)

	// result: 0
	num := int(0)

	// result: 
	// num := int(0)

	result := addDigits(num)
	fmt.Printf("result = %v\n", result)
}

