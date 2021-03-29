package main
import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func originalDigits(s string) string {
	digits := [][]interface{}{
		{0, "z", []int{}},
		{2, "w", []int{}},
		{4, "u", []int{}},
		{6, "x", []int{}},
		{8, "g", []int{}},
		{5, "f", []int{4}},
		{7, "s", []int{6}},
		{3, "h", []int{8}},
		{9, "i", []int{6, 8, 5}},
		{1, "o", []int{0, 2, 4}},
	}

	resultNums := map[int]int{}
	for i := 0; i < 10; i++ {
		item := digits[i]
		digit := item[0].(int)
		char := item[1].(string)
		rems := item[2].([]int)
		count := strings.Count(s, char)
		for _, rem := range rems {
			count -= resultNums[rem]
		}
		resultNums[digit] += count
	}
	result := ""
	for i := 0; i < 10; i++ {
		for j := 0; j < resultNums[i]; j++ {
			result += strconv.Itoa(i)
		}
	}
	return result
}

// Slow & Use more memory
func mySolution(s string) string {
	// zero -> z is unique, two -> w is unique, six -> x is unique, eight -> g is unique
	firstLoopDigits := map[string]string{"z": "zero", "w": "two", "x": "six", "g": "eight"}
	// three -> t is unique, five -> f is unique, seven -> s is unique
	secondLoopDigits := map[string]string{"t": "three", "s": "seven"}
	// four -> r is unique
	thirdLoopDigits := map[string]string{"r": "four"}
	// one -> o is unique, five -> f is unique 
	fourthLoopDigits := map[string]string{"o": "one", "f": "five"}
	fifthLoopDigits := map[string]string{"n": "nine"}
	// digits := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	digitsMap := map[string]string{
		"zero": "0",
		"one": "1",
		"two": "2",
		"three": "3",
		"four": "4",
		"five": "5",
		"six": "6",
		"seven": "7",
		"eight": "8",
		"nine": "9",
	}

	nums := []string{}
	loopDigits := []map[string]string{
		firstLoopDigits,
		secondLoopDigits,
		thirdLoopDigits,
		fourthLoopDigits,
		fifthLoopDigits,
	}
	for _, digits := range loopDigits {
		for c, digit := range digits {
			sIndex := strings.Index(s, c)
			// fmt.Printf("s = %v, c = %v\n", s, c)
			for sIndex != -1 {
				nums = append(nums, digitsMap[digit])
				for _, r := range digit {
					// fmt.Printf("r = %v\n", string(r))
					s = strings.Replace(s, string(r), "", 1)
				}
				// fmt.Printf("replaced s = %v\n", s)
				sIndex = strings.Index(s, c)
			}
		}
	}
	sort.Slice(nums, func (a, b int) bool { return nums[a] < nums[b] })
	return strings.Join(nums, "")
}

func main() {
	// s := "owoztneoer"
	// s := "fviefuro"
	s := "zeroonetwothreefourfivesixseveneightnine"

	result := originalDigits(s)
	fmt.Printf("result = %v\n", result)
}

