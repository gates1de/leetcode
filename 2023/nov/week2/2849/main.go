package main
import (
	"fmt"
)

func isReachableAtTime(sx int, sy int, fx int, fy int, t int) bool {
	width := abs(sx, fx)
	height := abs(sy, fy)
	if width == 0 && height == 0 && t == 1 {
		return false
	}

	return t >= max(width, height)
}

func abs(a, b int) int {
	if b > a {
		return b - a
	}
	return a - b
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func main() {
	// result: true
	// sx := int(2)
	// sy := int(4)
	// fx := int(7)
	// fy := int(7)
	// t := int(6)

	// result: false
	sx := int(3)
	sy := int(1)
	fx := int(7)
	fy := int(3)
	t := int(3)

	// result: 
	// sx := int(0)
	// sy := int(0)
	// fx := int(0)
	// fy := int(0)
	// t := int(0)

	result := isReachableAtTime(sx, sy, fx, fy, t)
	fmt.Printf("result = %v\n", result)
}

