package main
import (
	"fmt"
	"strings"
)

func wordSubsets(words1 []string, words2 []string) []string {
    BFreq := map[string]int{}
    result := []string{}
    countMax := int(0)
    for _, word := range words2 {
        for _, r := range word {
            char := string(r)
            count := strings.Count(word, char)
            if BFreq[char] != 0 {
                diff := count - BFreq[char]
                if diff > 0 {
                    BFreq[char] = count
                    countMax += diff
                }
            } else {
                BFreq[char] = count
                countMax += count
            }
        }
        if countMax > 10 {
            return result
        }
    }

TOP:
    for _, word := range words1 {
        if len(word) < countMax {
            continue
        }
        for char, _ := range BFreq {
            if strings.Count(word, char) < BFreq[char] {
                continue TOP
            }
        }
        result = append(result, word)
    }
    return result
}

const a = byte(97)

// Time Limit Exceeded
func ngSolution(words1 []string, words2 []string) []string {
	m := map[rune]byte{}
	convertedWords2 := make([][]byte, len(words2))
	for i, word2 := range words2 {
		convertedWords2[i] = convert(word2, m)
	}

	convertedWords1 := make([][]byte, len(words1))
	for i, word1 := range words1 {
		convertedWords1[i] = convert(word1, m)
	}

	validIndexes := make([]int, len(words1))
	for i, _ := range validIndexes {
		validIndexes[i] = i
	}

	for _, bytes2 := range convertedWords2 {
		newValidIndexes := []int{}

		for _, words1Index := range validIndexes {
			bytes1 := convertedWords1[words1Index]
			n := len(bytes2)
			if len(bytes1) < n {
				continue
			}

			if isSubset(bytes1, bytes2) {
				newValidIndexes = append(newValidIndexes, words1Index)
			}
		}

		validIndexes = newValidIndexes
	}

	result := make([]string, len(validIndexes))
	for i, word1Index := range validIndexes {
		result[i] = words1[word1Index]
	}
	return result
}

func convert(s string, m map[rune]byte) []byte {
	newS := ""
	for _, char := range s {
		if m[char] == 0 {
			m[char] = byte(len(m) + 1)
		}
		newS += string(a + m[char] - 1)
	}

	result := []byte(newS)
	return result
}

func isSubset(bytes1 []byte, bytes2 []byte) bool {
	m := map[byte]int{}
	for _, byte2 := range bytes2 {
		m[byte2]++
	}

	for _, byte1 := range bytes1 {
		if _, ok := m[byte1]; !ok {
			continue
		}

		m[byte1]--
		if m[byte1] == 0 {
			delete(m, byte1)
		}

		if len(m) == 0 {
			return true
		}
	}

	return false
}

func main() {
	// result: ["facebook","google","leetcode"]
	// words1 := []string{"amazon","apple","facebook","google","leetcode"}
	// words2 := []string{"e","o"}

	// result: ["apple","google","leetcode"]
	// words1 := []string{"amazon","apple","facebook","google","leetcode"}
	// words2 := []string{"l","e"}

	// result: ["facebook","leetcode"]
	// words1 := []string{"amazon","apple","facebook","google","leetcode"}
	// words2 := []string{"ec","o"}

	// result: ["leetcode"]
	words1 := []string{"amazon","apple","facebook","google","leetcode"}
	words2 := []string{"ece","o"}

	// result: 
	// words1 := []string{}
	// words2 := []string{}

	result := wordSubsets(words1, words2)
	fmt.Printf("result = %v\n", result)
}

