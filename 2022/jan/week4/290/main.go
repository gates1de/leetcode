package main
import (
	"fmt"
	"reflect"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	patternMap := map[rune]int{}
	patternArr := make([]int, len(pattern))
	i := int(1)
	arrIndex := int(0)
	for _, char := range pattern {
		if patternMap[char] == 0 {
			patternMap[char] = i
			i++
		}

		patternArr[arrIndex] = patternMap[char]
		arrIndex++
	}

	sMap := map[string]int{}
	sArr := make([]int, len(words))
	i = int(1)
	arrIndex = int(0)

	for _, word := range words {
		if sMap[word] == 0 {
			sMap[word] = i
			i++
		}

		sArr[arrIndex] = sMap[word]
		arrIndex++
	}

	// fmt.Printf("patternArr = %v, sArr = %v\n", patternArr, sArr)
	return reflect.DeepEqual(patternArr, sArr)
}

func main() {
	// result: true
	// pattern := "abba"
	// s := "dog cat cat dog"

	// result: false
	// pattern := "abba"
	// s := "dog cat cat fish"

	// result: false
	// pattern := "aaaa"
	// s := "dog cat cat dog"

	// result: true
	// pattern := "abcdefg"
	// s := "a b c d e f g"

	// result: false
	// pattern := "xxcxx"
	// s := "cherry coconut banana cranberry custardapple"

	// result: true
	pattern := "xxcxx"
	s := "cherry cherry banana cherry cherry"

	// result: 
	// pattern := ""
	// s := ""

	result := wordPattern(pattern, s)
	fmt.Printf("result = %v\n", result)
}

