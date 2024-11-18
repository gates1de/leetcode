package main
import (
	"fmt"
)

func decrypt(code []int, k int) []int {
	n := len(code)
	result := make([]int, n)

	if k == 0 {
		return result
	}

	for i, _ := range code {
		for j := 0; j < abs(k); j++ {
			index := i + j + 1

			if k < 0 {
				index = i - j - 1
			}

			if index < 0 {
				index += n
			} else if index >= n {
				index -= n
			}

			result[i] += code[index]
		}
	}

	return result
}

func abs(num int) int {
	if num <= 0 {
		return -num
	}
	return num
}

func main() {
	// result: [12,10,16,13]
	// code := []int{5,7,1,4}
	// k := int(3)

	// result: [0,0,0,0]
	// code := []int{1,2,3,4}
	// k := int(0)

	// result: [12,5,6,13]
	code := []int{2,4,9,3}
	k := int(-2)

	// result: []
	// code := []int{}
	// k := int(0)

	result := decrypt(code, k)
	fmt.Printf("result = %v\n", result)
}

