package main

import (
	"fmt"
	"sort"
)

func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
	sort.Ints(hBars)
	sort.Ints(vBars)
	hmax := int(1)
	vmax := int(1)
	hcur := int(1)
	vcur := int(1)

	for i := 1; i < len(hBars); i++ {
		if hBars[i] == hBars[i - 1] + 1 {
			hcur++
		} else {
			hcur = 1
		}
		hmax = max(hmax, hcur)
	}

	for i := 1; i < len(vBars); i++ {
		if vBars[i] == vBars[i - 1] + 1 {
			vcur++
		} else {
			vcur = 1
		}
		vmax = max(vmax, vcur)
	}

	side := min(hmax, vmax) + 1
	return side * side
}

func main() {
	// result: 4
	// n := int(2)
	// m := int(1)
	// hBars := []int{2,3}
	// vBars := []int{2}

	// result: 4
	// n := int(1)
	// m := int(1)
	// hBars := []int{2}
	// vBars := []int{2}

	// result: 4
	n := int(2)
	m := int(3)
	hBars := []int{2,3}
	vBars := []int{2,4}

	// result: 
	// n := int(0)
	// m := int(0)
	// hBars := []int{}
	// vBars := []int{}

	result := maximizeSquareHoleArea(n, m, hBars, vBars)
	fmt.Printf("result = %v\n", result)
}

