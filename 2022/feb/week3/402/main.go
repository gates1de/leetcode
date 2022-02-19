package main
import (
	"fmt"
	"strings"
)

func removeKdigits(num string, k int) string {
	zeroCount := int(0)
	for _, char := range num {
		if char == rune('0') {
			zeroCount++
		}
	}

	if len(num) - k <= zeroCount {
		return "0"
	}

	result := string(num[0])

	for i, char := range num[1:] {
		remain := len(num) - i - 1
		shortage := (len(num) - k) - len(result)
		if remain == shortage {
			result += num[i + 1:]
			break
		}

		insertIndex := int(-1)
		for j := 1; j <= len(result) && j <= remain - shortage; j++ {
			rChar := rune(result[len(result) - j])
			if char < rChar {
				insertIndex = len(result) - j
				continue
			}
			break
		}

		if insertIndex == -1 {
			result += string(char)
		} else {
			result = result[:insertIndex] + string(char)
		}

		if len(result) > len(num) - k {
			result = result[:len(num) - k]
		}
		// fmt.Printf("result = %v\n", result)
	}

	// fmt.Printf("after: result = %v\n", result)

	for _, char := range result {
		if char != rune('0') {
			break
		}
		result = strings.Replace(result, "0", "", 1)
	}

	if result == "" {
		return "0"
	}

	return result
}

func main() {
	// result: "1219"
	// num := "1432219"
	// k := int(3)

	// result: "200"
	// num := "10200"
	// k := int(1)

	// result: "0"
	// num := "10"
	// k := int(2)

	// result: "0"
	// num := "102"
	// k := int(2)

	// result: "1"
	num := "1020901"
	k := int(3)

	// result: 
	// num := ""
	// k := int(0)

	result := removeKdigits(num, k)
	fmt.Printf("result = %v\n", result)
}

