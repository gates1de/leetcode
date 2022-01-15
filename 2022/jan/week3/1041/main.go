package main
import (
	"fmt"
)

func isRobotBounded(instructions string) bool {
	direct := [2]int{0, 1}
	current := [2]int{0, 0}
	for i := 0; i < 4; i++ {
		for _, inst := range instructions {
			if inst == rune('G') {
				current[0] += direct[0]
				current[1] += direct[1]
				continue
			} else if inst == rune('L') {
				if direct[0] == 0 && direct[1] == 1 {
					direct[0] = -1
					direct[1] = 0
				} else if direct[0] == 1 && direct[1] == 0 {
					direct[0] = 0
					direct[1] = 1
				} else if direct[0] == 0 && direct[1] == -1 {
					direct[0] = 1
					direct[1] = 0
				} else if direct[0] == -1 && direct[1] == 0 {
					direct[0] = 0
					direct[1] = -1
				}
			} else if inst == rune('R') {
				if direct[0] == 0 && direct[1] == 1 {
					direct[0] = 1
					direct[1] = 0
				} else if direct[0] == 1 && direct[1] == 0 {
					direct[0] = 0
					direct[1] = -1
				} else if direct[0] == 0 && direct[1] == -1 {
					direct[0] = -1
					direct[1] = 0
				} else if direct[0] == -1 && direct[1] == 0 {
					direct[0] = 0
					direct[1] = 1
				}
			}
		}
	}

	if current == [2]int{0, 0} {
		return true
	}

	return false
}

func main() {
	// result: true
	// instructions := "GGLLGG"

	// result: false
	// instructions := "GG"

	// result: true
	// instructions := "GL"

	// result: false
	// instructions := "GGLRLGRRGGL"

	// result: true
	instructions := "GGLRLGRRGGLLG"

	// result: 
	// instructions := ""

	result := isRobotBounded(instructions)
	fmt.Printf("result = %v\n", result)
}

