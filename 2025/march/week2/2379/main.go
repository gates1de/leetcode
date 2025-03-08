package main
import (
	"fmt"
	"math"
)

func minimumRecolors(blocks string, k int) int {
	left := int(0)
	numOfWhites := int(0)
	result := math.MaxInt32

	for right := 0; right < len(blocks); right++ {
		if blocks[right] == 'W' {
			numOfWhites++
		}

		if right - left + 1 == k {
			result = min(result, numOfWhites)

			if blocks[left] == 'W' {
				numOfWhites--
			}

			left++
		}
	}

	return result
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func main() {
	// result: 3
	// blocks := "WBBWWBBWBW"
	// k := int(7)

	// result: 0
	blocks := "WBWBBBW"
	k := int(2)

	// result: 
	// blocks := ""
	// k := int(0)

	result := minimumRecolors(blocks, k)
	fmt.Printf("result = %v\n", result)
}

