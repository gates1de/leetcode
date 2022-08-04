package main
import (
	"fmt"
)

func mirrorReflection(p int, q int) int {
	return helper(0, p, false, p, q)
}

func helper(x int, y int, isDown bool, p int, q int) int {
	if x == p && y == p {
		return 0
	} else if x == p && y == 0 {
		return 1
	} else if x == 0 && y == 0 {
		return 2
	}

	nextX := p
	if x == p {
		nextX = 0
	}

	nextY := y
	if isDown {
		nextY = y + q
		if nextY >= p {
			isDown = false
			nextY = p - (nextY - p)
		}
	} else {
		nextY = y - q
		if nextY <= 0 {
			isDown = true
			nextY = -nextY
		}
	}

	return helper(nextX, nextY, isDown, p, q)
}

func main() {
	// result: 2
	// p := int(2)
	// q := int(1)

	// result: 1
	// p := int(3)
	// q := int(1)

	// result: 1
	p := int(13)
	q := int(3)

	// result: 
	// p := int(0)
	// q := int(0)

	result := mirrorReflection(p, q)
	fmt.Printf("result = %v\n", result)
}

