package main
import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func findSubstring(s string, words []string) []int {
	result := []int{}
    if len(words) == 0 {
		return result
	}

    wordLen := len(words[0])
    for offset := 0; offset < wordLen; offset++ {
        begin, end := offset, offset
        ws := append(make([]string, 0), words...)

        TOP:
        for end + wordLen <= len(s) {
            for i, word := range ws {
                if word == s[end:end + wordLen] {
                    end += wordLen
                    ws = append(ws[:i], ws[i + 1:]...)
                    if len(ws) == 0 {
                        result = append(result, begin)
                    }
                    continue TOP
                }
            }

            if begin < end {
                ws = append(ws, s[begin:begin + wordLen])
            } else {
                end += wordLen
            }

            begin += wordLen
        }
    }

    return result
}

const A = byte(65)

// Wrong Answer
func ngSolution(s string, words []string) []int {
	m := map[string]byte{}
	convertedWords := []byte{}
	for _, word := range words {
		if val, ok := m[word]; ok {
			convertedWords = append(convertedWords, val)
			continue
		}
		m[word] = A + byte(len(m))
		convertedWords = append(convertedWords, m[word])
	}
	sort.Slice(convertedWords, func(a, b int) bool {
		return convertedWords[a] < convertedWords[b]
	})

	for word, b := range m {
		s = strings.ReplaceAll(s, word, string(b))
	}

	wordsLen := len(words)
	wordLen := len(words[0])
	sLen := len(s)

	indexMap := map[int]int{}
	index := int(0)
	for i := 0; i < sLen; i++ {
		char := s[i]
		if A <= char && char <= A + 30 {
			index += wordLen
		} else {
			index++
		}
		indexMap[i] = index
	}

	result := []int{}
	for i := 0; i <= sLen - wordsLen; i++ {
		substr := []byte(s[i:i + wordsLen])
		sort.Slice(substr, func(a, b int) bool { return substr[a] < substr[b] })
		if reflect.DeepEqual(convertedWords, substr) {
			result = append(result, indexMap[i - 1])
		}
	}

	return result
}

func main() {
	// result: [0,9]
	// s := "barfoothefoobarman"
	// words := []string{"foo","bar"}

	// result: []
	// s := "wordgoodgoodgoodbestword"
	// words := []string{"word","good","best","word"}

	// result: [6,9,12]
	// s := "barfoofoobarthefoobarman"
	// words := []string{"bar","foo","the"}

	// result: [7,10,13]
	// s := "abarfoofoobarthefoobarman"
	// words := []string{"bar","foo","the"}

	// result: [8]
	// s := "wordgoodgoodgoodbestword"
	// words := []string{"word","good","best","good"}

	// result: [0,1,2,3,4,5,6,7,8,9,10]
	s := "aaaaaaaaaaaaaa"
	words := []string{"aa","aa"}

	// result: 
	// s := ""
	// words := []string{}

	result := findSubstring(s, words)
	fmt.Printf("result = %v\n", result)
}

