package main
import (
	"fmt"
)

func minFlipsMonoIncr(s string) int {
	zero := 0
	for _, v := range s {
		if v == '0' {
			zero++
		}
	}
	curFlip := 0
	minFlip := len(s)
	for i := 0; i < len(s) && curFlip < minFlip; i++ {
		if s[i] == '0' {
			zero--
			continue
		}
		if curFlip + zero < minFlip {
			minFlip = curFlip + zero
		}
		curFlip++
	}
	if curFlip < minFlip {
		minFlip = curFlip
	}
	return minFlip
}

// Slow & Use more memory
func mySolution(s string) int {
	zeroCounter := map[int]int{}
	oneCounter := map[int]int{}
	for i, digit := range s {
		if digit == rune('0') {
			zeroCounter[i] = zeroCounter[i - 1] + 1
			oneCounter[i] = oneCounter[i - 1]
		} else {
			oneCounter[i] = oneCounter[i - 1] + 1
			zeroCounter[i] = zeroCounter[i - 1]
		}
	}
	allZeroCount := zeroCounter[len(s) - 1]
	allOneCount := oneCounter[len(s) - 1]
	// fmt.Printf("zeroCounter = %v, allZeroCount = %v\n", zeroCounter, allZeroCount)
	// fmt.Printf("oneCounter = %v, allOneCount = %v\n", oneCounter, allOneCount)
	result := len(s)
	for i, _ := range s {
		flipCount := oneCounter[i] + allZeroCount - zeroCounter[i]
		// fmt.Printf("flipCount = %v\n", flipCount)

		if result > flipCount {
			result = flipCount
		}
	}
	if len(s) - allZeroCount < result {
		result = len(s) - allZeroCount
	}
	if len(s) - allOneCount < result {
		result = len(s) - allOneCount
	}
	return result
}

func main() {
	// result: 1
	// s := "00110"

	// result: 2
	// s := "010110"

	// result: 2
	// s := "00011000"

	// result: 22
	// s := "10100101011010110101010110011010010101010101010"

	// result: 1
	s := "11011"

	// result: 
	// s := ""

	result := minFlipsMonoIncr(s)
	fmt.Printf("result = %v\n", result)
}

