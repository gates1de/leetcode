package main
import (
	"fmt"
)

func gcdOfStrings(str1 string, str2 string) string {
	if str1 + str2 != str2 + str1 {
		return ""
	}

	gcdLen := gcd(len(str1), len(str2))
	return string(str1[:gcdLen])
}

func gcd(x int, y int) int {
	if y == 0 {
		return x
	}
	return gcd(y, x % y)
}

// Wrong Answer
func ngSolution(str1 string, str2 string) string {
	if len(str1) == 0 || len(str2) == 0 {
		return ""
	}

	short := str1
	long := str2
	if len(short) > len(long) {
		short, long = long, short
	}

	repeatChars := make([]rune, 0, len(short))
	repeatChars = append(repeatChars, rune(short[0]))
	repeatIndex := int(0)
	for i, char1 := range short {
		char2 := rune(long[i])
		if char1 != char2 {
			return ""
		}

		rc := repeatChars[repeatIndex]
		if rc != char1 {
			repeatChars = append(repeatChars, char1)
			repeatIndex = 0
		} else {
			repeatIndex++
			if repeatIndex >= len(repeatChars) {
				repeatIndex = 0
			}
		}
	}

	maxLen := len(long) / len(repeatChars)
	gcdLen := len(short) / len(repeatChars)
	for maxLen % gcdLen != 0 {
		if gcdLen % 2 == 1 {
			gcdLen = 1
			break
		}
		gcdLen /= 2
	}

	repeatStr := string(repeatChars)
	concatStr := ""
	for i := 0; i < maxLen; i++ {
		concatStr += repeatStr
	}
	fmt.Println(concatStr)
	if concatStr != long {
		return ""
	}

	result := ""
	for i := 0; i < gcdLen; i++ {
		result += repeatStr
	}


	return result
}

func main() {
	// result: "ABC"
	// str1 := "ABCABC"
	// str2 := "ABC"

	// result: "AB"
	// str1 := "ABABAB"
	// str2 := "ABAB"

	// result: ""
	// str1 := "LEET"
	// str2 := "CODE"

	// result: "TAUXX"
	str1 := "TAUXXTAUXXTAUXXTAUXXTAUXX"
	str2 := "TAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXXTAUXX"

	// result: 
	// str1 := ""
	// str2 := ""

	result := gcdOfStrings(str1, str2)
	fmt.Printf("result = %v\n", result)
}

