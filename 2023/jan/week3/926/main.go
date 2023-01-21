package main
import (
	"fmt"
)

func minFlipsMonoIncr(s string) int {
    zeroCount := int(0)
    for _, char := range s {
        if char == '0' {
            zeroCount++
        }
    }

    currentFlip := int(0)
    minFlip := len(s)
    for i := 0; i < len(s) && currentFlip < minFlip; i++ {
        if s[i] == '0' {
            zeroCount--
            continue
        }
        if currentFlip + zeroCount < minFlip {
            minFlip = currentFlip + zeroCount
        }
        currentFlip++
    }
    if currentFlip < minFlip {
        minFlip = currentFlip
    }
    return minFlip
}

func main() {
	// result: 1
	// s := "00110"

	// result: 2
	// s := "010110"

	// result: 2
	s := "00011000"

	// result: 
	// s := ""

	result := minFlipsMonoIncr(s)
	fmt.Printf("result = %v\n", result)
}

