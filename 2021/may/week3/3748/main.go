package main
import (
	"fmt"
	"math"
	"sort"
)

func minMoves2(nums []int) int {
    sort.Ints(nums)
    n := len(nums)
    if n <= 1 { return 0 }
    a, b := nums[n / 2], nums[n / 2 - 1]
    A, B := 0, 0
    for i := 0; i < n; i++ {
        A += abs(a, nums[i])
        B += abs(b, nums[i])
    }
    return min(A, B)
}

// Slow & Use more memory
func myAnswer(nums []int) int {
	m := map[int]int{}
	for _, num := range nums {
		sum := int(0)
		for _, n := range nums {
			sum += abs(num, n)
		}
		m[num] = sum
	}
	fmt.Printf("m = %v\n", m)

	result := math.MaxInt32
	for _, num := range m {
		if result > num {
			result = num
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

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 2
	// nums := []int{1, 2, 3}

	// result: 16
	nums := []int{1, 10, 2, 9}

	// result: 
	// nums := []int{}

	result := minMoves2(nums)
	fmt.Printf("result = %v\n", result)
}

