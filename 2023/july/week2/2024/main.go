package main
import (
	"fmt"
)

func maxConsecutiveAnswers(answerKey string, k int) int {
	result := int(0)
	counter := make(map[byte]int)
	for right := 0; right < len(answerKey); right++ {
		counter[answerKey[right]] += 1
		minor := min(counter['T'], counter['F'])

		if minor <= k {
			result++
		} else {
			counter[answerKey[right - result]] -= 1
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
	// result: 4
	// answerKey := "TTFF"
	// k := int(2)

	// result: 3
	// answerKey := "TFFT"
	// k := int(1)

	// result: 5
	answerKey := "TTFTTFTT"
	k := int(1)

	// result: 
	// answerKey := ""
	// k := int(0)

	result := maxConsecutiveAnswers(answerKey, k)
	fmt.Printf("result = %v\n", result)
}

