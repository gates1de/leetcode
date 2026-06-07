package main

import (
	"fmt"
	"sort"
)

func asteroidsDestroyed(mass int, asteroids []int) bool {
	current := int64(mass)
	sort.Ints(asteroids)
	for _, asteroid := range asteroids {
		if current < int64(asteroid) {
			return false
		}

		current += int64(asteroid)
	}

	return true
}

func main() {
	// result: true
	// mass := int(10)
	// asteroids := []int{3,9,19,5,21}

	// result: false
	mass := int(5)
	asteroids := []int{4,9,23,4}

	// result: false
	// mass := int(0)
	// asteroids := []int{}

	result := asteroidsDestroyed(mass, asteroids)
	fmt.Printf("result = %v\n", result)
}
