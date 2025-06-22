package main
import (
	"fmt"
)

func divideString(s string, k int, fill byte) []string {
	result := make([]string, 0, len(s) / k + 1)
	chars := []byte(s)

	for len(chars) > k {
		subs := string(chars[:k])
		result = append(result, subs)
		chars = chars[k:]
	}

	if len(chars) > 0 {
		subs := make([]byte, 0, k)
		for _, char := range chars {
			subs = append(subs, char)
		}
		for i := len(subs); i < k; i++ {
			subs = append(subs, fill)
		}

		result = append(result, string(subs))
	}

	return result
}

func main() {
	// result: ["abc","def","ghi"]
	// s := "abcdefghi"
	// k := int(3)
	// fill := byte('x')

	// result: ["abc","def","ghi","jxx"]
	s := "abcdefghij"
	k := int(3)
	fill := byte('x')

	// result: []
	// s := ""
	// k := int(0)
	// fill := byte('')

	result := divideString(s, k, fill)
	fmt.Printf("result = %v\n", result)
}

