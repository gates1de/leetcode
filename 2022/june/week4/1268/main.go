package main
import (
	"fmt"
	"sort"
	"strings"
)

func suggestedProducts(products []string, searchWord string) [][]string {
	sort.Strings(products)
	result := make([][]string, len(searchWord))
	helper(0, searchWord, products, result)
	return result
}

func helper(index int, searchWord string, words []string, result [][]string) {
	if index >= len(searchWord) {
		return
	}

	targetWord := searchWord[:index + 1]
	newWords := []string{} 
	for _, word := range words {
		if strings.HasPrefix(word, targetWord) {
			newWords = append(newWords, word)
		}
		if word[0] > targetWord[0] {
			break
		}
	}

	if len(newWords) >= 3 {
		result[index] = newWords[:3]
	} else {
		result[index] = newWords
	}
	helper(index + 1, searchWord, newWords, result)
}

func main() {
	// result: [
	//	 ["mobile","moneypot","monitor"],
	// 	 ["mobile","moneypot","monitor"],
	// 	 ["mouse","mousepad"],
	// 	 ["mouse","mousepad"],
	// 	 ["mouse","mousepad"]
	// ]
	// products := []string{"mobile","mouse","moneypot","monitor","mousepad"}
	// searchWord := "mouse"

	// result: [["havana"],["havana"],["havana"],["havana"],["havana"],["havana"]]
	// products := []string{"havana"}
	// searchWord := "havana"

	// result: [["baggage","bags","banner"],["baggage","bags","banner"],["baggage","bags"],["bags"]]
	products := []string{"bags","baggage","banner","box","cloths"}
	searchWord := "bags"

	// result: 
	// products := []string{}
	// searchWord := ""

	result := suggestedProducts(products, searchWord)
	fmt.Printf("result = %v\n", result)
}

