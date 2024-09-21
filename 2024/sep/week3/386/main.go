package main
import (
	"fmt"
)

func lexicalOrder(n int) []int {
	result := make([]int, 0)
	num := int(1)

	for i := 0; i < n; i++ {
		result = append(result, num)

		if num * 10 <= n {
			num *= 10
		} else {
			for num % 10 == 9 || num >= n {
				num /= 10
			}

			num++
		}
	}

	return result
}

func main() {
	// result: [1,10,11,12,13,2,3,4,5,6,7,8,9]
	// n := int(13)

	// result: [1,2]
	n := int(2)

	// result: []
	// n := int(0)

	result := lexicalOrder(n)
	fmt.Printf("result = %v\n", result)
}

