package main
import (
	"fmt"
)

func removeDuplicates(s string, k int) string {
    var sb strings.Builder
    stack := []rune{}
    count := []int{}
    for _, char := range s {
		//empty so always push
        if len(stack) == 0 {
            stack = append(stack, char)
            count = append(count, 1)
        } else {
			//new character
            if stack[len(stack) - 1] != char {
                stack = append(stack, char)
                count = append(count, 1)
            } else { //same char
                if count[len(count) - 1] == k - 1 {
                    count = count[:len(count) - 1]
                    stack = stack[:len(stack) - 1]
                } else {
                    count[len(count) - 1]++
                }
            }
        }
    }
    for i, _ := range stack {
        for j:=0; j < count[i]; j++ {
            sb.WriteString(string(stack[i]))
        }
    }
    ret := sb.String()
    return ret
}

// Time Limit Exceeded
func ngSolution(s string, k int) string {
	result := s
	for true {
		newResult := recursive(result, k)
		if result == newResult {
			break
		}
		result = newResult
		// fmt.Printf("current = %v\n", result)
	}
	return result
}

func recursive(s string, k int) string {
	counter := int(1)
	result := ""

	for i, r := range s {
		if counter == k {
			counter = 1
			continue
		}

		if i == len(s) - 1 {
			for j := 0; j < counter; j++ {
				result += string(r)
			}
			break
		}

		nextRune := rune(s[i + 1])
		if r == nextRune {
			counter++
		} else {
			for j := 0; j < counter; j++ {
				result += string(r)
			}
			counter = 1
		}
	}
	return result
}

func main() {
	// result: abcd
	// s := "abcd"
	// k := int(2)

	// result: aa
	// s := "deeedbbcccbdaa"
	// k := int(3)

	// result: ps
	// s := "pbbcggttciiippooaais"
	// k := int(2)

	// result: ybth
	// s := "yfttttfbbbbnnnnffbgffffgbbbbgssssgthyyyy"
	// k := int(4)

	// result: izzlz
	// s := "iiiixxxxxiiccccczzffffflllllllllfffffllyyyyyuuuuuz"
	// k := int(5)

	// result: ghayqgq
	s := "ghanyhhhhhttttttthhyyyyyynnnnnnyqkkkkkkkrrrrrrjjjjjjjryyyyyyfffffffygq"
	k := int(7)

	// result: 
	// s := ""
	// k := int(0)

	result := removeDuplicates(s, k)
	fmt.Printf("result = %v\n", result)
}

