package main
import (
	"fmt"
)

func countAndSay(n int) string {
	result := "1"
	if n == 1 {
		return result
	}

	for i := 2; i <= n; i++ {
		chars := make([]byte, 0, n)
		count := int(1)
		tmp := result[0]

		for j := 1; j < len(result); j++ {
			currentChar := result[j]
			if currentChar == tmp {
				count++
			} else {
				chars = append(chars, []byte(fmt.Sprintf("%d", count))...)
				chars = append(chars, tmp)
				tmp = currentChar
				count = 1
			}
		}

		chars = append(chars, []byte(fmt.Sprintf("%d", count))...)
		chars = append(chars, tmp)
		result = string(chars)
	}

	return result
}

func main() {
	// result: "1211"
	// n := int(4)

	// result: "1"
	n := int(1)

	// result: ""
	// n := int(0)

	result := countAndSay(n)
	fmt.Printf("result = %v\n", result)
}

