package main
import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func closeStrings(word1 string, word2 string) bool {
    chars1Count := map[string]int{}
    chars2Count := map[string]int{}
    keys1 := []string{}
    keys2 := []string{}
    counts1 := []int{}
    counts2 := []int{}

    chars1 := strings.Split(word1, "")
    chars2 := strings.Split(word2, "")
    for _, c := range chars1 {
        chars1Count[c]++
    }
    for _, c := range chars2 {
        chars2Count[c]++
    }

    for key, value := range chars1Count {
        keys1 = append(keys1, key)
        counts1 = append(counts1, value)
    }
    for key, value := range chars2Count {
        keys2 = append(keys2, key)
        counts2 = append(counts2, value)
    }

    sort.Slice(keys1, func (i, j int) bool { return keys1[i] < keys1[j] })
    sort.Slice(keys2, func (i, j int) bool { return keys2[i] < keys2[j] })
    sort.Slice(counts1, func (i, j int) bool { return counts1[i] < counts1[j] })
    sort.Slice(counts2, func (i, j int) bool { return counts2[i] < counts2[j] })

    if !reflect.DeepEqual(keys1, keys2) {
        return false
    }

    return reflect.DeepEqual(counts1, counts2)
}

func main() {
	// result: true
	// word1 := "abc"
	// word2 := "bca"

	// result: false
	// word1 := "a"
	// word2 := "aa"

	// result: true
	// word1 := "cabbba"
	// word2 := "abbccc"

	// result: false
	word1 := "uau"
	word2 := "ssx"

	// result: 
	// word1 := ""
	// word2 := ""

	result := closeStrings(word1, word2)
	fmt.Printf("result = %v\n", result)
}

