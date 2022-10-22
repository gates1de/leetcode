package main
import (
	"fmt"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	pre := countAndSay(n - 1) + " "
	result := ""

	preChar := rune(pre[0])
	count := int(1)
	for _, char := range pre[1:] {
		if preChar == char {
			count++
			continue
		}

		result += fmt.Sprintf("%v%v", count, string(preChar))
		preChar = char
		count = 1
	}

	return result
}

func main() {
	// result: "1"
	// n := int(1)

	// result: "1211"
	// n := int(4)

	// result: "13211311123113112211"
	n := int(10)

	// result: 
	// n := int(0)

	result := countAndSay(n)
	fmt.Printf("result = %v\n", result)
}

