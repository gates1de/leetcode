package main
import (
	"fmt"
	"strconv"
)

func maximumSwap(num int) int {
	numStr := fmt.Sprintf("%d", num)
	n := len(numStr)
	maxDigitIndex := int(-1)
	swapIndex1 := int(-1)
	swapIndex2 := int(-1)

	for i := n - 1; i >= 0; i-- {
		if maxDigitIndex == -1 || numStr[i] > numStr[maxDigitIndex] {
			maxDigitIndex = i
		} else if numStr[i] < numStr[maxDigitIndex] {
			swapIndex1 = i
			swapIndex2 = maxDigitIndex
		}
	}

	numChars := []byte(numStr)
	if swapIndex1 != -1 && swapIndex2 != -1 {
		numChars[swapIndex1], numChars[swapIndex2] = numChars[swapIndex2], numChars[swapIndex1]
	}

	result, err := strconv.Atoi(string(numChars))
	if err != nil {
		return -1
	}

	return result
}

func main() {
	// result: 7236
	// num := int(2736)

	// result: 9973
	num := int(9973)

	// result: 
	// num := int(0)

	result := maximumSwap(num)
	fmt.Printf("result = %v\n", result)
}

