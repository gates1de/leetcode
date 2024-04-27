package main
import (
	"fmt"
	"math"
)

func findRotateSteps(ring string, key string) int {
	ringLen := len(ring)
	keyLen := len(key)

	current := make([]int, ringLen)
	prev := make([]int, ringLen)

	for keyIndex := keyLen - 1; keyIndex >= 0; keyIndex-- {
		for i, _ := range current {
			current[i] = math.MaxInt32
		}

		for ringIndex := 0; ringIndex < ringLen; ringIndex++ {
			for charIndex := 0; charIndex < ringLen; charIndex++ {
				if ring[charIndex] == key[keyIndex] {
					current[ringIndex] = min(
						current[ringIndex],
						1 + countSteps(ringIndex, charIndex, ringLen) + prev[charIndex],
					)
				}
			}
		}

		copy(prev, current)
	}

	return prev[0]
}

func countSteps(current int, next int, ringLen int) int {
	stepsBetween := abs(current - next)
	stepsAround := ringLen - stepsBetween
	return min(stepsBetween, stepsAround)
}


func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func main() {
	// result: 4
	// ring := "godding"
	// key := "gd"

	// result: 13
	ring := "godding"
	key := "godding"

	// result: 
	// ring := ""
	// key := ""

	result := findRotateSteps(ring, key)
	fmt.Printf("result = %v\n", result)
}

