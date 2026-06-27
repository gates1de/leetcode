package main

import (
	"fmt"
)

const modulo = int(1e9 + 7)

func zigZagArrays(n int, l int, r int) int {
	m := r - l + 1
	upPrev := make([]int, m + 2)
	downPrev := make([]int, m + 2)
	for v := 1; v <= m; v++ {
		upPrev[v] = 1
		downPrev[v] = 1
	}

	upCur := make([]int, m + 2)
	downCur := make([]int, m + 2)

	for length := 2; length <= n; length++ {
		prefix := 0
		for v := 1; v <= m; v++ {
			prefix += downPrev[v - 1]
			if prefix >= modulo {
				prefix -= modulo
			}
			upCur[v] = prefix
		}

		suffix := 0
		for v := m; v >= 1; v-- {
			suffix += upPrev[v + 1]
			if suffix >= modulo {
				suffix -= modulo
			}
			downCur[v] = suffix
		}

		upPrev, upCur = upCur, upPrev
		downPrev, downCur = downCur, downPrev
	}

	result := int(0)
	for v := 1; v <= m; v++ {
		result += upPrev[v]
		if result >= modulo {
			result -= modulo
		}
		result += downPrev[v]
		if result >= modulo {
			result -= modulo
		}
	}

	return result
}

func main() {
	// result: 2
	// n := int(3)
	// l := int(4)
	// r := int(5)

	// result: 10
	n := int(3)
	l := int(1)
	r := int(3)

	// result: 
	// n := int(0)
	// l := int(0)
	// r := int(0)

	result := zigZagArrays(n, l, r)
	fmt.Printf("result = %v\n", result)
}
