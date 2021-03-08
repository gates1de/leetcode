package main
import (
	"fmt"
	"reflect"
	"sort"
	"strings"
)

type SortRunes []rune

func (s SortRunes) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s SortRunes) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func (s SortRunes) Len() int {
    return len(s)
}

func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    dict := make(map[rune]int)
    for _, v := range s {
        dict[v]++
    }
    for _, v := range t {
        dict[v]--
    }
    for k, _ := range dict {
        if dict[k] != 0 {
            return false
        }
    }
    return true
}

func myAnswer2(s string, t string) bool {
	sChars := strings.Split(s, "")
	tChars := strings.Split(t, "")
	sort.Slice(sChars, func (i, j int) bool { return sChars[i] < sChars[j] })
	sort.Slice(tChars, func (i, j int) bool { return tChars[i] < tChars[j] })
	fmt.Printf("s = %v, t = %v\n", sChars, tChars)
	return reflect.DeepEqual(sChars, tChars)
}

func myAnswer1(s string, t string) bool {
	sortedS := sortString(s)
	sortedT := sortString(t)
	fmt.Printf("s = %v, t = %v\n", sortedS, sortedT)
	return sortedS == sortedT
}

func sortString(s string) string {
    r := []rune(s)
    sort.Sort(SortRunes(r))
    return string(r)
}

func main() {
	s := "anagram"
	t := "nagaram"

	// s := "rat"
	// t := "car"

	result := isAnagram(s, t)
	fmt.Printf("result = %v\n", result)
}

