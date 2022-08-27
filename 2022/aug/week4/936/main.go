package main
import (
	"fmt"
	"strings"
)

func movesToStamp(stamp string, target string) []int {
    letters := []byte(target)
    toErase := len(target)
    contains := func(offset int) bool {
        empty := true
        for i := 0; i < len(stamp); i++ {
            if l := letters[offset + i]; l != '?' && l != stamp[i] {
                return false
            } else if l != '?' {
                empty = false
            }
        }
        return !empty
    }

    erase := func(offset int) {
        for i := 0; i < len(stamp); i++ {
            if letters[offset + i] != '?' {
                letters[offset + i] = '?'
                toErase--
            }
        }
    }

    var result []int
    for toErase != 0 {
        erased := false
        for i := 0; i < len(letters) - len(stamp) + 1; i++ {
            if contains(i) {
                erase(i)
                result = append(result, i)
                erased = true
                i += len(stamp)
                if toErase == 0 {
                    break
                }
            }
        }
        if !erased {
            return []int{}
        }
    }

    for l, r := 0, len(result) - 1; l < r; l, r = l + 1, r - 1 {
        result[l], result[r] = result[r], result[l]
    }

    return result
}

// Time Limis Exceedea
func ngSolution(stamp string, target string) []int {
	stampLen := len(stamp)
	targetLen := len(target)
	if stamp[stampLen - 1] != target[targetLen - 1] {
		return []int{}
	}

	maxIndex := len(target) - len(stamp)
	validIndexes := []int{}
	for i := 0; i <= maxIndex; i++ {
		if isValid(stamp, target[i:i + len(stamp)]) {
			validIndexes = append(validIndexes, i)
		}
	}

	s := strings.Repeat("?", len(target))

	return helper(s, []int{}, validIndexes, stamp, target)
}

func isValid(s1 string, s2 string) bool {
	for i := 0; i < len(s1); i++ {
		if s1[i] == s2[i] {
			return true
		}
	}
	return false
}

func helper(s string, indexes []int, validIndexes []int, stamp string, target string) []int {
	fmt.Println(s, validIndexes)
	if s == target {
		return indexes
	}

	if len(validIndexes) == 0 {
		return []int{}
	}

	for i, index := range validIndexes {
		newS := replaceRange(s, index, index + len(stamp), stamp)

		newIndexes := make([]int, len(indexes) + 1)
		copy(newIndexes, indexes)
		newIndexes[len(indexes)] = index

		newValidIndexes := make([]int, len(validIndexes))
		copy(newValidIndexes, validIndexes)
		newValidIndexes = remove(newValidIndexes, i)

		newIndexes = helper(newS, newIndexes, newValidIndexes, stamp, target)
		if len(newIndexes) > 0 {
			return newIndexes
		}
	}

	return []int{}
}

func replaceRange(s string, start int, end int, subs string) string {
	return s[:start] + subs + s[end:]
}

func remove(arr []int, i int) []int {
	arr[i] = arr[len(arr) - 1]
    return arr[:len(arr) - 1]
}

func main() {
	// result: [0,2]
	// stamp := "abc"
	// target := "ababc"

	// result: [3,0,1]
	// stamp := "abca"
	// target := "aabcaca"

	// result: [0,3,1,2]
	// stamp := "oz"
	// target := "ooozz"

	// result: [3,5,4,7,6,9,8,11,10,1,0,12,2]
	stamp := "ffc"
	target := "ffffcfffffffffc"

	// result: 
	// stamp := ""
	// target := ""

	result := movesToStamp(stamp, target)
	fmt.Printf("result = %v\n", result)
}

