package main
import (
	"fmt"
	"strings"
)

func areSentencesSimilar(sentence1 string, sentence2 string) bool {
	s1Words := strings.Split(sentence1, " ")
	s2Words := strings.Split(sentence2, " ")
	n1 := len(s1Words)
	n2 := len(s2Words)
	start := int(0)
	end1 := n1 - 1
	end2 := n2 - 1

	if (n1 > n2) {
		return areSentencesSimilar(sentence2, sentence1)
	}

	for start < n1 && s1Words[start] == s2Words[start] {
		start++
	}

	for end1 >= 0 && s1Words[end1] == s2Words[end2] {
		end1--
		end2--
	}

	return end1 < start
}

func main() {
	// result: true
	// sentence1 := "My name is Haley"
	// sentence2 := "My Haley"

	// result: false
	// sentence1 := "of"
	// sentence2 := "A lot of words"

	// result: 
	sentence1 := "Eating right now"
	sentence2 := "Eating"

	// result: 
	// sentence1 := ""
	// sentence2 := ""

	result := areSentencesSimilar(sentence1, sentence2)
	fmt.Printf("result = %v\n", result)
}

