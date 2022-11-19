package main
import (
	"fmt"
)

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	aRectSize := (ax2 - ax1) * (ay2 - ay1)
	bRectSize := (bx2 - bx1) * (by2 - by1)
	if ax1 > bx2 || ax2 < bx1 || ay1 > by2 || ay2 < by1 {
		return aRectSize + bRectSize
	}

	insideRectLeft := int(0)
	insideRectRight := int(0)
	insideRectTop := int(0)
	insideRectBottom := int(0)

	if ax1 < bx1 {
		insideRectLeft = bx1
		insideRectRight = ax2
		if bx2 < ax2 {
			insideRectRight = bx2
		}
	} else {
		insideRectLeft = ax1
		insideRectRight = bx2
		if ax2 < bx2 {
			insideRectRight = ax2
		}
	}

	if ay1 < by1 {
		insideRectBottom = by1
		insideRectTop = ay2
		if by2 < ay2 {
			insideRectTop = by2
		}
	} else {
		insideRectBottom = ay1
		insideRectTop = by2
		if ay2 < by2 {
			insideRectTop = ay2
		}
	}

	insideRectSize := (insideRectRight - insideRectLeft) * (insideRectTop - insideRectBottom)
	return aRectSize + bRectSize - insideRectSize
}

func main() {
	// result: 45
	// ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 := int(-3), int(0), int(3), int(4), int(0), int(-1), int(9), int(2)

	// result: 16
	ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 := int(-2), int(-2), int(2), int(2), int(-2), int(-2), int(2), int(2)

	// result: 
	// ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 := int(0), int(0), int(0), int(0), int(0), int(0), int(0), int(0)

	result := computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2)
	fmt.Printf("result = %v\n", result)
}

