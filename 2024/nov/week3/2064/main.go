package main
import (
	"fmt"
)

func minimizedMaximum(n int, quantities []int) int {
	left := int(0)
	right := int(0)

	for _, quantity := range quantities {
		if quantity > right {
			right = quantity
		}
	}

	for left < right {
		mid := (left + right) / 2

		if canDistribute(mid, quantities, n) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func canDistribute(val int, quantities []int, n int) bool {
	i := int(0)
	remaining := quantities[i]

	for j := 0; j < n; j++ {
		if remaining <= val {
			i++

			if i == len(quantities) {
				return true
			} else {
				remaining = quantities[i]
			}
		} else {
			remaining -= val
		}
	}

	return false
}

func main() {
	// result: 3
	// n := int(6)
	// quantities := []int{11,6}

	// result: 5
	// n := int(7)
	// quantities := []int{15,10,10}

	// result: 100000
	n := int(1)
	quantities := []int{100000}

	// result: 
	// n := int(0)
	// quantities := []int{}

	result := minimizedMaximum(n, quantities)
	fmt.Printf("result = %v\n", result)
}

