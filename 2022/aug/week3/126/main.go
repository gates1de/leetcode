package main
import (
	"fmt"
	"math"
	"sort"
)

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	adjacencyList := make(map[string][]string)
	for _, word := range wordList {
		for i := 0; i < len(word); i++ {
			pattern := fmt.Sprintf("%s.%s", word[:i], word[i + 1:])
			adjacencyList[pattern] = append(adjacencyList[pattern], word)
		}
	}

	reverseDict := map[string][]string{beginWord: {}}

	q := []string{beginWord}

	for len(q) > 0  {
		n := len(q)
		visitedThisLevel := make(map[string][]string)
		for i := 0; i < n; i++ {
			word := q[0]
			q = q[1:]
			for i := 0; i < len(word); i++ {
				pattern := fmt.Sprintf("%s.%s", word[:i], word[i + 1:])
				for _, nextWord := range adjacencyList[pattern] {
					if _, ok1 := reverseDict[nextWord]; !ok1 {
						if _, ok2 := visitedThisLevel[nextWord]; !ok2 {
							visitedThisLevel[nextWord] = []string{word}
							q = append(q, nextWord)
						} else {
							visitedThisLevel[nextWord] = append(visitedThisLevel[nextWord], word)
						}
					}
				}
			}
		}

		for k, v := range visitedThisLevel {
			reverseDict[k] = append(reverseDict[k], v...)
		}
	}

	if !contains(wordList, endWord) {
		return [][]string{}
	}

	return dfs(endWord, beginWord, reverseDict)
}

func contains(seq []string, s string) bool {
	for _, v := range seq {
		if v == s {
			return true
		}
	}
	return false
}

func dfs(node string, beginWord string, reverseDict map[string][]string) [][]string {
	if node == beginWord {
		return [][]string{{beginWord}}
	}
	if _, ok := reverseDict[node]; !ok {
		return [][]string{}
	}
	result := [][]string{}

	for _, parent := range reverseDict[node] {
		parentPathes := dfs(parent, beginWord, reverseDict)
		for i := 0; i < len(parentPathes); i++ {
			result = append(result, parentPathes[i])
		}
		for i := len(result) - 1; i > len(result) - len(parentPathes) - 1; i-- {
			result[i] = append(result[i], node)
		}
	}

	return result
}

// Time Limit Exceeded
func ngSolution(beginWord string, endWord string, wordList []string) [][]string {
	result := [][]string{}
	beginWordIndex := int(-1)
	for i, word := range wordList {
		if beginWord == word {
			beginWordIndex = i
			break
		}
	}

	minPath := math.MaxInt32
	if beginWordIndex == -1 {
		visited := make([]bool, len(wordList) + 1)
		wordList = append([]string{beginWord}, wordList...)
		visited[0] = true
		helper([]string{beginWord}, 0, endWord, wordList, visited, &result, &minPath)
	} else {
		visited := make([]bool, len(wordList))
		visited[beginWordIndex] = true
		helper([]string{beginWord}, beginWordIndex, endWord, wordList, visited, &result, &minPath)
	}

	if len(result) == 0 {
		return result
	}

	sort.Slice(result, func(a, b int) bool {
		return len(result[a]) < len(result[b])
	})

	minPath = len(result[0])
	for i, words := range result {
		if len(words) <= minPath {
			continue
		}
		return result[:i]
	}

	return result
}

func helper(
	words []string,
	wordIndex int,
	endWord string,
	wordList []string,
	visited []bool,
	result *[][]string,
	minPath *int,
) {
	currentWord := wordList[wordIndex]
	if currentWord == endWord {
		*result = append(*result, words)
		if *minPath >= len(words) {
			*minPath = len(words)
		}
		return
	}

	if *minPath <= len(words) {
		return
	}

	visited[wordIndex] = true
	for i, word := range wordList {
		if visited[i] {
			continue
		}

		if isValid(currentWord, word) {
			newVisited := make([]bool, len(visited))
			copy(newVisited, visited)
			newWords := make([]string, len(words) + 1)
			copy(newWords, words)
			newWords[len(words)] = word
			helper(newWords, i, endWord, wordList, newVisited, result, minPath)
		}
	}
}

func isValid(word1 string, word2 string) bool {
	diffCount := int(0)
	for i, char1 := range word1 {
		char2 := rune(word2[i])
		if char1 != char2 {
			diffCount++
		}
	}

	return diffCount <= 1
}

func main() {
	// result: [["hit","hot","dot","dog","cog"],["hit","hot","lot","log","cog"]]
	// beginWord := "hit"
	// endWord := "cog"
	// wordList := []string{"hot","dot","dog","lot","log","cog"}

	// result: []
	// beginWord := "hit"
	// endWord := "cog"
	// wordList := []string{"hot","dot","dog","lot","log"}

	// result: [["hot","dot","dog"],["hot","hog","dog"]]
	beginWord := "hot"
	endWord := "dog"
	wordList := []string{"hot","cog","dog","tot","hog","hop","pot","dot"}

	// result: 
	// beginWord := ""
	// endWord := ""
	// wordList := []string{}

	result := findLadders(beginWord, endWord, wordList)
	fmt.Printf("result = %v\n", result)
}

