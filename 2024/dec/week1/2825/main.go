package main
import (
	"fmt"
)

func canMakeSubsequence(str1 string, str2 string) bool {
	str2Index := int(0)
	str1Len := len(str1)
	str2Len := len(str2)

	for i := 0; i < str1Len && str2Index < str2Len; i++ {
		if str1[i] == str2[str2Index] ||
			str1[i] + 1 == str2[str2Index] ||
			str1[i] - 25 == str2[str2Index] {
				str2Index++
			}
	}

	return str2Index == str2Len
}

func main() {
	// result: true
	// str1 := "abc"
	// str2 := "ad"

	// result: true
	// str1 := "zc"
	// str2 := "ad"

	// result: false
	str1 := "ab"
	str2 := "d"

	// result: 
	// str1 := ""
	// str2 := ""

	result := canMakeSubsequence(str1, str2)
	fmt.Printf("result = %v\n", result)
}

