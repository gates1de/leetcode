package main
import (
	"fmt"
)

func numDecodings(s string) int {
    if len(s) == 0 || s[0] == '0' {
        return 0
    } else if len(s) == 1 {
        return 1
    }

    result := int(1)
    prevCount := int(1)
    for i := 1; i < len(s); i++ {
        currCount := 0
        if s[i - 1] == '1' || (s[i - 1] == '2' && s[i] < '7') {
            currCount += prevCount
        }

        if s[i] > '0' {
            currCount += result
        }

        prevCount = result
        result = currCount
    }

    return result
}

func main() {
	// result: 2
	// s := "12"

	// result: 3
	// s := "226"

	// result: 0
	// s := "06"

	// result: 0
	// s := "1203401514131902130"

	// result: 4
	// s := "12031419"

	// result: 1836311903
	s := "111111111111111111111111111111111111111111111"

	// result: 
	// s := ""

	result := numDecodings(s)
	fmt.Printf("result = %v\n", result)
}

