package main

import (
	"fmt"
)

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return rec1[0] < rec2[2] && rec2[0] < rec1[2] &&
		rec1[1] < rec2[3] && rec2[1] < rec1[3]
}

func main() {
	// result: true
	// rec1 := []int{0,0,2,2}
	// rec2 := []int{1,1,3,3}

	// result: false
	// rec1 := []int{0,0,1,1}
	// rec2 := []int{1,0,2,1}

	// result: false
	rec1 := []int{0, 0, 1, 1}
	rec2 := []int{2, 2, 3, 3}

	// result: false
	// rec1 := []int{}
	// rec2 := []int{}

	result := isRectangleOverlap(rec1, rec2)
	fmt.Printf("result = %v\n", result)
}
