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

// Wrong Answer
func ngSolution(stamp string, target string) []int {
	result := []int{}
	for _, sRune := range stamp {
		if !strings.Contains(target, string(sRune)) {
			return result
		}
	}

	sames := []int{}
	i := int(0)
	for i + len(stamp) - 1 < len(target) {
		split := target[i:i + len(stamp)]
		fmt.Printf("split = %v\n", split)
		if split == stamp {
			sames = append(sames, i)
			i++
			continue
		}
		for j, _ := range split {
			if  stamp[j] == split[j] {
				stampIndex := i + j
				// fmt.Printf("stampIndex = %v\n", stampIndex)
				if !contains(stampIndex, result) {
					result = append(result, i)
				}
				break
			}
		}
		i++
	}
	// fmt.Printf("sames = %v\n", sames)
	result = append(result, sames...)
	return result
}

func contains(num int, nums []int) bool {
	for _, n := range nums {
		if num == n {
			return true
		}
	}
	return false
}

func reverse(nums []int) []int {
	for i := 0; i < len(nums) / 2; i++ {
		nums[i], nums[len(nums) - i - 1] = nums[len(nums) - i - 1], nums[i]
	}
	return nums
}

func main() {
	// stamp := "abc"
	// target := "ababc"

	// stamp := "abca"
	// target := "aabcaca"

	// stamp := "oz"
	// target := "ooozz"

	// stamp := "aye"
	// target := "eyeye"

	stamp := "cab"
	target := "cabbb"

	result := movesToStamp(stamp, target)
	fmt.Printf("result = %v\n", result)
}

