package main
import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type Counter struct {
	key string
	count int
}

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

// Same speed & Use less memory than above
func mySolution1(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	chars1 := strings.Split(word1, "")
	chars2 := strings.Split(word2, "")
	chars1Count := map[string]int{}
	chars2Count := map[string]int{}
	chars1Counter := []Counter{}
	chars2Counter := []Counter{}
	counts1 := []int{}
	counts2 := []int{}
	for _, c := range chars1 {
		chars1Count[c]++
	}
	for _, c := range chars2 {
		chars2Count[c]++
	}
	for key, value := range chars1Count {
		counts1 = append(counts1, value)
		chars1Counter = append(chars1Counter, Counter{key: key, count: value})
	}
	for key, value := range chars2Count {
		counts2 = append(counts2, value)
		chars2Counter = append(chars2Counter, Counter{key: key, count: value})
	}
	sort.Slice(counts1, func (i, j int) bool { return counts1[i] < counts1[j] })
	sort.Slice(counts2, func (i, j int) bool { return counts2[i] < counts2[j] })
	fmt.Printf("counts1 = %v, counts2 = %v\n", counts1, counts2)
	if !reflect.DeepEqual(counts1, counts2) {
		return false
	}

	sort.Slice(chars1Counter, func (i, j int) bool { return chars1Counter[i].count < chars1Counter[j].count })
	sort.Slice(chars2Counter, func (i, j int) bool { return chars2Counter[i].count < chars2Counter[j].count })
	fmt.Printf("chars1Counter = %v, chars2Counter = %v\n", chars1Counter, chars2Counter)

	for i, counter1 := range chars1Counter {
		counter2 := chars2Counter[i]
		if counter1.key == counter2.key && counter1.count != counter2.count {
			return false
		}
		if counter1.key != counter2.key && chars1Count[counter2.key] == 0 {
			return false
		}
	}
	return true
}

func main() {
	// word1 := "abc"
	// word2 := "cba"

	// word1 := "a"
	// word2 := "aa"

	// word1 := "cabbba"
	// word2 := "abbccc"

	// word1 := "cabbba"
	// word2 := "aabbss"

	// word1 := "abbbzcf"
	// word2 := "babzzcz"

	// word1 := "uau"
	// word2 := "ssx"

	// word1 := "abbzzca"
	// word2 := "babzzcz"

	word1 := "uuukuuuukkuusuususuuuukuskuusuuusuusuuuuuuk"
	word2 := "kssskkskkskssskksskskksssssksskksskskksksuu"

	result := closeStrings(word1, word2)
	fmt.Printf("result = %v\n", result)
}

