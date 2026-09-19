package main

import (
	"fmt"
)

func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
	closestX := xCenter
	if closestX < x1 {
		closestX = x1
	} else if closestX > x2 {
		closestX = x2
	}

	closestY := yCenter
	if closestY < y1 {
		closestY = y1
	} else if closestY > y2 {
		closestY = y2
	}

	dx := int64(xCenter - closestX)
	dy := int64(yCenter - closestY)
	return dx*dx+dy*dy <= int64(radius)*int64(radius)
}

func main() {
	// result: true
	// radius := int(1)
	// xCenter := int(0)
	// yCenter := int(0)
	// x1 := int(1)
	// y1 := int(-1)
	// x2 := int(3)
	// y2 := int(1)

	// result: false
	// radius := int(1)
	// xCenter := int(1)
	// yCenter := int(1)
	// x1 := int(1)
	// y1 := int(-3)
	// x2 := int(2)
	// y2 := int(-1)

	// result: true
	radius := int(1)
	xCenter := int(0)
	yCenter := int(0)
	x1 := int(-1)
	y1 := int(0)
	x2 := int(0)
	y2 := int(1)

	// result: true
	// radius := int(0)
	// xCenter := int(0)
	// yCenter := int(0)
	// x1 := int(0)
	// y1 := int(0)
	// x2 := int(0)
	// y2 := int(0)

	result := checkOverlap(radius, xCenter, yCenter, x1, y1, x2, y2)
	fmt.Printf("result = %v\n", result)
}
