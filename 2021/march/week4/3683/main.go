package main
import (
	"fmt"
	"sort"
)

func advantageCount(A []int, B []int) []int {
	if len(A) == 0 || len(B) == 0 {
		return B
	}

	sort.Slice(A, func (a, b int) bool { return A[a] < A[b] })
	result := make([]int, len(A))
	for i, _ := range result {
		result[i] = -1
	}
	for i, b := range B {
		for j, a := range A {
			if result[i] >= 0 && a >= b {
				break
			}
			if a <= b {
				continue
			}
			result[i] = a
			A[j] = -1
		}
	}

	for i, b := range B {
		if result[i] < 0 {
			minDiff := int(1000000000)
			changeIndexA := int(-1)
			for j, a := range A {
				if a < 0 {
					continue
				}
				diff := abs(a, b)
				if diff <= minDiff {
					minDiff = diff
					changeIndexA = j
				}
				break
			}
			if changeIndexA >= 0 {
				result[i] = A[changeIndexA]
				A[changeIndexA] = -1
			}
		}
	}
	return result
}

func abs(a, b int) int {
	if b > a {
		return b - a
	}
	return a - b
}

func main() {
	// A := []int{2, 7, 11, 15}
	// B := []int{1, 10, 4, 11}

	// A := []int{12, 24, 8, 32}
	// B := []int{13, 25, 32, 11}

	A := []int{5621, 1743, 5532, 3549, 9581}
	B := []int{913, 9787, 4121, 5039, 1481}

	result := advantageCount(A, B)
	fmt.Printf("result = %v\n", result)
}

