package main
import (
	"fmt"
)

type WordFilter struct {
	words map[string][]string
	wordIndex map[string]int
}


func Constructor(words []string) WordFilter {
	obj := WordFilter{words: map[string][]string{}, wordIndex: map[string]int{}}
	for i := len(words) - 1; i >= 0; i-- {
		word := words[i]
		if obj.wordIndex[word] > 0 {
			continue
		}
		sideString := string(word[0]) + string(word[len(word) - 1])
		obj.words[sideString] = append(obj.words[sideString], word)
		obj.wordIndex[word] = i
	}
	// fmt.Printf("obj = %v\n", obj)
	return obj
}


func (this *WordFilter) F(prefix string, suffix string) int {
	preSufString := string(prefix[0]) + string(suffix[len(suffix) - 1])
	filteredWords := this.words[preSufString]
	// fmt.Printf("filteredWords = %v\n", filteredWords)
	result := int(-1)
	for _, word := range filteredWords {
		wordPrefix := word[:len(prefix)]
		wordSuffix := word[len(word) - len(suffix):]
		// fmt.Printf("wordPrefix = %v, wordSuffix = %v\n", wordPrefix, wordSuffix)
		if wordPrefix == prefix && suffix == wordSuffix && result < this.wordIndex[word] {
			result = this.wordIndex[word]
		}
	}
	return result
}


/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */

func main() {
	// result: 0
	// words := []string{"apple"}
	// prefix := "a"
	// suffix := "e"

	// result: 5
	words := []string{"cabaabaaaa","ccbcababac","bacaabccba","bcbbcbacaa","abcaccbcaa","accabaccaa","cabcbbbcca","ababccabcb","caccbbcbab","bccbacbcba"}
	prefix := "a"
	suffix := "aa"

	// result: 
	// words := []string{}
	// prefix := ""
	// suffix := ""

	obj := Constructor(words)
	result := obj.F(prefix, suffix)
	fmt.Printf("result = %v\n", result)
}

