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
    return obj
}

func (this *WordFilter) F(prefix string, suffix string) int {
    preSufString := string(prefix[0]) + string(suffix[len(suffix) - 1])
    filteredWords := this.words[preSufString]
    result := int(-1)
    for _, word := range filteredWords {
        wordPrefix := word[:len(prefix)]
        wordSuffix := word[len(word) - len(suffix):]
        if wordPrefix == prefix && suffix == wordSuffix && result < this.wordIndex[word] {
            result = this.wordIndex[word]
        }
    }
    return result
}

func main() {
	// result: [null, 0]
	words := []string{"apple"}
	fixes := [][]string{{"a", "e"}}

	// result: [null, 1]
	// words := []string{}
	// controls := [][]string{{"", ""}}

	obj := Constructor(words)
	for _, fix := range fixes {
		fmt.Printf("obj.F(%v, %v) = %v\n", fix[0], fix[1], obj.F(fix[0], fix[1]))
	}
}

