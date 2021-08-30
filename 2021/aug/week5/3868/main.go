package main
import (
	"fmt"
)

func maxCount(m int, n int, ops [][]int) int {
	if len(ops) == 0 {
		return m * n
	}

	minA := m
	minB := n
	for i := 0; i < len(ops); i++ {
		o := ops[i]
		if minA > o[0] {
			minA = o[0]
		}
		if minB > o[1] {
			minB = o[1]
		}
	}

	return minA * minB
}

func main() {
	// result: 4
	// m := int(3)
	// n := int(3)
	// ops := [][]int{{2, 2}, {3, 3}}

	// result: 4
	// m := int(3)
	// n := int(3)
	// ops := [][]int{{2,2},{3,3},{3,3},{3,3},{2,2},{3,3},{3,3},{3,3},{2,2},{3,3},{3,3},{3,3}}

	// result: 9
	// m := int(3)
	// n := int(3)
	// ops := [][]int{}

	// result: 2
	m := int(8)
	n := int(6)
	ops := [][]int{{3, 2}, {3, 3}, {8, 1}, {3, 5}}

	// result: 
	// m := int(0)
	// n := int(0)
	// ops := [][]int{}

	result := maxCount(m, n, ops)
	fmt.Printf("result = %v\n", result)
}

