package main
import (
	"fmt"
)

func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	result := make([]int, n)
	freq := make([]int, n + 1)
	commonCount := int(0)

	for i := 0; i < n; i++ {
		freq[A[i]] += 1
		if freq[A[i]] == 2 {
			commonCount++
		}

		freq[B[i]] += 1
		if freq[B[i]] == 2 {
			commonCount++
		}

		result[i] = commonCount
	}

	return result
}

func main() {
	// result: [0,2,3,4]
	// A := []int{1,3,2,4}
	// B := []int{3,1,2,4}

	// result: [0,1,3]
	A := []int{2,3,1}
	B := []int{3,1,2}

	// result: []
	// A := []int{}
	// B := []int{}

	result := findThePrefixCommonArray(A, B)
	fmt.Printf("result = %v\n", result)
}

