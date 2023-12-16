package main
import (
	"fmt"
)

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    dict := make(map[rune]int)
    for _, v := range s {
        dict[v]++
    }
    for _, v := range t {
        dict[v]--
    }
    for k, _ := range dict {
        if dict[k] != 0 {
            return false
        }
    }

    return true
}

func main() {
	// result: true
	// s := "anagram"
	// t := "nagaram"

	// result: false
	s := "rat"
	t := "car"

	// result: 
	// s := ""
	// t := ""

	result := isAnagram(s, t)
	fmt.Printf("result = %v\n", result)
}

