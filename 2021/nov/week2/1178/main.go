package main
import (
	"fmt"
)

func findNumOfValidWords(words []string, puzzles []string) []int {
    freqWords := make(map[int]int)
    for _, word := range words {
        b := bits(word)
        if _, ok := freqWords[b]; !ok {
            freqWords[b] = 1
        } else {
            freqWords[b] += 1
        }
    }

    result := make([]int, len(puzzles))
    for i, puzzle := range puzzles {
        mask := bits(puzzle)
        num := 0

        fb := 1 << (puzzle[0] - 'a')

        for cur := mask; cur != 0; cur = (cur-1) & mask {
            if cur & fb == 0 {
                continue
            }
            if v, ok := freqWords[cur]; ok {
                num += v
            }
        }
        result[i] = num
    }
    return result
}


func bits(s string) int {
    var bits int
    for _, l := range s {
        bits |= 1 << (l - 'a')
    }
    return bits
}

// Time Limit Exceeded
func ngSolution(words []string, puzzles []string) []int {
	result := make([]int, len(puzzles))

	m := make([]map[rune]int, len(words))
	for i, word := range words {
		m[i] = map[rune]int{}
		for _, c := range word {
			m[i][c]++
		}
	}

	for i, puzzle := range puzzles {
		for j, word := range words {
			if isOK(puzzle, word, m[j]) {
				result[i]++
			}
		}
	}

	return result
}

func isOK(puzzle string, word string, m map[rune]int) bool {
	count := int(0)
	for _, char := range puzzle {
		if count == 0 && m[char] == 0 {
			return false
		}
		count += m[char]
		if count == len(word) {
			break
		}
	}
	return count == len(word)
}

func main() {
	// result: [1,1,3,2,4,0]
	words := []string{"aaaa","asas","able","ability","actt","actor","access"}
	puzzles := []string{"aboveyz","abrodyz","abslute","absoryz","actresz","gaswxyz"}

	// result: [0,1,3,2,0]
	// words := []string{"apple","pleas","please"}
	// puzzles := []string{"aelwxyz","aelpxyz","aelpsxy","saelpxy","xaelpsy"}

	// result: 
	// words := []string{}
	// puzzles := []string{}

	result := findNumOfValidWords(words, puzzles)
	fmt.Printf("result = %v\n", result)
}

