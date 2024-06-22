package main
import (
	"fmt"
)

func judgeSquareSum(c int) bool {
    m := map[int]bool{}
    for i := 0; i * i <= c; i++ {
        m[i * i] = true
    }
    for square, _ := range m {
        if m[c - square] {
            return true
        }
    }
    return false
}

func main() {
	// result: true
	// c := int(5)

	// result: false
	c := int(3)

	// result: 
	// c := int(0)

	result := judgeSquareSum(c)
	fmt.Printf("result = %v\n", result)
}

