package main
import (
	"fmt"
	"sort"
)

func groupAnagrams(strs []string) [][]string {
    m := map[string][]string{}
    for _, str := range strs {
        chars := []byte(str)
        sort.Slice(chars, func(a, b int) bool { return chars[a] < chars[b] })
        sortedStr := string(chars)
        if _, ok := m[sortedStr]; ok {
            m[sortedStr] = append(m[sortedStr], str)
        } else {
            m[sortedStr] = []string{str}
        }
    }

    result := make([][]string, 0, len(m))
    for _, val := range m {
        result = append(result, val)
    }

    return result
}

func main() {
	// result: [["bat"],["nat","tan"],["ate","eat","tea"]]
	// strs := []string{"eat","tea","tan","ate","nat","bat"}

	// result: [[""]]
	// strs := []string{""}

	// result: [["a"]]
	strs := []string{"a"}

	// result: 
	// strs := []string{}

	result := groupAnagrams(strs)
	fmt.Printf("result = %v\n", result)
}

