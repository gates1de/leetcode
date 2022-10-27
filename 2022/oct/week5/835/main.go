package main
import (
	"fmt"
)

func largestOverlap(img1 [][]int, img2 [][]int) int {
	n := len(img1)

	result := int(0)
	for yShift := 0; yShift < n; yShift++ {
		for xShift := 0; xShift < n; xShift++ {
			result = max(result, shiftAndCount(xShift, yShift, img1, img2))
			result = max(result, shiftAndCount(xShift, yShift, img2, img1))
		}
	}

	return result
}

func shiftAndCount(xShift int, yShift int, m [][]int, r [][]int) int {
	leftShiftCount := int(0)
	rightShiftCount := int(0)
	rRow := int(0)
	for mRow := yShift; mRow < len(m); mRow++ {
		rCol := int(0)
		for mCol := xShift; mCol < len(m); mCol++ {
			if m[mRow][mCol] == 1 && m[mRow][mCol] == r[rRow][rCol] {
				leftShiftCount++
			}
			if m[mRow][rCol] == 1 && m[mRow][rCol] == r[rRow][mCol] {
				rightShiftCount++
			}
			rCol++
		}
		rRow++
	}

    return max(leftShiftCount, rightShiftCount)
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// img1 := [][]int{{1,1,0},{0,1,0},{0,1,0}}
	// img2 := [][]int{{0,0,0},{0,1,1},{0,0,1}}

	// result: 1
	// img1 := [][]int{{1}}
	// img2 := [][]int{{1}}

	// result: 0
	img1 := [][]int{{0}}
	img2 := [][]int{{0}}

	// result: 
	// img1 := [][]int{}
	// img2 := [][]int{}

	result := largestOverlap(img1, img2)
	fmt.Printf("result = %v\n", result)
}

