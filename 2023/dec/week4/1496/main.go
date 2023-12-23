package main
import (
	"fmt"
)

func isPathCrossing(path string) bool {
	x := int(0)
	y := int(0)
	visited := make(map[[2]int]bool)
	visited[[2]int{x, y}] = true

	for _, direction := range path {
		if direction == 'N' {
			y++
		} else if direction == 'S' {
			y--
		} else if direction == 'E' {
			x++
		} else if direction == 'W' {
			x--
		}

		if visited[[2]int{x, y}] {
			return true
		}

		visited[[2]int{x, y}] = true
	}

	return false
}

func main() {
	// result: false
	// path := "NES"

	// result: true
	path := "NESWW"

	// result: 
	// path := ""

	result := isPathCrossing(path)
	fmt.Printf("result = %v\n", result)
}

