package main
import (
	"fmt"
)

func checkIfPangram(sentence string) bool {
	chars := make(map[byte]bool, 26)
	for _, char := range sentence {
		chars[byte(char)] = true
	}
	return len(chars) == 26
}

func main() {
	// result: true
	// sentence := "thequickbrownfoxjumpsoverthelazydog"

	// result: false
	sentence := "leetcode"

	// result: 
	// sentence := ""

	result := checkIfPangram(sentence)
	fmt.Printf("result = %v\n", result)
}

