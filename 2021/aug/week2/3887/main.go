package main
import (
	"fmt"
)

func groupAnagrams(words []string) [][]string {
	cache := make(map[[26]byte]int)
	result := make([][]string, 0)
	for i := range words {
		list := [26]byte{}
		for j := range words[i] {
			list[words[i][j]-'a']++
		}
		if idx, ok := cache[list]; ok {
			result[idx] = append(result[idx], words[i])
		} else {
			result = append(result, []string{words[i]})
			cache[list] = len(result) - 1
		}
	}
	return result
}

// Slow & Use more memory
func myAnswer(strs []string) [][]string {
	strsPerLen := map[int][]string{}
	for _, str := range strs {
		length := len(str)
		if strsPerLen[length] == nil {
			strsPerLen[length] = []string{str}
		} else {
			strsPerLen[length] = append(strsPerLen[length], str)
		}
	}
	// fmt.Printf("strsPerLen = %v\n", strsPerLen)
	result := [][]string{}

	for _, arr := range strsPerLen {
		for i, s := range arr {
			if s == "0" {
				continue
			}
			group := []string{s}
			arr[i] = "0"
			for j, t := range arr {
				if t == "0" {
					continue
				}
				if isAnagram(s, t) {
					group = append(group, t)
					arr[j] = "0"
				}
			}
			result = append(result, group)
		}
	}
	return result
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

func main() {
	// result: [["bat"],["nat","tan"],["ate","eat","tea"]]
	// strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}

	// result: [[""]]
	// strs := []string{""}

	// result: [["a"]]
	// strs := []string{"a"}

	// result: [[eat tea ate] [tan nat] [bat] [t] [j] [ti it]]
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat", "t", "ti", "j", "it"}

	// result: 
	// strs := []string{}

	result := groupAnagrams(strs)
	fmt.Printf("result = %v\n", result)
}

