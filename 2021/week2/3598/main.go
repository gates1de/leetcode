package main
import (
	"fmt"
	"reflect"
	"strings"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if !contains(endWord, wordList) ||
		reflect.DeepEqual(wordList, []string{beginWord, endWord}) {
		return 0
	}
	if diff(beginWord, endWord) == 1 {
		return 2
	}

	result := int(1)
	currentWord := beginWord

	for {
		nextWord := currentWord

		fmt.Printf("wordList = %v, currentWord = %v\n", wordList, currentWord)
		for _, word := range wordList {
			fmt.Printf("nextWord = %v\n", nextWord)
			if beginWord == word {
				fmt.Printf("word is beginWord\n")
				continue
			}
			if endWord == word {
				fmt.Printf("word is endWord\n")
				continue
			}

			if diff(beginWord, word) == 1 && diff(endWord, word) == 1 {
				fmt.Printf("word is almost beginWord and endWord\n")
				nextWord = word
				break
			}

			if diff(currentWord, word) == 1 {
				fmt.Printf("diff is 1\n")
				nextWord = word
				break
			}
		}

		if nextWord == beginWord {
			fmt.Printf("Not changed nextWord from beginWord\n")
			return 0
		}
		if currentWord == nextWord {
			fmt.Printf("Not changed nextWord from currentWord\n")
			return 0
		}
		if diff(nextWord, endWord) == 1 {
			fmt.Printf("nextWord is almost endWord\n")
			result++
			break
		}

		currentWord = nextWord
		fmt.Printf("++ currentWord = %v\n", currentWord)
		wordList = removeByElement(wordList, currentWord)
		result++
	}

	// ++ is endWord count 
	result++
	return result
}

func removeByElement(s []string, target string) []string {
	result := make([]string, 0, len(s))
	for _, str := range s {
		if str != target {
			result = append(result, str)
		}
	}
	return result
}

func contains(word string, wordList []string) bool {
	for _, w := range wordList {
		if word == w {
			return true
		}
	}
	return false
}

func diff(word1 string, word2 string) int {
	count := int(0)
	characters1 := strings.Split(word1, "")
	characters2 := strings.Split(word2, "")
	for i, c1 := range characters1 {
		c2 := characters2[i]
		if c1 != c2 {
			count++
		}
	}
	return count
}

func main() {
	// beginWord := "hit"
	// endWord := "cog"
	// wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	// beginWord := "hot"
	// endWord := "dog"
	// wordList := []string{"hot", "dog", "dot"}
	// wordList := []string{"hot","cog","dog","tot","hog","hop","pot","dot"}

	// beginWord := "a"
	// endWord := "c"
	// wordList := []string{"a", "b", "c"}

	// beginWord := "hit"
	// endWord := "cog"
	// wordList := []string{"hot","dot","tog","cog"}

	beginWord := "teach"
	endWord := "place"
	wordList := []string{"peale","wilts","place","fetch","purer","pooch","peace","poach","berra","teach","rheum","peach"}

	result := ladderLength(beginWord, endWord, wordList)
	fmt.Printf("result = %v\n", result)
}

