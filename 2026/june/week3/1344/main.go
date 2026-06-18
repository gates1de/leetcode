package main

import (
	"fmt"
)

func angleClock(hour int, minutes int) float64 {
	hour %= 12

	hourAngle := float64(hour) * 30.0 + float64(minutes) * 0.5
	minuteAngle := float64(minutes) * 6.0
	diff := hourAngle - minuteAngle
	if diff < 0 {
		diff = -diff
	}
	if diff > 180.0 {
		diff = 360.0 - diff
	}

	return diff
}

func main() {
	// result: 165
	// hour := int(12)
	// minutes := int(30)

	// result: 75
	// hour := int(3)
	// minutes := int(30)

	// result: 7.5
	hour := int(3)
	minutes := int(15)

	// result: 
	// hour := int(0)
	// minutes := int(0)

	result := angleClock(hour, minutes)
	fmt.Printf("result = %v\n", result)
}
