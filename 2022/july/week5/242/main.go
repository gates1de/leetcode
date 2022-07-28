package main
import (
	"fmt"
	"sort"
	"reflect"
)

func isAnagram(s string, t string) bool {
	sChars := []rune(s)
	tChars := []rune(t)
	sort.Slice(sChars, func(a, b int) bool { return sChars[a] < sChars[b] })
	sort.Slice(tChars, func(a, b int) bool { return tChars[a] < tChars[b] })
	return reflect.DeepEqual(sChars, tChars)
}

func main() {
	// result: true
	// s := "anagram"
	// t := "margana"

	// result: false
	// s := "rat"
	// t := "car"

	// result: false
	// s := "s"
	// t := "t"

	// result: true
	s := "a"
	t := "a"

	// result: 
	// s := ""
	// t := ""

	result := isAnagram(s, t)
	fmt.Printf("result = %v\n", result)
}

